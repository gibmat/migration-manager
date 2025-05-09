package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	incusTLS "github.com/lxc/incus/v6/shared/tls"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/internal/server/auth"
	"github.com/FuturFusion/migration-manager/internal/server/response"
	"github.com/FuturFusion/migration-manager/internal/server/util"
	apiSource "github.com/FuturFusion/migration-manager/internal/source"
	"github.com/FuturFusion/migration-manager/internal/transaction"
	"github.com/FuturFusion/migration-manager/shared/api"
)

var sourcesCmd = APIEndpoint{
	Path: "sources",

	Get:  APIEndpointAction{Handler: sourcesGet, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanView)},
	Post: APIEndpointAction{Handler: sourcesPost, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanCreate)},
}

var sourceCmd = APIEndpoint{
	Path: "sources/{name}",

	Delete: APIEndpointAction{Handler: sourceDelete, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanDelete)},
	Get:    APIEndpointAction{Handler: sourceGet, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanView)},
	Put:    APIEndpointAction{Handler: sourcePut, AccessHandler: allowPermission(auth.ObjectTypeServer, auth.EntitlementCanEdit)},
}

// swagger:operation GET /1.0/sources sources sources_get
//
//	Get the sources
//
//	Returns a list of sources (URLs).
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: API sources
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          type: array
//	          description: List of sources
//                items:
//                  type: string
//                example: |-
//                  [
//                    "/1.0/sources/foo",
//                    "/1.0/sources/bar"
//                  ]
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"

// swagger:operation GET /1.0/sources?recursion=1 sources sources_get_recursion
//
//	Get the sources
//
//	Returns a list of sources (structs).
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: API sources
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          type: array
//	          description: List of sources
//	          items:
//	            $ref: "#/definitions/Source"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func sourcesGet(d *Daemon, r *http.Request) response.Response {
	// Parse the recursion field.
	recursion, err := strconv.Atoi(r.FormValue("recursion"))
	if err != nil {
		recursion = 0
	}

	if recursion == 1 {
		sources, err := d.source.GetAll(r.Context())
		if err != nil {
			return response.SmartError(err)
		}

		result := make([]api.Source, 0, len(sources))
		for _, source := range sources {
			result = append(result, source.ToAPI())
		}

		return response.SyncResponse(true, result)
	}

	sourceNames, err := d.source.GetAllNames(r.Context())
	if err != nil {
		return response.SmartError(err)
	}

	result := make([]string, 0, len(sourceNames))
	for _, name := range sourceNames {
		result = append(result, fmt.Sprintf("/%s/sources/%s", api.APIVersion, name))
	}

	return response.SyncResponse(true, result)
}

// swagger:operation POST /1.0/sources sources sources_post
//
//	Add a source
//
//	Creates a new source.
//
//	---
//	consumes:
//	  - application/json
//	produces:
//	  - application/json
//	parameters:
//	  - in: body
//	    name: source
//	    description: Source configuration
//	    required: true
//	    schema:
//	      $ref: "#/definitions/Source"
//	responses:
//	  "200":
//	    $ref: "#/responses/EmptySyncResponse"
//	  "400":
//	    $ref: "#/responses/BadRequest"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func sourcesPost(d *Daemon, r *http.Request) response.Response {
	var source api.Source

	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		return response.BadRequest(err)
	}

	src, err := d.source.Create(r.Context(), migration.Source{
		Name:       source.Name,
		SourceType: source.SourceType,
		Properties: source.Properties,
		EndpointFunc: func(s api.Source) (migration.SourceEndpoint, error) {
			return apiSource.NewInternalVMwareSourceFrom(s)
		},
	})
	if err != nil {
		return response.SmartError(fmt.Errorf("Failed creating source %q: %w", source.Name, err))
	}

	// Trigger a scan of this new source for instances.
	if src.GetExternalConnectivityStatus() == api.EXTERNALCONNECTIVITYSTATUS_OK {
		err = d.syncOneSource(r.Context(), src)
		if err != nil {
			return response.SmartError(fmt.Errorf("Failed to initiate sync from source %q: %w", src.Name, err))
		}
	}

	metadata := make(map[string]string)
	metadata["ConnectivityStatus"] = string(src.GetExternalConnectivityStatus())

	// If waiting on fingerprint confirmation, return it to the user.
	if src.GetExternalConnectivityStatus() == api.EXTERNALCONNECTIVITYSTATUS_TLS_CONFIRM_FINGERPRINT {
		metadata["certFingerprint"] = incusTLS.CertFingerprint(src.GetServerCertificate())
	}

	return response.SyncResponseLocation(true, metadata, "/"+api.APIVersion+"/sources/"+source.Name)
}

// swagger:operation DELETE /1.0/sources/{name} sources source_delete
//
//	Delete the source
//
//	Removes the source.
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    $ref: "#/responses/EmptySyncResponse"
//	  "400":
//	    $ref: "#/responses/BadRequest"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func sourceDelete(d *Daemon, r *http.Request) response.Response {
	name := r.PathValue("name")

	err := d.source.DeleteByName(r.Context(), name, d.instance)
	if err != nil {
		return response.SmartError(err)
	}

	return response.EmptySyncResponse
}

// swagger:operation GET /1.0/sources/{name} sources source_get
//
//	Get the source
//
//	Gets a specific source.
//
//	---
//	produces:
//	  - application/json
//	responses:
//	  "200":
//	    description: Source
//	    schema:
//	      type: object
//	      description: Sync response
//	      properties:
//	        type:
//	          type: string
//	          description: Response type
//	          example: sync
//	        status:
//	          type: string
//	          description: Status description
//	          example: Success
//	        status_code:
//	          type: integer
//	          description: Status code
//	          example: 200
//	        metadata:
//	          $ref: "#/definitions/Source"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func sourceGet(d *Daemon, r *http.Request) response.Response {
	name := r.PathValue("name")

	source, err := d.source.GetByName(r.Context(), name)
	if err != nil {
		return response.SmartError(err)
	}

	return response.SyncResponseETag(
		true,
		source.ToAPI(),
		source,
	)
}

// swagger:operation PUT /1.0/sources/{name} sources source_put
//
//	Update the source
//
//	Updates the source definition.
//
//	---
//	consumes:
//	  - application/json
//	produces:
//	  - application/json
//	parameters:
//	  - in: body
//	    name: source
//	    description: Source definition
//	    required: true
//	    schema:
//	      $ref: "#/definitions/Source"
//	responses:
//	  "200":
//	    $ref: "#/responses/EmptySyncResponse"
//	  "400":
//	    $ref: "#/responses/BadRequest"
//	  "403":
//	    $ref: "#/responses/Forbidden"
//	  "412":
//	    $ref: "#/responses/PreconditionFailed"
//	  "500":
//	    $ref: "#/responses/InternalServerError"
func sourcePut(d *Daemon, r *http.Request) response.Response {
	name := r.PathValue("name")

	var source api.SourcePut

	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		return response.BadRequest(err)
	}

	ctx, trans := transaction.Begin(r.Context())
	defer func() {
		rollbackErr := trans.Rollback()
		if rollbackErr != nil {
			response.SmartError(fmt.Errorf("Transaction rollback failed: %v, reason: %w", rollbackErr, err))
		}
	}()

	currentSource, err := d.source.GetByName(ctx, name)
	if err != nil {
		return response.SmartError(err)
	}

	// Validate ETag
	err = util.EtagCheck(r, currentSource)
	if err != nil {
		return response.PreconditionFailed(err)
	}

	src := &migration.Source{
		ID:         currentSource.ID,
		Name:       source.Name,
		SourceType: currentSource.SourceType,
		Properties: source.Properties,
		EndpointFunc: func(s api.Source) (migration.SourceEndpoint, error) {
			return apiSource.NewInternalVMwareSourceFrom(s)
		},
	}

	err = d.source.Update(ctx, name, src, d.instance)
	if err != nil {
		return response.SmartError(fmt.Errorf("Failed updating source %q: %w", source.Name, err))
	}

	err = trans.Commit()
	if err != nil {
		return response.SmartError(fmt.Errorf("Failed commit transaction: %w", err))
	}

	// Trigger a scan of this new source for instances.
	if src.GetExternalConnectivityStatus() == api.EXTERNALCONNECTIVITYSTATUS_OK {
		err = d.syncOneSource(r.Context(), *src)
		if err != nil {
			return response.SmartError(fmt.Errorf("Failed to initiate sync from source %q: %w", src.Name, err))
		}
	}

	metadata := make(map[string]string)
	metadata["ConnectivityStatus"] = string(src.GetExternalConnectivityStatus())

	// If waiting on fingerprint confirmation, return it to the user.
	if src.GetExternalConnectivityStatus() == api.EXTERNALCONNECTIVITYSTATUS_TLS_CONFIRM_FINGERPRINT {
		metadata["certFingerprint"] = incusTLS.CertFingerprint(src.GetServerCertificate())
	}

	return response.SyncResponseLocation(true, metadata, "/"+api.APIVersion+"/sources/"+source.Name)
}
