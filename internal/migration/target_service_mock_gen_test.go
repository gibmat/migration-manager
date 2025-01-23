// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package migration_test

import (
	"context"
	"sync"

	"github.com/FuturFusion/migration-manager/internal/migration"
)

// Ensure, that TargetServiceMock does implement migration.TargetService.
// If this is not the case, regenerate this file with moq.
var _ migration.TargetService = &TargetServiceMock{}

// TargetServiceMock is a mock implementation of migration.TargetService.
//
//	func TestSomethingThatUsesTargetService(t *testing.T) {
//
//		// make and configure a mocked migration.TargetService
//		mockedTargetService := &TargetServiceMock{
//			CreateFunc: func(ctx context.Context, target migration.Target) (migration.Target, error) {
//				panic("mock out the Create method")
//			},
//			DeleteByNameFunc: func(ctx context.Context, name string) error {
//				panic("mock out the DeleteByName method")
//			},
//			GetAllFunc: func(ctx context.Context) (migration.Targets, error) {
//				panic("mock out the GetAll method")
//			},
//			GetAllNamesFunc: func(ctx context.Context) ([]string, error) {
//				panic("mock out the GetAllNames method")
//			},
//			GetByIDFunc: func(ctx context.Context, id int) (migration.Target, error) {
//				panic("mock out the GetByID method")
//			},
//			GetByNameFunc: func(ctx context.Context, name string) (migration.Target, error) {
//				panic("mock out the GetByName method")
//			},
//			UpdateByIDFunc: func(ctx context.Context, target migration.Target) (migration.Target, error) {
//				panic("mock out the UpdateByID method")
//			},
//		}
//
//		// use mockedTargetService in code that requires migration.TargetService
//		// and then make assertions.
//
//	}
type TargetServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, target migration.Target) (migration.Target, error)

	// DeleteByNameFunc mocks the DeleteByName method.
	DeleteByNameFunc func(ctx context.Context, name string) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) (migration.Targets, error)

	// GetAllNamesFunc mocks the GetAllNames method.
	GetAllNamesFunc func(ctx context.Context) ([]string, error)

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(ctx context.Context, id int) (migration.Target, error)

	// GetByNameFunc mocks the GetByName method.
	GetByNameFunc func(ctx context.Context, name string) (migration.Target, error)

	// UpdateByIDFunc mocks the UpdateByID method.
	UpdateByIDFunc func(ctx context.Context, target migration.Target) (migration.Target, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Target is the target argument value.
			Target migration.Target
		}
		// DeleteByName holds details about calls to the DeleteByName method.
		DeleteByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllNames holds details about calls to the GetAllNames method.
		GetAllNames []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID int
		}
		// GetByName holds details about calls to the GetByName method.
		GetByName []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// UpdateByID holds details about calls to the UpdateByID method.
		UpdateByID []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Target is the target argument value.
			Target migration.Target
		}
	}
	lockCreate       sync.RWMutex
	lockDeleteByName sync.RWMutex
	lockGetAll       sync.RWMutex
	lockGetAllNames  sync.RWMutex
	lockGetByID      sync.RWMutex
	lockGetByName    sync.RWMutex
	lockUpdateByID   sync.RWMutex
}

// Create calls CreateFunc.
func (mock *TargetServiceMock) Create(ctx context.Context, target migration.Target) (migration.Target, error) {
	if mock.CreateFunc == nil {
		panic("TargetServiceMock.CreateFunc: method is nil but TargetService.Create was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Target migration.Target
	}{
		Ctx:    ctx,
		Target: target,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, target)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedTargetService.CreateCalls())
func (mock *TargetServiceMock) CreateCalls() []struct {
	Ctx    context.Context
	Target migration.Target
} {
	var calls []struct {
		Ctx    context.Context
		Target migration.Target
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DeleteByName calls DeleteByNameFunc.
func (mock *TargetServiceMock) DeleteByName(ctx context.Context, name string) error {
	if mock.DeleteByNameFunc == nil {
		panic("TargetServiceMock.DeleteByNameFunc: method is nil but TargetService.DeleteByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockDeleteByName.Lock()
	mock.calls.DeleteByName = append(mock.calls.DeleteByName, callInfo)
	mock.lockDeleteByName.Unlock()
	return mock.DeleteByNameFunc(ctx, name)
}

// DeleteByNameCalls gets all the calls that were made to DeleteByName.
// Check the length with:
//
//	len(mockedTargetService.DeleteByNameCalls())
func (mock *TargetServiceMock) DeleteByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockDeleteByName.RLock()
	calls = mock.calls.DeleteByName
	mock.lockDeleteByName.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *TargetServiceMock) GetAll(ctx context.Context) (migration.Targets, error) {
	if mock.GetAllFunc == nil {
		panic("TargetServiceMock.GetAllFunc: method is nil but TargetService.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedTargetService.GetAllCalls())
func (mock *TargetServiceMock) GetAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetAllNames calls GetAllNamesFunc.
func (mock *TargetServiceMock) GetAllNames(ctx context.Context) ([]string, error) {
	if mock.GetAllNamesFunc == nil {
		panic("TargetServiceMock.GetAllNamesFunc: method is nil but TargetService.GetAllNames was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAllNames.Lock()
	mock.calls.GetAllNames = append(mock.calls.GetAllNames, callInfo)
	mock.lockGetAllNames.Unlock()
	return mock.GetAllNamesFunc(ctx)
}

// GetAllNamesCalls gets all the calls that were made to GetAllNames.
// Check the length with:
//
//	len(mockedTargetService.GetAllNamesCalls())
func (mock *TargetServiceMock) GetAllNamesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAllNames.RLock()
	calls = mock.calls.GetAllNames
	mock.lockGetAllNames.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *TargetServiceMock) GetByID(ctx context.Context, id int) (migration.Target, error) {
	if mock.GetByIDFunc == nil {
		panic("TargetServiceMock.GetByIDFunc: method is nil but TargetService.GetByID was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  int
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	mock.lockGetByID.Unlock()
	return mock.GetByIDFunc(ctx, id)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//
//	len(mockedTargetService.GetByIDCalls())
func (mock *TargetServiceMock) GetByIDCalls() []struct {
	Ctx context.Context
	ID  int
} {
	var calls []struct {
		Ctx context.Context
		ID  int
	}
	mock.lockGetByID.RLock()
	calls = mock.calls.GetByID
	mock.lockGetByID.RUnlock()
	return calls
}

// GetByName calls GetByNameFunc.
func (mock *TargetServiceMock) GetByName(ctx context.Context, name string) (migration.Target, error) {
	if mock.GetByNameFunc == nil {
		panic("TargetServiceMock.GetByNameFunc: method is nil but TargetService.GetByName was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockGetByName.Lock()
	mock.calls.GetByName = append(mock.calls.GetByName, callInfo)
	mock.lockGetByName.Unlock()
	return mock.GetByNameFunc(ctx, name)
}

// GetByNameCalls gets all the calls that were made to GetByName.
// Check the length with:
//
//	len(mockedTargetService.GetByNameCalls())
func (mock *TargetServiceMock) GetByNameCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockGetByName.RLock()
	calls = mock.calls.GetByName
	mock.lockGetByName.RUnlock()
	return calls
}

// UpdateByID calls UpdateByIDFunc.
func (mock *TargetServiceMock) UpdateByID(ctx context.Context, target migration.Target) (migration.Target, error) {
	if mock.UpdateByIDFunc == nil {
		panic("TargetServiceMock.UpdateByIDFunc: method is nil but TargetService.UpdateByID was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Target migration.Target
	}{
		Ctx:    ctx,
		Target: target,
	}
	mock.lockUpdateByID.Lock()
	mock.calls.UpdateByID = append(mock.calls.UpdateByID, callInfo)
	mock.lockUpdateByID.Unlock()
	return mock.UpdateByIDFunc(ctx, target)
}

// UpdateByIDCalls gets all the calls that were made to UpdateByID.
// Check the length with:
//
//	len(mockedTargetService.UpdateByIDCalls())
func (mock *TargetServiceMock) UpdateByIDCalls() []struct {
	Ctx    context.Context
	Target migration.Target
} {
	var calls []struct {
		Ctx    context.Context
		Target migration.Target
	}
	mock.lockUpdateByID.RLock()
	calls = mock.calls.UpdateByID
	mock.lockUpdateByID.RUnlock()
	return calls
}
