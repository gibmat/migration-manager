// Code generated by generate-database from the incus project - DO NOT EDIT.

package entities

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
)

var instanceOverrideObjects = RegisterStmt(`
SELECT instance_overrides.properties, instance_overrides.id, instance_overrides.uuid, instance_overrides.last_update, instance_overrides.comment, instance_overrides.disable_migration
  FROM instance_overrides
  ORDER BY instance_overrides.uuid
`)

var instanceOverrideObjectsByUUID = RegisterStmt(`
SELECT instance_overrides.properties, instance_overrides.id, instance_overrides.uuid, instance_overrides.last_update, instance_overrides.comment, instance_overrides.disable_migration
  FROM instance_overrides
  WHERE ( instance_overrides.uuid = ? )
  ORDER BY instance_overrides.uuid
`)

var instanceOverrideID = RegisterStmt(`
SELECT instance_overrides.id FROM instance_overrides
  WHERE instance_overrides.uuid = ?
`)

var instanceOverrideCreate = RegisterStmt(`
INSERT INTO instance_overrides (properties, uuid, last_update, comment, disable_migration)
  VALUES (?, ?, ?, ?, ?)
`)

var instanceOverrideUpdate = RegisterStmt(`
UPDATE instance_overrides
  SET properties = ?, uuid = ?, last_update = ?, comment = ?, disable_migration = ?
 WHERE id = ?
`)

var instanceOverrideDeleteByUUID = RegisterStmt(`
DELETE FROM instance_overrides WHERE uuid = ?
`)

// GetInstanceOverrideID return the ID of the instance_override with the given key.
// generator: instance_override ID
func GetInstanceOverrideID(ctx context.Context, db tx, uuid uuid.UUID) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Instance_override")
	}()

	stmt, err := Stmt(db, instanceOverrideID)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"instanceOverrideID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, uuid)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, ErrNotFound
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to get \"instance_overrides\" ID: %w", err)
	}

	return id, nil
}

// GetInstanceOverride returns the instance_override with the given key.
// generator: instance_override GetOne
func GetInstanceOverride(ctx context.Context, db dbtx, uuid uuid.UUID) (_ *migration.InstanceOverride, _err error) {
	defer func() {
		_err = mapErr(_err, "Instance_override")
	}()

	filter := InstanceOverrideFilter{}
	filter.UUID = &uuid

	objects, err := GetInstanceOverrides(ctx, db, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instance_overrides\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, ErrNotFound
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"instance_overrides\" entry matches")
	}
}

// instanceOverrideColumns returns a string of column names to be used with a SELECT statement for the entity.
// Use this function when building statements to retrieve database entries matching the InstanceOverride entity.
func instanceOverrideColumns() string {
	return "instance_overrides.properties, instance_overrides.id, instance_overrides.uuid, instance_overrides.last_update, instance_overrides.comment, instance_overrides.disable_migration"
}

// getInstanceOverrides can be used to run handwritten sql.Stmts to return a slice of objects.
func getInstanceOverrides(ctx context.Context, stmt *sql.Stmt, args ...any) ([]migration.InstanceOverride, error) {
	objects := make([]migration.InstanceOverride, 0)

	dest := func(scan func(dest ...any) error) error {
		i := migration.InstanceOverride{}
		var propertiesStr string
		err := scan(&propertiesStr, &i.ID, &i.UUID, &i.LastUpdate, &i.Comment, &i.DisableMigration)
		if err != nil {
			return err
		}

		err = unmarshalJSON(propertiesStr, &i.Properties)
		if err != nil {
			return err
		}

		objects = append(objects, i)

		return nil
	}

	err := selectObjects(ctx, stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instance_overrides\" table: %w", err)
	}

	return objects, nil
}

// getInstanceOverridesRaw can be used to run handwritten query strings to return a slice of objects.
func getInstanceOverridesRaw(ctx context.Context, db dbtx, sql string, args ...any) ([]migration.InstanceOverride, error) {
	objects := make([]migration.InstanceOverride, 0)

	dest := func(scan func(dest ...any) error) error {
		i := migration.InstanceOverride{}
		var propertiesStr string
		err := scan(&propertiesStr, &i.ID, &i.UUID, &i.LastUpdate, &i.Comment, &i.DisableMigration)
		if err != nil {
			return err
		}

		err = unmarshalJSON(propertiesStr, &i.Properties)
		if err != nil {
			return err
		}

		objects = append(objects, i)

		return nil
	}

	err := scan(ctx, db, sql, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instance_overrides\" table: %w", err)
	}

	return objects, nil
}

// GetInstanceOverrides returns all available instance_overrides.
// generator: instance_override GetMany
func GetInstanceOverrides(ctx context.Context, db dbtx, filters ...InstanceOverrideFilter) (_ []migration.InstanceOverride, _err error) {
	defer func() {
		_err = mapErr(_err, "Instance_override")
	}()

	var err error

	// Result slice.
	objects := make([]migration.InstanceOverride, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	args := []any{}
	queryParts := [2]string{}

	if len(filters) == 0 {
		sqlStmt, err = Stmt(db, instanceOverrideObjects)
		if err != nil {
			return nil, fmt.Errorf("Failed to get \"instanceOverrideObjects\" prepared statement: %w", err)
		}
	}

	for i, filter := range filters {
		if filter.UUID != nil {
			args = append(args, []any{filter.UUID}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, instanceOverrideObjectsByUUID)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"instanceOverrideObjectsByUUID\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(instanceOverrideObjectsByUUID)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"instanceOverrideObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.UUID == nil {
			return nil, fmt.Errorf("Cannot filter on empty InstanceOverrideFilter")
		} else {
			return nil, fmt.Errorf("No statement exists for the given Filter")
		}
	}

	// Select.
	if sqlStmt != nil {
		objects, err = getInstanceOverrides(ctx, sqlStmt, args...)
	} else {
		queryStr := strings.Join(queryParts[:], "ORDER BY")
		objects, err = getInstanceOverridesRaw(ctx, db, queryStr, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instance_overrides\" table: %w", err)
	}

	return objects, nil
}

// CreateInstanceOverride adds a new instance_override to the database.
// generator: instance_override Create
func CreateInstanceOverride(ctx context.Context, db dbtx, object migration.InstanceOverride) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Instance_override")
	}()

	args := make([]any, 5)

	// Populate the statement arguments.
	marshaledProperties, err := marshalJSON(object.Properties)
	if err != nil {
		return -1, err
	}

	args[0] = marshaledProperties
	args[1] = object.UUID
	args[2] = object.LastUpdate
	args[3] = object.Comment
	args[4] = object.DisableMigration

	// Prepared statement to use.
	stmt, err := Stmt(db, instanceOverrideCreate)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"instanceOverrideCreate\" prepared statement: %w", err)
	}

	// Execute the statement.
	result, err := stmt.Exec(args...)
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return -1, ErrConflict
		}
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to create \"instance_overrides\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"instance_overrides\" entry ID: %w", err)
	}

	return id, nil
}

// UpdateInstanceOverride updates the instance_override matching the given key parameters.
// generator: instance_override Update
func UpdateInstanceOverride(ctx context.Context, db tx, uuid uuid.UUID, object migration.InstanceOverride) (_err error) {
	defer func() {
		_err = mapErr(_err, "Instance_override")
	}()

	id, err := GetInstanceOverrideID(ctx, db, uuid)
	if err != nil {
		return err
	}

	stmt, err := Stmt(db, instanceOverrideUpdate)
	if err != nil {
		return fmt.Errorf("Failed to get \"instanceOverrideUpdate\" prepared statement: %w", err)
	}

	marshaledProperties, err := marshalJSON(object.Properties)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(marshaledProperties, object.UUID, object.LastUpdate, object.Comment, object.DisableMigration, id)
	if err != nil {
		return fmt.Errorf("Update \"instance_overrides\" entry failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}

// DeleteInstanceOverride deletes the instance_override matching the given key parameters.
// generator: instance_override DeleteOne-by-UUID
func DeleteInstanceOverride(ctx context.Context, db dbtx, uuid uuid.UUID) (_err error) {
	defer func() {
		_err = mapErr(_err, "Instance_override")
	}()

	stmt, err := Stmt(db, instanceOverrideDeleteByUUID)
	if err != nil {
		return fmt.Errorf("Failed to get \"instanceOverrideDeleteByUUID\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(uuid)
	if err != nil {
		return fmt.Errorf("Delete \"instance_overrides\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n == 0 {
		return ErrNotFound
	} else if n > 1 {
		return fmt.Errorf("Query deleted %d InstanceOverride rows instead of 1", n)
	}

	return nil
}
