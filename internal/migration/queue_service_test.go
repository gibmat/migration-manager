package migration_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/internal/ptr"
	"github.com/FuturFusion/migration-manager/internal/testing/boom"
	"github.com/FuturFusion/migration-manager/internal/testing/queue"
	"github.com/FuturFusion/migration-manager/shared/api"
)

func TestQueueService_GetAll(t *testing.T) {
	tests := []struct {
		name                       string
		batchSvcGetAllBatches      migration.Batches
		batchSvcGetAllErr          error
		instanceSvcGetAllByBatchID []queue.Item[migration.Instances]

		assertErr      require.ErrorAssertionFunc
		wantQueueItems migration.QueueEntries
	}{
		{
			name:                  "success - no batches",
			batchSvcGetAllBatches: nil,

			assertErr: require.NoError,
		},
		{
			name: "success - with batches",
			batchSvcGetAllBatches: migration.Batches{
				{
					ID:     1,
					Name:   "one",
					Status: api.BATCHSTATUS_UNKNOWN, // this batch is ignored
				},
				{
					ID:     2,
					Name:   "two",
					Status: api.BATCHSTATUS_DEFINED, // this batch is ignored
				},
				{
					ID:     3,
					Name:   "three",
					Status: api.BATCHSTATUS_QUEUED,
				},
				{
					ID:     4,
					Name:   "four",
					Status: api.BATCHSTATUS_RUNNING,
				},
				{
					ID:     5,
					Name:   "five",
					Status: api.BATCHSTATUS_RUNNING,
				},
			},
			instanceSvcGetAllByBatchID: []queue.Item[migration.Instances]{
				// Instances for batch 3
				{
					Value: migration.Instances{
						{
							UUID:                  uuidA,
							InventoryPath:         "/some/instance/A",
							MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
							MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
							Batch:                 ptr.To("three"),
						},
					},
				},
				// Instances for batch 4
				{
					Value: migration.Instances{
						{
							UUID:                  uuidB,
							InventoryPath:         "/some/instance/B",
							MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
							MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
							Batch:                 ptr.To("four"),
						},
					},
				},
				// Instances for batch 5
				{
					Value: migration.Instances{
						{
							UUID:                  uuidC,
							InventoryPath:         "/some/instance/C",
							MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
							MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
							Batch:                 ptr.To("five"),
						},
					},
				},
			},

			assertErr: require.NoError,
			wantQueueItems: migration.QueueEntries{
				{
					InstanceUUID:          uuidA,
					InstanceName:          "A",
					MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
					MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
					BatchName:             "three",
				},
				{
					InstanceUUID:          uuidB,
					InstanceName:          "B",
					MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
					MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
					BatchName:             "four",
				},
				{
					InstanceUUID:          uuidC,
					InstanceName:          "C",
					MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
					MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
					BatchName:             "five",
				},
			},
		},
		{
			name:              "error - batch.GetAll",
			batchSvcGetAllErr: boom.Error,

			assertErr: boom.ErrorIs,
		},
		{
			name: "error - instance.GetAllByBatchID",
			batchSvcGetAllBatches: migration.Batches{
				{
					ID:     1,
					Name:   "one",
					Status: api.BATCHSTATUS_RUNNING,
				},
			},
			instanceSvcGetAllByBatchID: []queue.Item[migration.Instances]{
				{
					Err: boom.Error,
				},
			},

			assertErr: boom.ErrorIs,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			batchSvc := &BatchServiceMock{
				GetAllFunc: func(ctx context.Context) (migration.Batches, error) {
					return tc.batchSvcGetAllBatches, tc.batchSvcGetAllErr
				},
			}

			instanceSvc := &InstanceServiceMock{
				GetAllByBatchFunc: func(ctx context.Context, batch string, withOverrides bool) (migration.Instances, error) {
					return queue.Pop(t, &tc.instanceSvcGetAllByBatchID)
				},
			}

			queueSvc := migration.NewQueueService(batchSvc, instanceSvc, nil)

			// Run test
			queueItems, err := queueSvc.GetAll(context.Background())

			// Assert
			tc.assertErr(t, err)
			require.Equal(t, tc.wantQueueItems, queueItems)

			// Ensure queues are completely drained.
			require.Empty(t, tc.instanceSvcGetAllByBatchID)
		})
	}
}

func TestInstanceService_GetByInstanceID(t *testing.T) {
	tests := []struct {
		name                       string
		uuidArg                    uuid.UUID
		instanceSvcGetByIDInstance migration.Instance
		instanceSvcGetByIDErr      error
		batchSvcGetByIDBatch       migration.Batch
		batchSvcGetByIDErr         error

		assertErr      require.ErrorAssertionFunc
		wantQueueEntry migration.QueueEntry
	}{
		{
			name:    "success",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
				MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
				Batch:                 ptr.To("one"),
			},
			batchSvcGetByIDBatch: migration.Batch{
				ID:   1,
				Name: "one",
			},

			assertErr: require.NoError,
			wantQueueEntry: migration.QueueEntry{
				InstanceUUID:          uuidA,
				InstanceName:          "A",
				MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
				MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
				BatchName:             "one",
			},
		},
		{
			name:                  "error - instance.GetByID",
			uuidArg:               uuidA,
			instanceSvcGetByIDErr: boom.Error,

			assertErr: boom.ErrorIs,
		},
		{
			name:    "error - instance not assigned to batch",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
				MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
				Batch:                 nil, // not assigned to batch
			},

			assertErr: func(tt require.TestingT, err error, a ...any) {
				require.ErrorIs(tt, err, migration.ErrNotFound, a...)
			},
		},
		{
			name:    "error - instance not in migration state",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH, // not in migration state
				MigrationStatusString: api.MIGRATIONSTATUS_NOT_ASSIGNED_BATCH.String(),
				Batch:                 ptr.To("one"),
			},

			assertErr: func(tt require.TestingT, err error, a ...any) {
				require.ErrorIs(tt, err, migration.ErrNotFound, a...)
			},
		},
		{
			name:    "error - batch.GetByID",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_CREATING,
				MigrationStatusString: api.MIGRATIONSTATUS_CREATING.String(),
				Batch:                 ptr.To("one"),
			},
			batchSvcGetByIDErr: boom.Error,

			assertErr: boom.ErrorIs,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			instanceSvc := &InstanceServiceMock{
				GetByUUIDFunc: func(ctx context.Context, id uuid.UUID, withOverrides bool) (*migration.Instance, error) {
					return &tc.instanceSvcGetByIDInstance, tc.instanceSvcGetByIDErr
				},
			}

			batchSvc := &BatchServiceMock{
				GetByNameFunc: func(ctx context.Context, name string) (*migration.Batch, error) {
					return &tc.batchSvcGetByIDBatch, tc.batchSvcGetByIDErr
				},
			}

			queueSvc := migration.NewQueueService(batchSvc, instanceSvc, nil)

			// Run test
			queueEntry, err := queueSvc.GetByInstanceID(context.Background(), tc.uuidArg)

			// Assert
			tc.assertErr(t, err)
			require.Equal(t, tc.wantQueueEntry, queueEntry)
		})
	}
}

func TestInstanceService_NewWorkerCommandByInstanceUUID(t *testing.T) {
	tests := []struct {
		name                           string
		uuidArg                        uuid.UUID
		instanceSvcGetByIDInstance     migration.Instance
		instanceSvcGetByIDErr          error
		sourceSvcGetByIDSource         migration.Source
		sourceSvcGetByIDErr            error
		batchSvcGetByIDBatch           migration.Batch
		batchSvcGetByIDErr             error
		instanceSvcUpdateStatusByIDErr error

		assertErr                 require.ErrorAssertionFunc
		wantMigrationStatus       api.MigrationStatusType
		wantMigrationStatusString string
		wantWorkerCommand         migration.WorkerCommand
	}{
		{
			name:    "success - without migration window start time",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},
			sourceSvcGetByIDSource: migration.Source{
				ID:         1,
				Name:       "one",
				SourceType: api.SOURCETYPE_VMWARE,
			},
			batchSvcGetByIDBatch: migration.Batch{
				ID:   2,
				Name: "two",
			},

			assertErr: require.NoError,
			wantWorkerCommand: migration.WorkerCommand{
				Command:       api.WORKERCOMMAND_FINALIZE_IMPORT,
				InventoryPath: "/some/instance/A",
				SourceType:    api.SOURCETYPE_VMWARE,
				Source: migration.Source{
					ID:         1,
					Name:       "one",
					SourceType: api.SOURCETYPE_VMWARE,
				},
				OS:        "ubuntu",
				OSVersion: "24.04",
			},
			wantMigrationStatus:       api.MIGRATIONSTATUS_FINAL_IMPORT,
			wantMigrationStatusString: api.MIGRATIONSTATUS_FINAL_IMPORT.String(),
		},
		{
			name:    "success - background disk sync",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: true,
					},
				},
				NeedsDiskImport: true,
			},
			sourceSvcGetByIDSource: migration.Source{
				ID:         1,
				Name:       "one",
				SourceType: api.SOURCETYPE_VMWARE,
			},
			batchSvcGetByIDBatch: migration.Batch{
				ID:   2,
				Name: "two",
			},

			assertErr: require.NoError,
			wantWorkerCommand: migration.WorkerCommand{
				Command:       api.WORKERCOMMAND_IMPORT_DISKS,
				InventoryPath: "/some/instance/A",
				SourceType:    api.SOURCETYPE_VMWARE,
				Source: migration.Source{
					ID:         1,
					Name:       "one",
					SourceType: api.SOURCETYPE_VMWARE,
				},
				OS:        "ubuntu",
				OSVersion: "24.04",
			},
			wantMigrationStatus:       api.MIGRATIONSTATUS_BACKGROUND_IMPORT,
			wantMigrationStatusString: api.MIGRATIONSTATUS_BACKGROUND_IMPORT.String(),
		},
		{
			name:    "success - migration window started",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},
			sourceSvcGetByIDSource: migration.Source{
				ID:         1,
				Name:       "one",
				SourceType: api.SOURCETYPE_VMWARE,
			},
			batchSvcGetByIDBatch: migration.Batch{
				ID:                   2,
				Name:                 "two",
				MigrationWindowStart: time.Now().Add(-1 * time.Minute),
			},

			assertErr: require.NoError,
			wantWorkerCommand: migration.WorkerCommand{
				Command:       api.WORKERCOMMAND_FINALIZE_IMPORT,
				InventoryPath: "/some/instance/A",
				SourceType:    api.SOURCETYPE_VMWARE,
				Source: migration.Source{
					ID:         1,
					Name:       "one",
					SourceType: api.SOURCETYPE_VMWARE,
				},
				OS:        "ubuntu",
				OSVersion: "24.04",
			},
			wantMigrationStatus:       api.MIGRATIONSTATUS_FINAL_IMPORT,
			wantMigrationStatusString: api.MIGRATIONSTATUS_FINAL_IMPORT.String(),
		},
		{
			name:                  "error - instance.GetByID",
			uuidArg:               uuidA,
			instanceSvcGetByIDErr: boom.Error,

			assertErr: boom.ErrorIs,
		},
		{
			name:    "error - not assigned to batch",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 nil, // not assigned to batch
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},

			assertErr: func(tt require.TestingT, err error, a ...any) {
				require.ErrorIs(tt, err, migration.ErrNotFound, a...)
			},
		},
		{
			name:    "error - instance is not in idle state",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_ASSIGNED_BATCH,
				MigrationStatusString: api.MIGRATIONSTATUS_ASSIGNED_BATCH.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},

			assertErr: func(tt require.TestingT, err error, a ...any) {
				require.ErrorIs(tt, err, migration.ErrOperationNotPermitted, a...)
			},
		},
		{
			name:    "error - source.GetByID",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},
			sourceSvcGetByIDErr: boom.Error,

			assertErr: boom.ErrorIs,
		},
		{
			name:    "error - batch.GetByID",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},
			sourceSvcGetByIDSource: migration.Source{
				ID:         1,
				Name:       "one",
				SourceType: api.SOURCETYPE_VMWARE,
			},
			batchSvcGetByIDErr: boom.Error,

			assertErr: boom.ErrorIs,
		},
		{
			name:    "error - instance.UpdateStatusByID",
			uuidArg: uuidA,
			instanceSvcGetByIDInstance: migration.Instance{
				UUID:                  uuidA,
				InventoryPath:         "/some/instance/A",
				MigrationStatus:       api.MIGRATIONSTATUS_IDLE,
				MigrationStatusString: api.MIGRATIONSTATUS_IDLE.String(),
				Batch:                 ptr.To("one"),
				OS:                    "ubuntu",
				OSVersion:             "24.04",
				Disks: []api.InstanceDiskInfo{
					{
						Type:                      "HDD",
						DifferentialSyncSupported: false,
					},
				},
				NeedsDiskImport: true,
			},
			sourceSvcGetByIDSource: migration.Source{
				ID:         1,
				Name:       "one",
				SourceType: api.SOURCETYPE_VMWARE,
			},
			batchSvcGetByIDBatch: migration.Batch{
				ID:   2,
				Name: "two",
			},
			instanceSvcUpdateStatusByIDErr: boom.Error,

			assertErr:                 boom.ErrorIs,
			wantMigrationStatus:       api.MIGRATIONSTATUS_FINAL_IMPORT,
			wantMigrationStatusString: api.MIGRATIONSTATUS_FINAL_IMPORT.String(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			instanceSvc := &InstanceServiceMock{
				GetByUUIDFunc: func(ctx context.Context, id uuid.UUID, withOverrides bool) (*migration.Instance, error) {
					return &tc.instanceSvcGetByIDInstance, tc.instanceSvcGetByIDErr
				},
				UpdateStatusByUUIDFunc: func(ctx context.Context, id uuid.UUID, status api.MigrationStatusType, statusString string, needsDiskImport bool) (*migration.Instance, error) {
					require.Equal(t, tc.wantMigrationStatus, status)
					require.Equal(t, tc.wantMigrationStatusString, statusString)
					return nil, tc.instanceSvcUpdateStatusByIDErr
				},
			}

			sourceSvc := &SourceServiceMock{
				GetByNameFunc: func(ctx context.Context, name string) (*migration.Source, error) {
					return &tc.sourceSvcGetByIDSource, tc.sourceSvcGetByIDErr
				},
			}

			batchSvc := &BatchServiceMock{
				GetByNameFunc: func(ctx context.Context, name string) (*migration.Batch, error) {
					return &tc.batchSvcGetByIDBatch, tc.batchSvcGetByIDErr
				},
			}

			queueSvc := migration.NewQueueService(batchSvc, instanceSvc, sourceSvc)

			// Run test
			workerCommand, err := queueSvc.NewWorkerCommandByInstanceUUID(context.Background(), tc.uuidArg)

			// Assert
			tc.assertErr(t, err)
			require.Equal(t, tc.wantWorkerCommand, workerCommand)
		})
	}
}
