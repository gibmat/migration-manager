// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package migration_test

import (
	"context"
	"sync"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/shared/api"
	"github.com/google/uuid"
)

// Ensure, that InstanceServiceMock does implement migration.InstanceService.
// If this is not the case, regenerate this file with moq.
var _ migration.InstanceService = &InstanceServiceMock{}

// InstanceServiceMock is a mock implementation of migration.InstanceService.
//
//	func TestSomethingThatUsesInstanceService(t *testing.T) {
//
//		// make and configure a mocked migration.InstanceService
//		mockedInstanceService := &InstanceServiceMock{
//			CreateFunc: func(ctx context.Context, instance migration.Instance) (migration.Instance, error) {
//				panic("mock out the Create method")
//			},
//			CreateOverridesFunc: func(ctx context.Context, overrides migration.InstanceOverride) (migration.InstanceOverride, error) {
//				panic("mock out the CreateOverrides method")
//			},
//			DeleteByUUIDFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the DeleteByUUID method")
//			},
//			DeleteOverridesByUUIDFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the DeleteOverridesByUUID method")
//			},
//			GetAllFunc: func(ctx context.Context, withOverrides bool) (migration.Instances, error) {
//				panic("mock out the GetAll method")
//			},
//			GetAllByBatchFunc: func(ctx context.Context, batch string, withOverrides bool) (migration.Instances, error) {
//				panic("mock out the GetAllByBatch method")
//			},
//			GetAllByBatchAndStateFunc: func(ctx context.Context, batch string, status api.MigrationStatusType, withOverrides bool) (migration.Instances, error) {
//				panic("mock out the GetAllByBatchAndState method")
//			},
//			GetAllBySourceFunc: func(ctx context.Context, source string, withOverrides bool) (migration.Instances, error) {
//				panic("mock out the GetAllBySource method")
//			},
//			GetAllByStateFunc: func(ctx context.Context, status api.MigrationStatusType, withOverrides bool) (migration.Instances, error) {
//				panic("mock out the GetAllByState method")
//			},
//			GetAllUUIDsFunc: func(ctx context.Context) ([]uuid.UUID, error) {
//				panic("mock out the GetAllUUIDs method")
//			},
//			GetAllUnassignedFunc: func(ctx context.Context, withOverrides bool) (migration.Instances, error) {
//				panic("mock out the GetAllUnassigned method")
//			},
//			GetByUUIDFunc: func(ctx context.Context, id uuid.UUID, withOverrides bool) (*migration.Instance, error) {
//				panic("mock out the GetByUUID method")
//			},
//			GetOverridesByUUIDFunc: func(ctx context.Context, id uuid.UUID) (*migration.InstanceOverride, error) {
//				panic("mock out the GetOverridesByUUID method")
//			},
//			ProcessWorkerUpdateFunc: func(ctx context.Context, id uuid.UUID, workerResponseTypeArg api.WorkerResponseType, statusString string) (migration.Instance, error) {
//				panic("mock out the ProcessWorkerUpdate method")
//			},
//			UnassignFromBatchFunc: func(ctx context.Context, id uuid.UUID) error {
//				panic("mock out the UnassignFromBatch method")
//			},
//			UpdateFunc: func(ctx context.Context, instance *migration.Instance) error {
//				panic("mock out the Update method")
//			},
//			UpdateOverridesFunc: func(ctx context.Context, overrides *migration.InstanceOverride) error {
//				panic("mock out the UpdateOverrides method")
//			},
//			UpdateStatusByUUIDFunc: func(ctx context.Context, i uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (*migration.Instance, error) {
//				panic("mock out the UpdateStatusByUUID method")
//			},
//		}
//
//		// use mockedInstanceService in code that requires migration.InstanceService
//		// and then make assertions.
//
//	}
type InstanceServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, instance migration.Instance) (migration.Instance, error)

	// CreateOverridesFunc mocks the CreateOverrides method.
	CreateOverridesFunc func(ctx context.Context, overrides migration.InstanceOverride) (migration.InstanceOverride, error)

	// DeleteByUUIDFunc mocks the DeleteByUUID method.
	DeleteByUUIDFunc func(ctx context.Context, id uuid.UUID) error

	// DeleteOverridesByUUIDFunc mocks the DeleteOverridesByUUID method.
	DeleteOverridesByUUIDFunc func(ctx context.Context, id uuid.UUID) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context, withOverrides bool) (migration.Instances, error)

	// GetAllByBatchFunc mocks the GetAllByBatch method.
	GetAllByBatchFunc func(ctx context.Context, batch string, withOverrides bool) (migration.Instances, error)

	// GetAllByBatchAndStateFunc mocks the GetAllByBatchAndState method.
	GetAllByBatchAndStateFunc func(ctx context.Context, batch string, status api.MigrationStatusType, withOverrides bool) (migration.Instances, error)

	// GetAllBySourceFunc mocks the GetAllBySource method.
	GetAllBySourceFunc func(ctx context.Context, source string, withOverrides bool) (migration.Instances, error)

	// GetAllByStateFunc mocks the GetAllByState method.
	GetAllByStateFunc func(ctx context.Context, status api.MigrationStatusType, withOverrides bool) (migration.Instances, error)

	// GetAllUUIDsFunc mocks the GetAllUUIDs method.
	GetAllUUIDsFunc func(ctx context.Context) ([]uuid.UUID, error)

	// GetAllUnassignedFunc mocks the GetAllUnassigned method.
	GetAllUnassignedFunc func(ctx context.Context, withOverrides bool) (migration.Instances, error)

	// GetByUUIDFunc mocks the GetByUUID method.
	GetByUUIDFunc func(ctx context.Context, id uuid.UUID, withOverrides bool) (*migration.Instance, error)

	// GetOverridesByUUIDFunc mocks the GetOverridesByUUID method.
	GetOverridesByUUIDFunc func(ctx context.Context, id uuid.UUID) (*migration.InstanceOverride, error)

	// ProcessWorkerUpdateFunc mocks the ProcessWorkerUpdate method.
	ProcessWorkerUpdateFunc func(ctx context.Context, id uuid.UUID, workerResponseTypeArg api.WorkerResponseType, statusString string) (migration.Instance, error)

	// UnassignFromBatchFunc mocks the UnassignFromBatch method.
	UnassignFromBatchFunc func(ctx context.Context, id uuid.UUID) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, instance *migration.Instance) error

	// UpdateOverridesFunc mocks the UpdateOverrides method.
	UpdateOverridesFunc func(ctx context.Context, overrides *migration.InstanceOverride) error

	// UpdateStatusByUUIDFunc mocks the UpdateStatusByUUID method.
	UpdateStatusByUUIDFunc func(ctx context.Context, i uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (*migration.Instance, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Instance is the instance argument value.
			Instance migration.Instance
		}
		// CreateOverrides holds details about calls to the CreateOverrides method.
		CreateOverrides []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Overrides is the overrides argument value.
			Overrides migration.InstanceOverride
		}
		// DeleteByUUID holds details about calls to the DeleteByUUID method.
		DeleteByUUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// DeleteOverridesByUUID holds details about calls to the DeleteOverridesByUUID method.
		DeleteOverridesByUUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetAllByBatch holds details about calls to the GetAllByBatch method.
		GetAllByBatch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch string
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetAllByBatchAndState holds details about calls to the GetAllByBatchAndState method.
		GetAllByBatchAndState []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Batch is the batch argument value.
			Batch string
			// Status is the status argument value.
			Status api.MigrationStatusType
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetAllBySource holds details about calls to the GetAllBySource method.
		GetAllBySource []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Source is the source argument value.
			Source string
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetAllByState holds details about calls to the GetAllByState method.
		GetAllByState []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Status is the status argument value.
			Status api.MigrationStatusType
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetAllUUIDs holds details about calls to the GetAllUUIDs method.
		GetAllUUIDs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllUnassigned holds details about calls to the GetAllUnassigned method.
		GetAllUnassigned []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetByUUID holds details about calls to the GetByUUID method.
		GetByUUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
			// WithOverrides is the withOverrides argument value.
			WithOverrides bool
		}
		// GetOverridesByUUID holds details about calls to the GetOverridesByUUID method.
		GetOverridesByUUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// ProcessWorkerUpdate holds details about calls to the ProcessWorkerUpdate method.
		ProcessWorkerUpdate []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
			// WorkerResponseTypeArg is the workerResponseTypeArg argument value.
			WorkerResponseTypeArg api.WorkerResponseType
			// StatusString is the statusString argument value.
			StatusString string
		}
		// UnassignFromBatch holds details about calls to the UnassignFromBatch method.
		UnassignFromBatch []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID uuid.UUID
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Instance is the instance argument value.
			Instance *migration.Instance
		}
		// UpdateOverrides holds details about calls to the UpdateOverrides method.
		UpdateOverrides []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Overrides is the overrides argument value.
			Overrides *migration.InstanceOverride
		}
		// UpdateStatusByUUID holds details about calls to the UpdateStatusByUUID method.
		UpdateStatusByUUID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// I is the i argument value.
			I uuid.UUID
			// Status is the status argument value.
			Status api.MigrationStatusType
			// StatusString is the statusString argument value.
			StatusString string
			// NeedsDiskImport is the needsDiskImport argument value.
			NeedsDiskImport bool
		}
	}
	lockCreate                sync.RWMutex
	lockCreateOverrides       sync.RWMutex
	lockDeleteByUUID          sync.RWMutex
	lockDeleteOverridesByUUID sync.RWMutex
	lockGetAll                sync.RWMutex
	lockGetAllByBatch         sync.RWMutex
	lockGetAllByBatchAndState sync.RWMutex
	lockGetAllBySource        sync.RWMutex
	lockGetAllByState         sync.RWMutex
	lockGetAllUUIDs           sync.RWMutex
	lockGetAllUnassigned      sync.RWMutex
	lockGetByUUID             sync.RWMutex
	lockGetOverridesByUUID    sync.RWMutex
	lockProcessWorkerUpdate   sync.RWMutex
	lockUnassignFromBatch     sync.RWMutex
	lockUpdate                sync.RWMutex
	lockUpdateOverrides       sync.RWMutex
	lockUpdateStatusByUUID    sync.RWMutex
}

// Create calls CreateFunc.
func (mock *InstanceServiceMock) Create(ctx context.Context, instance migration.Instance) (migration.Instance, error) {
	if mock.CreateFunc == nil {
		panic("InstanceServiceMock.CreateFunc: method is nil but InstanceService.Create was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Instance migration.Instance
	}{
		Ctx:      ctx,
		Instance: instance,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, instance)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedInstanceService.CreateCalls())
func (mock *InstanceServiceMock) CreateCalls() []struct {
	Ctx      context.Context
	Instance migration.Instance
} {
	var calls []struct {
		Ctx      context.Context
		Instance migration.Instance
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// CreateOverrides calls CreateOverridesFunc.
func (mock *InstanceServiceMock) CreateOverrides(ctx context.Context, overrides migration.InstanceOverride) (migration.InstanceOverride, error) {
	if mock.CreateOverridesFunc == nil {
		panic("InstanceServiceMock.CreateOverridesFunc: method is nil but InstanceService.CreateOverrides was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Overrides migration.InstanceOverride
	}{
		Ctx:       ctx,
		Overrides: overrides,
	}
	mock.lockCreateOverrides.Lock()
	mock.calls.CreateOverrides = append(mock.calls.CreateOverrides, callInfo)
	mock.lockCreateOverrides.Unlock()
	return mock.CreateOverridesFunc(ctx, overrides)
}

// CreateOverridesCalls gets all the calls that were made to CreateOverrides.
// Check the length with:
//
//	len(mockedInstanceService.CreateOverridesCalls())
func (mock *InstanceServiceMock) CreateOverridesCalls() []struct {
	Ctx       context.Context
	Overrides migration.InstanceOverride
} {
	var calls []struct {
		Ctx       context.Context
		Overrides migration.InstanceOverride
	}
	mock.lockCreateOverrides.RLock()
	calls = mock.calls.CreateOverrides
	mock.lockCreateOverrides.RUnlock()
	return calls
}

// DeleteByUUID calls DeleteByUUIDFunc.
func (mock *InstanceServiceMock) DeleteByUUID(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteByUUIDFunc == nil {
		panic("InstanceServiceMock.DeleteByUUIDFunc: method is nil but InstanceService.DeleteByUUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteByUUID.Lock()
	mock.calls.DeleteByUUID = append(mock.calls.DeleteByUUID, callInfo)
	mock.lockDeleteByUUID.Unlock()
	return mock.DeleteByUUIDFunc(ctx, id)
}

// DeleteByUUIDCalls gets all the calls that were made to DeleteByUUID.
// Check the length with:
//
//	len(mockedInstanceService.DeleteByUUIDCalls())
func (mock *InstanceServiceMock) DeleteByUUIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDeleteByUUID.RLock()
	calls = mock.calls.DeleteByUUID
	mock.lockDeleteByUUID.RUnlock()
	return calls
}

// DeleteOverridesByUUID calls DeleteOverridesByUUIDFunc.
func (mock *InstanceServiceMock) DeleteOverridesByUUID(ctx context.Context, id uuid.UUID) error {
	if mock.DeleteOverridesByUUIDFunc == nil {
		panic("InstanceServiceMock.DeleteOverridesByUUIDFunc: method is nil but InstanceService.DeleteOverridesByUUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDeleteOverridesByUUID.Lock()
	mock.calls.DeleteOverridesByUUID = append(mock.calls.DeleteOverridesByUUID, callInfo)
	mock.lockDeleteOverridesByUUID.Unlock()
	return mock.DeleteOverridesByUUIDFunc(ctx, id)
}

// DeleteOverridesByUUIDCalls gets all the calls that were made to DeleteOverridesByUUID.
// Check the length with:
//
//	len(mockedInstanceService.DeleteOverridesByUUIDCalls())
func (mock *InstanceServiceMock) DeleteOverridesByUUIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockDeleteOverridesByUUID.RLock()
	calls = mock.calls.DeleteOverridesByUUID
	mock.lockDeleteOverridesByUUID.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *InstanceServiceMock) GetAll(ctx context.Context, withOverrides bool) (migration.Instances, error) {
	if mock.GetAllFunc == nil {
		panic("InstanceServiceMock.GetAllFunc: method is nil but InstanceService.GetAll was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		WithOverrides bool
	}{
		Ctx:           ctx,
		WithOverrides: withOverrides,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx, withOverrides)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedInstanceService.GetAllCalls())
func (mock *InstanceServiceMock) GetAllCalls() []struct {
	Ctx           context.Context
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		WithOverrides bool
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetAllByBatch calls GetAllByBatchFunc.
func (mock *InstanceServiceMock) GetAllByBatch(ctx context.Context, batch string, withOverrides bool) (migration.Instances, error) {
	if mock.GetAllByBatchFunc == nil {
		panic("InstanceServiceMock.GetAllByBatchFunc: method is nil but InstanceService.GetAllByBatch was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		Batch         string
		WithOverrides bool
	}{
		Ctx:           ctx,
		Batch:         batch,
		WithOverrides: withOverrides,
	}
	mock.lockGetAllByBatch.Lock()
	mock.calls.GetAllByBatch = append(mock.calls.GetAllByBatch, callInfo)
	mock.lockGetAllByBatch.Unlock()
	return mock.GetAllByBatchFunc(ctx, batch, withOverrides)
}

// GetAllByBatchCalls gets all the calls that were made to GetAllByBatch.
// Check the length with:
//
//	len(mockedInstanceService.GetAllByBatchCalls())
func (mock *InstanceServiceMock) GetAllByBatchCalls() []struct {
	Ctx           context.Context
	Batch         string
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		Batch         string
		WithOverrides bool
	}
	mock.lockGetAllByBatch.RLock()
	calls = mock.calls.GetAllByBatch
	mock.lockGetAllByBatch.RUnlock()
	return calls
}

// GetAllByBatchAndState calls GetAllByBatchAndStateFunc.
func (mock *InstanceServiceMock) GetAllByBatchAndState(ctx context.Context, batch string, status api.MigrationStatusType, withOverrides bool) (migration.Instances, error) {
	if mock.GetAllByBatchAndStateFunc == nil {
		panic("InstanceServiceMock.GetAllByBatchAndStateFunc: method is nil but InstanceService.GetAllByBatchAndState was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		Batch         string
		Status        api.MigrationStatusType
		WithOverrides bool
	}{
		Ctx:           ctx,
		Batch:         batch,
		Status:        status,
		WithOverrides: withOverrides,
	}
	mock.lockGetAllByBatchAndState.Lock()
	mock.calls.GetAllByBatchAndState = append(mock.calls.GetAllByBatchAndState, callInfo)
	mock.lockGetAllByBatchAndState.Unlock()
	return mock.GetAllByBatchAndStateFunc(ctx, batch, status, withOverrides)
}

// GetAllByBatchAndStateCalls gets all the calls that were made to GetAllByBatchAndState.
// Check the length with:
//
//	len(mockedInstanceService.GetAllByBatchAndStateCalls())
func (mock *InstanceServiceMock) GetAllByBatchAndStateCalls() []struct {
	Ctx           context.Context
	Batch         string
	Status        api.MigrationStatusType
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		Batch         string
		Status        api.MigrationStatusType
		WithOverrides bool
	}
	mock.lockGetAllByBatchAndState.RLock()
	calls = mock.calls.GetAllByBatchAndState
	mock.lockGetAllByBatchAndState.RUnlock()
	return calls
}

// GetAllBySource calls GetAllBySourceFunc.
func (mock *InstanceServiceMock) GetAllBySource(ctx context.Context, source string, withOverrides bool) (migration.Instances, error) {
	if mock.GetAllBySourceFunc == nil {
		panic("InstanceServiceMock.GetAllBySourceFunc: method is nil but InstanceService.GetAllBySource was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		Source        string
		WithOverrides bool
	}{
		Ctx:           ctx,
		Source:        source,
		WithOverrides: withOverrides,
	}
	mock.lockGetAllBySource.Lock()
	mock.calls.GetAllBySource = append(mock.calls.GetAllBySource, callInfo)
	mock.lockGetAllBySource.Unlock()
	return mock.GetAllBySourceFunc(ctx, source, withOverrides)
}

// GetAllBySourceCalls gets all the calls that were made to GetAllBySource.
// Check the length with:
//
//	len(mockedInstanceService.GetAllBySourceCalls())
func (mock *InstanceServiceMock) GetAllBySourceCalls() []struct {
	Ctx           context.Context
	Source        string
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		Source        string
		WithOverrides bool
	}
	mock.lockGetAllBySource.RLock()
	calls = mock.calls.GetAllBySource
	mock.lockGetAllBySource.RUnlock()
	return calls
}

// GetAllByState calls GetAllByStateFunc.
func (mock *InstanceServiceMock) GetAllByState(ctx context.Context, status api.MigrationStatusType, withOverrides bool) (migration.Instances, error) {
	if mock.GetAllByStateFunc == nil {
		panic("InstanceServiceMock.GetAllByStateFunc: method is nil but InstanceService.GetAllByState was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		Status        api.MigrationStatusType
		WithOverrides bool
	}{
		Ctx:           ctx,
		Status:        status,
		WithOverrides: withOverrides,
	}
	mock.lockGetAllByState.Lock()
	mock.calls.GetAllByState = append(mock.calls.GetAllByState, callInfo)
	mock.lockGetAllByState.Unlock()
	return mock.GetAllByStateFunc(ctx, status, withOverrides)
}

// GetAllByStateCalls gets all the calls that were made to GetAllByState.
// Check the length with:
//
//	len(mockedInstanceService.GetAllByStateCalls())
func (mock *InstanceServiceMock) GetAllByStateCalls() []struct {
	Ctx           context.Context
	Status        api.MigrationStatusType
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		Status        api.MigrationStatusType
		WithOverrides bool
	}
	mock.lockGetAllByState.RLock()
	calls = mock.calls.GetAllByState
	mock.lockGetAllByState.RUnlock()
	return calls
}

// GetAllUUIDs calls GetAllUUIDsFunc.
func (mock *InstanceServiceMock) GetAllUUIDs(ctx context.Context) ([]uuid.UUID, error) {
	if mock.GetAllUUIDsFunc == nil {
		panic("InstanceServiceMock.GetAllUUIDsFunc: method is nil but InstanceService.GetAllUUIDs was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllUUIDs.Lock()
	mock.calls.GetAllUUIDs = append(mock.calls.GetAllUUIDs, callInfo)
	mock.lockGetAllUUIDs.Unlock()
	return mock.GetAllUUIDsFunc(ctx)
}

// GetAllUUIDsCalls gets all the calls that were made to GetAllUUIDs.
// Check the length with:
//
//	len(mockedInstanceService.GetAllUUIDsCalls())
func (mock *InstanceServiceMock) GetAllUUIDsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllUUIDs.RLock()
	calls = mock.calls.GetAllUUIDs
	mock.lockGetAllUUIDs.RUnlock()
	return calls
}

// GetAllUnassigned calls GetAllUnassignedFunc.
func (mock *InstanceServiceMock) GetAllUnassigned(ctx context.Context, withOverrides bool) (migration.Instances, error) {
	if mock.GetAllUnassignedFunc == nil {
		panic("InstanceServiceMock.GetAllUnassignedFunc: method is nil but InstanceService.GetAllUnassigned was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		WithOverrides bool
	}{
		Ctx:           ctx,
		WithOverrides: withOverrides,
	}
	mock.lockGetAllUnassigned.Lock()
	mock.calls.GetAllUnassigned = append(mock.calls.GetAllUnassigned, callInfo)
	mock.lockGetAllUnassigned.Unlock()
	return mock.GetAllUnassignedFunc(ctx, withOverrides)
}

// GetAllUnassignedCalls gets all the calls that were made to GetAllUnassigned.
// Check the length with:
//
//	len(mockedInstanceService.GetAllUnassignedCalls())
func (mock *InstanceServiceMock) GetAllUnassignedCalls() []struct {
	Ctx           context.Context
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		WithOverrides bool
	}
	mock.lockGetAllUnassigned.RLock()
	calls = mock.calls.GetAllUnassigned
	mock.lockGetAllUnassigned.RUnlock()
	return calls
}

// GetByUUID calls GetByUUIDFunc.
func (mock *InstanceServiceMock) GetByUUID(ctx context.Context, id uuid.UUID, withOverrides bool) (*migration.Instance, error) {
	if mock.GetByUUIDFunc == nil {
		panic("InstanceServiceMock.GetByUUIDFunc: method is nil but InstanceService.GetByUUID was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		ID            uuid.UUID
		WithOverrides bool
	}{
		Ctx:           ctx,
		ID:            id,
		WithOverrides: withOverrides,
	}
	mock.lockGetByUUID.Lock()
	mock.calls.GetByUUID = append(mock.calls.GetByUUID, callInfo)
	mock.lockGetByUUID.Unlock()
	return mock.GetByUUIDFunc(ctx, id, withOverrides)
}

// GetByUUIDCalls gets all the calls that were made to GetByUUID.
// Check the length with:
//
//	len(mockedInstanceService.GetByUUIDCalls())
func (mock *InstanceServiceMock) GetByUUIDCalls() []struct {
	Ctx           context.Context
	ID            uuid.UUID
	WithOverrides bool
} {
	var calls []struct {
		Ctx           context.Context
		ID            uuid.UUID
		WithOverrides bool
	}
	mock.lockGetByUUID.RLock()
	calls = mock.calls.GetByUUID
	mock.lockGetByUUID.RUnlock()
	return calls
}

// GetOverridesByUUID calls GetOverridesByUUIDFunc.
func (mock *InstanceServiceMock) GetOverridesByUUID(ctx context.Context, id uuid.UUID) (*migration.InstanceOverride, error) {
	if mock.GetOverridesByUUIDFunc == nil {
		panic("InstanceServiceMock.GetOverridesByUUIDFunc: method is nil but InstanceService.GetOverridesByUUID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetOverridesByUUID.Lock()
	mock.calls.GetOverridesByUUID = append(mock.calls.GetOverridesByUUID, callInfo)
	mock.lockGetOverridesByUUID.Unlock()
	return mock.GetOverridesByUUIDFunc(ctx, id)
}

// GetOverridesByUUIDCalls gets all the calls that were made to GetOverridesByUUID.
// Check the length with:
//
//	len(mockedInstanceService.GetOverridesByUUIDCalls())
func (mock *InstanceServiceMock) GetOverridesByUUIDCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockGetOverridesByUUID.RLock()
	calls = mock.calls.GetOverridesByUUID
	mock.lockGetOverridesByUUID.RUnlock()
	return calls
}

// ProcessWorkerUpdate calls ProcessWorkerUpdateFunc.
func (mock *InstanceServiceMock) ProcessWorkerUpdate(ctx context.Context, id uuid.UUID, workerResponseTypeArg api.WorkerResponseType, statusString string) (migration.Instance, error) {
	if mock.ProcessWorkerUpdateFunc == nil {
		panic("InstanceServiceMock.ProcessWorkerUpdateFunc: method is nil but InstanceService.ProcessWorkerUpdate was just called")
	}
	callInfo := struct {
		Ctx                   context.Context
		ID                    uuid.UUID
		WorkerResponseTypeArg api.WorkerResponseType
		StatusString          string
	}{
		Ctx:                   ctx,
		ID:                    id,
		WorkerResponseTypeArg: workerResponseTypeArg,
		StatusString:          statusString,
	}
	mock.lockProcessWorkerUpdate.Lock()
	mock.calls.ProcessWorkerUpdate = append(mock.calls.ProcessWorkerUpdate, callInfo)
	mock.lockProcessWorkerUpdate.Unlock()
	return mock.ProcessWorkerUpdateFunc(ctx, id, workerResponseTypeArg, statusString)
}

// ProcessWorkerUpdateCalls gets all the calls that were made to ProcessWorkerUpdate.
// Check the length with:
//
//	len(mockedInstanceService.ProcessWorkerUpdateCalls())
func (mock *InstanceServiceMock) ProcessWorkerUpdateCalls() []struct {
	Ctx                   context.Context
	ID                    uuid.UUID
	WorkerResponseTypeArg api.WorkerResponseType
	StatusString          string
} {
	var calls []struct {
		Ctx                   context.Context
		ID                    uuid.UUID
		WorkerResponseTypeArg api.WorkerResponseType
		StatusString          string
	}
	mock.lockProcessWorkerUpdate.RLock()
	calls = mock.calls.ProcessWorkerUpdate
	mock.lockProcessWorkerUpdate.RUnlock()
	return calls
}

// UnassignFromBatch calls UnassignFromBatchFunc.
func (mock *InstanceServiceMock) UnassignFromBatch(ctx context.Context, id uuid.UUID) error {
	if mock.UnassignFromBatchFunc == nil {
		panic("InstanceServiceMock.UnassignFromBatchFunc: method is nil but InstanceService.UnassignFromBatch was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  uuid.UUID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockUnassignFromBatch.Lock()
	mock.calls.UnassignFromBatch = append(mock.calls.UnassignFromBatch, callInfo)
	mock.lockUnassignFromBatch.Unlock()
	return mock.UnassignFromBatchFunc(ctx, id)
}

// UnassignFromBatchCalls gets all the calls that were made to UnassignFromBatch.
// Check the length with:
//
//	len(mockedInstanceService.UnassignFromBatchCalls())
func (mock *InstanceServiceMock) UnassignFromBatchCalls() []struct {
	Ctx context.Context
	ID  uuid.UUID
} {
	var calls []struct {
		Ctx context.Context
		ID  uuid.UUID
	}
	mock.lockUnassignFromBatch.RLock()
	calls = mock.calls.UnassignFromBatch
	mock.lockUnassignFromBatch.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *InstanceServiceMock) Update(ctx context.Context, instance *migration.Instance) error {
	if mock.UpdateFunc == nil {
		panic("InstanceServiceMock.UpdateFunc: method is nil but InstanceService.Update was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Instance *migration.Instance
	}{
		Ctx:      ctx,
		Instance: instance,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, instance)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedInstanceService.UpdateCalls())
func (mock *InstanceServiceMock) UpdateCalls() []struct {
	Ctx      context.Context
	Instance *migration.Instance
} {
	var calls []struct {
		Ctx      context.Context
		Instance *migration.Instance
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}

// UpdateOverrides calls UpdateOverridesFunc.
func (mock *InstanceServiceMock) UpdateOverrides(ctx context.Context, overrides *migration.InstanceOverride) error {
	if mock.UpdateOverridesFunc == nil {
		panic("InstanceServiceMock.UpdateOverridesFunc: method is nil but InstanceService.UpdateOverrides was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Overrides *migration.InstanceOverride
	}{
		Ctx:       ctx,
		Overrides: overrides,
	}
	mock.lockUpdateOverrides.Lock()
	mock.calls.UpdateOverrides = append(mock.calls.UpdateOverrides, callInfo)
	mock.lockUpdateOverrides.Unlock()
	return mock.UpdateOverridesFunc(ctx, overrides)
}

// UpdateOverridesCalls gets all the calls that were made to UpdateOverrides.
// Check the length with:
//
//	len(mockedInstanceService.UpdateOverridesCalls())
func (mock *InstanceServiceMock) UpdateOverridesCalls() []struct {
	Ctx       context.Context
	Overrides *migration.InstanceOverride
} {
	var calls []struct {
		Ctx       context.Context
		Overrides *migration.InstanceOverride
	}
	mock.lockUpdateOverrides.RLock()
	calls = mock.calls.UpdateOverrides
	mock.lockUpdateOverrides.RUnlock()
	return calls
}

// UpdateStatusByUUID calls UpdateStatusByUUIDFunc.
func (mock *InstanceServiceMock) UpdateStatusByUUID(ctx context.Context, i uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (*migration.Instance, error) {
	if mock.UpdateStatusByUUIDFunc == nil {
		panic("InstanceServiceMock.UpdateStatusByUUIDFunc: method is nil but InstanceService.UpdateStatusByUUID was just called")
	}
	callInfo := struct {
		Ctx             context.Context
		I               uuid.UUID
		Status          api.MigrationStatusType
		StatusString    string
		NeedsDiskImport bool
	}{
		Ctx:             ctx,
		I:               i,
		Status:          status,
		StatusString:    statusString,
		NeedsDiskImport: needsDiskImport,
	}
	mock.lockUpdateStatusByUUID.Lock()
	mock.calls.UpdateStatusByUUID = append(mock.calls.UpdateStatusByUUID, callInfo)
	mock.lockUpdateStatusByUUID.Unlock()
	return mock.UpdateStatusByUUIDFunc(ctx, i, status, statusString, needsDiskImport)
}

// UpdateStatusByUUIDCalls gets all the calls that were made to UpdateStatusByUUID.
// Check the length with:
//
//	len(mockedInstanceService.UpdateStatusByUUIDCalls())
func (mock *InstanceServiceMock) UpdateStatusByUUIDCalls() []struct {
	Ctx             context.Context
	I               uuid.UUID
	Status          api.MigrationStatusType
	StatusString    string
	NeedsDiskImport bool
} {
	var calls []struct {
		Ctx             context.Context
		I               uuid.UUID
		Status          api.MigrationStatusType
		StatusString    string
		NeedsDiskImport bool
	}
	mock.lockUpdateStatusByUUID.RLock()
	calls = mock.calls.UpdateStatusByUUID
	mock.lockUpdateStatusByUUID.RUnlock()
	return calls
}
