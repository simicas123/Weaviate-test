//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package usage

import (
	"context"
	"sort"
	"testing"
	"time"

	logrus "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/weaviate/weaviate/adapters/repos/db"
	"github.com/weaviate/weaviate/adapters/repos/db/vector/compressionhelpers"
	types "github.com/weaviate/weaviate/cluster/usage/types"
	"github.com/weaviate/weaviate/entities/backup"
	"github.com/weaviate/weaviate/entities/models"
	modulecapabilities "github.com/weaviate/weaviate/entities/modulecapabilities"
	entschema "github.com/weaviate/weaviate/entities/schema"
	"github.com/weaviate/weaviate/entities/vectorindex/hnsw"
	backupusecase "github.com/weaviate/weaviate/usecases/backup"
	"github.com/weaviate/weaviate/usecases/schema"
	"github.com/weaviate/weaviate/usecases/sharding"
)

func TestService_Usage_SingleTenant(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	className := "TestClass"
	replication := 1
	uniqueShards := 1
	shardName := "abcd"
	objectCount := 1000
	storageSize := int64(5000000)
	vectorName := "abcd"
	vectorType := "hnsw"
	compression := "standard"
	compressionRatio := 0.75
	dimensionality := 1536
	dimensionCount := 1000

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:             className,
					VectorIndexConfig: &hnsw.UserConfig{},
					ReplicationConfig: &models.ReplicationConfig{
						Factor: int64(replication),
					},
				},
			},
		},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	shardingState := &sharding.State{
		Physical: map[string]sharding.Physical{
			"": {
				Name:           shardName,
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusHOT,
			},
		},
	}
	shardingState.SetLocalName(nodeName)
	mockSchema.EXPECT().CopyShardingState(className).Return(shardingState)

	mockDB := db.NewMockIndexGetter(t)
	mockIndex := db.NewMockIndexLike(t)
	mockDB.EXPECT().GetIndexLike(entschema.ClassName(className)).Return(mockIndex)

	mockShard := db.NewMockShardLike(t)
	mockShard.EXPECT().ObjectCountAsync(ctx).Return(int64(objectCount), nil)
	mockShard.EXPECT().ObjectStorageSize(ctx).Return(storageSize, nil)
	mockShard.EXPECT().VectorStorageSize(ctx).Return(int64(0), nil)
	mockShard.EXPECT().DimensionsUsage(ctx, vectorName).Return(types.Dimensionality{
		Dimensions: dimensionality,
		Count:      dimensionCount,
	}, nil)

	mockVectorIndex := db.NewMockVectorIndex(t)
	mockCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockCompressionStats.EXPECT().CompressionRatio(dimensionality).Return(compressionRatio)
	mockVectorIndex.EXPECT().CompressionStats().Return(mockCompressionStats)

	mockIndex.On("ForEachShard", mock.AnythingOfType("func(string, db.ShardLike) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.ShardLike) error)
		f("", mockShard)
	})

	mockShard.On("ForEachVectorIndex", mock.AnythingOfType("func(string, db.VectorIndex) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.VectorIndex) error)
		f(vectorName, mockVectorIndex)
	})

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 1)

	collection := result.Collections[0]
	assert.Equal(t, className, collection.Name)
	assert.Equal(t, replication, collection.ReplicationFactor)
	assert.Equal(t, uniqueShards, collection.UniqueShardCount)
	assert.Len(t, collection.Shards, 1)

	shard := collection.Shards[0]
	assert.Equal(t, "", shard.Name)
	assert.Equal(t, int64(objectCount), shard.ObjectsCount)
	assert.Equal(t, uint64(storageSize), shard.ObjectsStorageBytes)
	assert.Len(t, shard.NamedVectors, 1)

	vector := shard.NamedVectors[0]
	assert.Equal(t, vectorName, vector.Name)
	assert.Equal(t, vectorType, vector.VectorIndexType)
	assert.Equal(t, compression, vector.Compression)
	assert.Equal(t, compressionRatio, vector.VectorCompressionRatio)
	assert.Len(t, vector.Dimensionalities, 1)

	dim := vector.Dimensionalities[0]
	assert.Equal(t, dimensionality, dim.Dimensions)
	assert.Equal(t, dimensionCount, dim.Count)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockIndex.AssertExpectations(t)
	mockShard.AssertExpectations(t)
	mockVectorIndex.AssertExpectations(t)
	mockCompressionStats.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}

func TestService_Usage_MultiTenant_HotAndCold(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	className := "MultiTenantClass"
	replication := 3
	uniqueShards := 2
	hotTenant := "tenant1"
	coldTenant := "tenant2"
	hotObjectCount := 1500
	coldObjectCount := 500
	hotStorageSize := int64(7500000)
	coldStorageSize := int64(2500000)
	vectorName := "abcd"
	vectorType := "hnsw"
	compression := "standard"
	compressionRatio := 0.8
	dimensionality := 1536
	dimensionCount := 1500

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:              className,
					VectorIndexConfig:  &hnsw.UserConfig{},
					ReplicationConfig:  &models.ReplicationConfig{Factor: int64(replication)},
					MultiTenancyConfig: &models.MultiTenancyConfig{Enabled: true},
				},
			},
		},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	shardingState := &sharding.State{
		PartitioningEnabled: true,
		Physical: map[string]sharding.Physical{
			hotTenant: {
				Name:           hotTenant,
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusHOT,
			},
			coldTenant: {
				Name:           coldTenant,
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusCOLD,
			},
		},
	}
	shardingState.SetLocalName(nodeName)
	mockSchema.EXPECT().CopyShardingState(className).Return(shardingState)

	mockDB := db.NewMockIndexGetter(t)
	mockIndex := db.NewMockIndexLike(t)
	mockDB.EXPECT().GetIndexLike(entschema.ClassName(className)).Return(mockIndex)

	mockShard := db.NewMockShardLike(t)
	mockShard.EXPECT().ObjectCountAsync(ctx).Return(int64(hotObjectCount), nil)
	mockShard.EXPECT().ObjectStorageSize(ctx).Return(hotStorageSize, nil)
	mockShard.EXPECT().VectorStorageSize(ctx).Return(int64(0), nil)
	mockShard.EXPECT().DimensionsUsage(ctx, vectorName).Return(types.Dimensionality{
		Dimensions: dimensionality,
		Count:      dimensionCount,
	}, nil)

	mockVectorIndex := db.NewMockVectorIndex(t)
	mockCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockCompressionStats.EXPECT().CompressionRatio(dimensionality).Return(compressionRatio)
	mockVectorIndex.EXPECT().CompressionStats().Return(mockCompressionStats)

	mockIndex.On("ForEachShard", mock.AnythingOfType("func(string, db.ShardLike) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.ShardLike) error)
		f(hotTenant, mockShard)
	})
	mockIndex.EXPECT().CalculateUnloadedObjectsMetrics(ctx, coldTenant).Return(types.ObjectUsage{
		Count:        int64(coldObjectCount),
		StorageBytes: coldStorageSize,
	}, nil)
	mockIndex.EXPECT().CalculateUnloadedVectorsMetrics(ctx, coldTenant).Return(int64(0), nil)

	mockShard.On("ForEachVectorIndex", mock.AnythingOfType("func(string, db.VectorIndex) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.VectorIndex) error)
		f(vectorName, mockVectorIndex)
	})

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 1)

	collection := result.Collections[0]
	assert.Equal(t, className, collection.Name)
	assert.Equal(t, replication, collection.ReplicationFactor)
	assert.Equal(t, uniqueShards, collection.UniqueShardCount)
	assert.Len(t, collection.Shards, 2)

	var hotShard, coldShard *types.ShardUsage
	for _, shard := range collection.Shards {
		switch shard.Name {
		case hotTenant:
			hotShard = shard
		case coldTenant:
			coldShard = shard
		}
	}

	require.NotNil(t, hotShard)
	assert.Equal(t, int64(hotObjectCount), hotShard.ObjectsCount)
	assert.Equal(t, uint64(hotStorageSize), hotShard.ObjectsStorageBytes)
	assert.Len(t, hotShard.NamedVectors, 1)

	require.NotNil(t, coldShard)
	assert.Equal(t, int64(coldObjectCount), coldShard.ObjectsCount)
	assert.Equal(t, uint64(coldStorageSize), coldShard.ObjectsStorageBytes)
	assert.Len(t, coldShard.NamedVectors, 0)

	vector := hotShard.NamedVectors[0]
	assert.Equal(t, vectorName, vector.Name)
	assert.Equal(t, vectorType, vector.VectorIndexType)
	assert.Equal(t, compression, vector.Compression)
	assert.Equal(t, compressionRatio, vector.VectorCompressionRatio)
	assert.Len(t, vector.Dimensionalities, 1)
	dim := vector.Dimensionalities[0]
	assert.Equal(t, dimensionality, dim.Dimensions)
	assert.Equal(t, dimensionCount, dim.Count)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockIndex.AssertExpectations(t)
	mockShard.AssertExpectations(t)
	mockVectorIndex.AssertExpectations(t)
	mockCompressionStats.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}

func TestService_Usage_WithBackups(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	backupID := "backup-1"
	backupStatus := backup.Success
	completionTime := time.Date(2025, 1, 15, 10, 30, 0, 0, time.UTC)
	preCompressionSizeBytes := int64(1073741824) // 1 GiB
	sizeInGib := 1.0
	backupType := "SUCCESS"
	class1 := "Class1"
	class2 := "Class2"
	class3 := "Class3"

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{Classes: []*models.Class{}},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	mockDB := db.NewMockIndexGetter(t)

	mockBackupBackend := modulecapabilities.NewMockBackupBackend(t)
	backups := []*backup.DistributedBackupDescriptor{
		{
			ID:                      backupID,
			Status:                  backupStatus,
			CompletedAt:             completionTime,
			PreCompressionSizeBytes: preCompressionSizeBytes,
			Nodes: map[string]*backup.NodeDescriptor{
				"node1": {Classes: []string{class1, class2}},
			},
		},
		{
			ID:                      "backup-2",
			Status:                  backup.Failed,
			CompletedAt:             completionTime,
			PreCompressionSizeBytes: 2147483648,
			Nodes: map[string]*backup.NodeDescriptor{
				"node1": {Classes: []string{class3}},
			},
		},
	}
	mockBackupBackend.EXPECT().AllBackups(ctx).Return(backups, nil)

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{mockBackupBackend})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 0)
	assert.Len(t, result.Backups, 1)

	backup := result.Backups[0]
	assert.Equal(t, backupID, backup.ID)
	assert.Equal(t, completionTime.Format(time.RFC3339), backup.CompletionTime)
	assert.Equal(t, sizeInGib, backup.SizeInGib)
	assert.Equal(t, backupType, backup.Type)

	collections := backup.Collections
	sort.Strings(collections)
	expectedCollections := []string{class1, class2}
	sort.Strings(expectedCollections)
	assert.Equal(t, expectedCollections, collections)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
	mockBackupBackend.AssertExpectations(t)
}

func TestService_Usage_WithNamedVectors(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	className := "NamedVectorClass"
	replication := 1
	shardName := ""
	objectCount := 2000
	storageSize := int64(10000000)
	vectorName := "abcd"
	textVectorName := "text"
	imageVectorName := "image"
	vectorType := "hnsw"
	compression := "standard"
	defaultCompressionRatio := 0.7
	textCompressionRatio := 0.7
	imageCompressionRatio := 0.8
	dimensionality := 1536
	textDimensionality := 768
	imageDimensionality := 1024
	dimensionCount := 2000

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:             className,
					VectorIndexConfig: &hnsw.UserConfig{},
					ReplicationConfig: &models.ReplicationConfig{Factor: int64(replication)},
				},
			},
		},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	shardingState := &sharding.State{
		Physical: map[string]sharding.Physical{
			shardName: {
				Name:           "",
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusHOT,
			},
		},
	}
	shardingState.SetLocalName(nodeName)
	mockSchema.EXPECT().CopyShardingState(className).Return(shardingState)

	mockDB := db.NewMockIndexGetter(t)
	mockIndex := db.NewMockIndexLike(t)
	mockDB.EXPECT().GetIndexLike(entschema.ClassName(className)).Return(mockIndex)

	mockShard := db.NewMockShardLike(t)
	mockShard.EXPECT().ObjectCountAsync(ctx).Return(int64(objectCount), nil)
	mockShard.EXPECT().ObjectStorageSize(ctx).Return(storageSize, nil)
	mockShard.EXPECT().VectorStorageSize(ctx).Return(int64(0), nil)
	mockShard.EXPECT().DimensionsUsage(ctx, vectorName).Return(types.Dimensionality{
		Dimensions: dimensionality,
		Count:      dimensionCount,
	}, nil)
	mockShard.EXPECT().DimensionsUsage(ctx, textVectorName).Return(types.Dimensionality{
		Dimensions: textDimensionality,
		Count:      dimensionCount,
	}, nil)
	mockShard.EXPECT().DimensionsUsage(ctx, imageVectorName).Return(types.Dimensionality{
		Dimensions: imageDimensionality,
		Count:      dimensionCount,
	}, nil)

	mockDefaultVectorIndex := db.NewMockVectorIndex(t)
	mockTextVectorIndex := db.NewMockVectorIndex(t)
	mockImageVectorIndex := db.NewMockVectorIndex(t)

	mockDefaultCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockDefaultCompressionStats.EXPECT().CompressionRatio(dimensionality).Return(defaultCompressionRatio)
	mockDefaultVectorIndex.EXPECT().CompressionStats().Return(mockDefaultCompressionStats)

	mockTextCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockTextCompressionStats.EXPECT().CompressionRatio(textDimensionality).Return(textCompressionRatio)
	mockTextVectorIndex.EXPECT().CompressionStats().Return(mockTextCompressionStats)

	mockImageCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockImageCompressionStats.EXPECT().CompressionRatio(imageDimensionality).Return(imageCompressionRatio)
	mockImageVectorIndex.EXPECT().CompressionStats().Return(mockImageCompressionStats)

	mockIndex.On("ForEachShard", mock.AnythingOfType("func(string, db.ShardLike) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.ShardLike) error)
		f(shardName, mockShard)
	})

	mockShard.On("ForEachVectorIndex", mock.AnythingOfType("func(string, db.VectorIndex) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.VectorIndex) error)
		f(vectorName, mockDefaultVectorIndex)
		f(textVectorName, mockTextVectorIndex)
		f(imageVectorName, mockImageVectorIndex)
	})

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 1)

	collection := result.Collections[0]
	assert.Equal(t, className, collection.Name)
	assert.Len(t, collection.Shards, 1)

	shard := collection.Shards[0]
	assert.Equal(t, shardName, shard.Name)
	assert.Equal(t, int64(objectCount), shard.ObjectsCount)
	assert.Equal(t, uint64(storageSize), shard.ObjectsStorageBytes)
	assert.Len(t, shard.NamedVectors, 3)

	defaultVector := shard.NamedVectors[0]
	assert.Equal(t, vectorName, defaultVector.Name)
	assert.Equal(t, vectorType, defaultVector.VectorIndexType)
	assert.Equal(t, compression, defaultVector.Compression)
	assert.Equal(t, defaultCompressionRatio, defaultVector.VectorCompressionRatio)
	assert.Len(t, defaultVector.Dimensionalities, 1)

	textVector := shard.NamedVectors[1]
	assert.Equal(t, textVectorName, textVector.Name)
	assert.Equal(t, vectorType, textVector.VectorIndexType)
	assert.Equal(t, compression, textVector.Compression)
	assert.Equal(t, textCompressionRatio, textVector.VectorCompressionRatio)
	assert.Len(t, textVector.Dimensionalities, 1)

	imageVector := shard.NamedVectors[2]
	assert.Equal(t, imageVectorName, imageVector.Name)
	assert.Equal(t, vectorType, imageVector.VectorIndexType)
	assert.Equal(t, compression, imageVector.Compression)
	assert.Equal(t, imageCompressionRatio, imageVector.VectorCompressionRatio)
	assert.Len(t, imageVector.Dimensionalities, 1)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockIndex.AssertExpectations(t)
	mockShard.AssertExpectations(t)
	mockDefaultVectorIndex.AssertExpectations(t)
	mockDefaultCompressionStats.AssertExpectations(t)
	mockTextVectorIndex.AssertExpectations(t)
	mockTextCompressionStats.AssertExpectations(t)
	mockImageVectorIndex.AssertExpectations(t)
	mockImageCompressionStats.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}

func TestService_Usage_EmptyCollections(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{Classes: []*models.Class{}},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	mockDB := db.NewMockIndexGetter(t)

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 0)
	assert.Len(t, result.Backups, 0)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}

func TestService_Usage_BackupError(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{Classes: []*models.Class{}},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	mockDB := db.NewMockIndexGetter(t)

	mockBackupBackend := modulecapabilities.NewMockBackupBackend(t)
	mockBackupBackend.EXPECT().AllBackups(ctx).Return(nil, assert.AnError)

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{mockBackupBackend})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	_, err := service.Usage(ctx)

	require.Error(t, err)
	require.ErrorIs(t, err, assert.AnError)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
	mockBackupBackend.AssertExpectations(t)
}

func TestService_Usage_VectorIndexError(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	className := "ErrorClass"
	replication := 1
	shardName := ""
	objectCount := 1000
	storageSize := int64(5000000)
	vectorName := "abcd"
	vectorType := "hnsw"
	compression := "standard"
	compressionRatio := 1.0
	dimensionality := 1536
	dimensionCount := 1000

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:             className,
					VectorIndexConfig: &hnsw.UserConfig{},
					ReplicationConfig: &models.ReplicationConfig{Factor: int64(replication)},
				},
			},
		},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	shardingState := &sharding.State{
		Physical: map[string]sharding.Physical{
			shardName: {
				Name:           "",
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusHOT,
			},
		},
	}
	shardingState.SetLocalName(nodeName)
	mockSchema.EXPECT().CopyShardingState(className).Return(shardingState)

	mockDB := db.NewMockIndexGetter(t)
	mockIndex := db.NewMockIndexLike(t)
	mockDB.EXPECT().GetIndexLike(entschema.ClassName(className)).Return(mockIndex)

	mockShard := db.NewMockShardLike(t)
	mockShard.EXPECT().ObjectCountAsync(ctx).Return(int64(objectCount), nil)
	mockShard.EXPECT().ObjectStorageSize(ctx).Return(storageSize, nil)
	mockShard.EXPECT().VectorStorageSize(ctx).Return(int64(0), nil)
	mockShard.EXPECT().DimensionsUsage(ctx, vectorName).Return(types.Dimensionality{
		Dimensions: dimensionality,
		Count:      dimensionCount,
	}, nil)

	mockVectorIndex := db.NewMockVectorIndex(t)
	mockVectorIndex.EXPECT().CompressionStats().Return(compressionhelpers.UncompressedStats{})

	mockIndex.On("ForEachShard", mock.AnythingOfType("func(string, db.ShardLike) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.ShardLike) error)
		f(shardName, mockShard)
	})

	mockShard.On("ForEachVectorIndex", mock.AnythingOfType("func(string, db.VectorIndex) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.VectorIndex) error)
		f(vectorName, mockVectorIndex)
	})

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 1)

	collection := result.Collections[0]
	assert.Len(t, collection.Shards, 1)

	shard := collection.Shards[0]
	assert.Len(t, shard.NamedVectors, 1)

	vector := shard.NamedVectors[0]
	assert.Equal(t, vectorName, vector.Name)
	assert.Equal(t, vectorType, vector.VectorIndexType)
	assert.Equal(t, compression, vector.Compression)
	assert.Equal(t, compressionRatio, vector.VectorCompressionRatio)
	assert.Len(t, vector.Dimensionalities, 1)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockIndex.AssertExpectations(t)
	mockShard.AssertExpectations(t)
	mockVectorIndex.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}

func TestService_Usage_NilVectorIndexConfig(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	className := "NilConfigClass"
	replication := 1
	shardName := ""
	objectCount := 1000
	storageSize := int64(5000000)
	vectorName := "abcd"
	vectorType := ""
	compression := "standard"
	compressionRatio := 0.75
	dimensionality := 1536
	dimensionCount := 1000

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:             className,
					VectorIndexConfig: nil,
					ReplicationConfig: &models.ReplicationConfig{Factor: int64(replication)},
				},
			},
		},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	shardingState := &sharding.State{
		Physical: map[string]sharding.Physical{
			shardName: {
				Name:           "",
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusHOT,
			},
		},
	}
	shardingState.SetLocalName(nodeName)
	mockSchema.EXPECT().CopyShardingState(className).Return(shardingState)

	mockDB := db.NewMockIndexGetter(t)
	mockIndex := db.NewMockIndexLike(t)
	mockDB.EXPECT().GetIndexLike(entschema.ClassName(className)).Return(mockIndex)

	mockShard := db.NewMockShardLike(t)
	mockShard.EXPECT().ObjectCountAsync(ctx).Return(int64(objectCount), nil)
	mockShard.EXPECT().ObjectStorageSize(ctx).Return(storageSize, nil)
	mockShard.EXPECT().VectorStorageSize(ctx).Return(int64(0), nil)
	mockShard.EXPECT().DimensionsUsage(ctx, vectorName).Return(types.Dimensionality{
		Dimensions: dimensionality,
		Count:      dimensionCount,
	}, nil)

	mockVectorIndex := db.NewMockVectorIndex(t)
	mockCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockCompressionStats.EXPECT().CompressionRatio(dimensionality).Return(compressionRatio)
	mockVectorIndex.EXPECT().CompressionStats().Return(mockCompressionStats)

	mockIndex.On("ForEachShard", mock.AnythingOfType("func(string, db.ShardLike) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.ShardLike) error)
		f(shardName, mockShard)
	})

	mockShard.On("ForEachVectorIndex", mock.AnythingOfType("func(string, db.VectorIndex) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.VectorIndex) error)
		f(vectorName, mockVectorIndex)
	})

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 1)

	collection := result.Collections[0]
	assert.Equal(t, className, collection.Name)
	assert.Len(t, collection.Shards, 1)

	shard := collection.Shards[0]
	assert.Equal(t, shardName, shard.Name)
	assert.Equal(t, int64(objectCount), shard.ObjectsCount)
	assert.Equal(t, uint64(storageSize), shard.ObjectsStorageBytes)
	assert.Len(t, shard.NamedVectors, 1)

	vector := shard.NamedVectors[0]
	assert.Equal(t, vectorName, vector.Name)
	assert.Equal(t, vectorType, vector.VectorIndexType)
	assert.Equal(t, compression, vector.Compression)
	assert.Equal(t, compressionRatio, vector.VectorCompressionRatio)
	assert.Len(t, vector.Dimensionalities, 1)
	dim := vector.Dimensionalities[0]
	assert.Equal(t, dimensionality, dim.Dimensions)
	assert.Equal(t, dimensionCount, dim.Count)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockIndex.AssertExpectations(t)
	mockShard.AssertExpectations(t)
	mockVectorIndex.AssertExpectations(t)
	mockCompressionStats.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}

func TestService_Usage_VectorStorageSize(t *testing.T) {
	ctx := context.Background()

	nodeName := "test-node"
	className := "VectorStorageClass"
	replication := 3
	uniqueShards := 2
	hotTenant := "hot-tenant"
	coldTenant := "cold-tenant"

	// Hot tenant metrics
	hotObjectCount := 2000
	hotStorageSize := int64(10000000)
	hotVectorStorageSize := int64(8000000) // 8MB for hot tenant vectors

	// Cold tenant metrics
	coldObjectCount := 1000
	coldStorageSize := int64(5000000)
	coldVectorStorageSize := int64(4000000) // 4MB for cold tenant vectors

	vectorName := "default"
	vectorType := "hnsw"
	compression := "standard"
	compressionRatio := 0.75
	dimensionality := 1536
	dimensionCount := 2000

	mockSchema := schema.NewMockSchemaGetter(t)
	mockSchema.EXPECT().GetSchemaSkipAuth().Return(entschema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class:              className,
					VectorIndexConfig:  &hnsw.UserConfig{},
					ReplicationConfig:  &models.ReplicationConfig{Factor: int64(replication)},
					MultiTenancyConfig: &models.MultiTenancyConfig{Enabled: true},
				},
			},
		},
	})
	mockSchema.EXPECT().NodeName().Return(nodeName)

	shardingState := &sharding.State{
		PartitioningEnabled: true,
		Physical: map[string]sharding.Physical{
			hotTenant: {
				Name:           hotTenant,
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusHOT,
			},
			coldTenant: {
				Name:           coldTenant,
				BelongsToNodes: []string{nodeName},
				Status:         models.TenantActivityStatusCOLD,
			},
		},
	}
	shardingState.SetLocalName(nodeName)
	mockSchema.EXPECT().CopyShardingState(className).Return(shardingState)

	mockDB := db.NewMockIndexGetter(t)
	mockIndex := db.NewMockIndexLike(t)
	mockDB.EXPECT().GetIndexLike(entschema.ClassName(className)).Return(mockIndex)

	// Mock hot tenant shard
	mockHotShard := db.NewMockShardLike(t)
	mockHotShard.EXPECT().ObjectCountAsync(ctx).Return(int64(hotObjectCount), nil)
	mockHotShard.EXPECT().ObjectStorageSize(ctx).Return(hotStorageSize, nil)
	mockHotShard.EXPECT().VectorStorageSize(ctx).Return(hotVectorStorageSize, nil) // Test actual vector storage size
	mockHotShard.EXPECT().DimensionsUsage(ctx, vectorName).Return(types.Dimensionality{
		Dimensions: dimensionality,
		Count:      dimensionCount,
	}, nil)

	mockVectorIndex := db.NewMockVectorIndex(t)
	mockCompressionStats := compressionhelpers.NewMockCompressionStats(t)
	mockCompressionStats.EXPECT().CompressionRatio(dimensionality).Return(compressionRatio)
	mockVectorIndex.EXPECT().CompressionStats().Return(mockCompressionStats)

	// Mock cold tenant calculations
	mockIndex.EXPECT().CalculateUnloadedObjectsMetrics(ctx, coldTenant).Return(types.ObjectUsage{
		Count:        int64(coldObjectCount),
		StorageBytes: coldStorageSize,
	}, nil)
	mockIndex.EXPECT().CalculateUnloadedVectorsMetrics(ctx, coldTenant).Return(coldVectorStorageSize, nil) // Test cold tenant vector storage

	mockIndex.On("ForEachShard", mock.AnythingOfType("func(string, db.ShardLike) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.ShardLike) error)
		f(hotTenant, mockHotShard)
	})

	mockHotShard.On("ForEachVectorIndex", mock.AnythingOfType("func(string, db.VectorIndex) error")).Return(nil).Run(func(args mock.Arguments) {
		f := args.Get(0).(func(string, db.VectorIndex) error)
		f(vectorName, mockVectorIndex)
	})

	mockBackupProvider := backupusecase.NewMockBackupBackendProvider(t)
	mockBackupProvider.EXPECT().EnabledBackupBackends().Return([]modulecapabilities.BackupBackend{})

	logger, _ := logrus.NewNullLogger()
	service := NewService(mockSchema, mockDB, mockBackupProvider, logger)

	result, err := service.Usage(ctx)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, nodeName, result.Node)
	assert.Len(t, result.Collections, 1)

	collection := result.Collections[0]
	assert.Equal(t, className, collection.Name)
	assert.Equal(t, replication, collection.ReplicationFactor)
	assert.Equal(t, uniqueShards, collection.UniqueShardCount)
	assert.Len(t, collection.Shards, 2)

	// Find hot and cold shards
	var hotShard, coldShard *types.ShardUsage
	for _, shard := range collection.Shards {
		switch shard.Name {
		case hotTenant:
			hotShard = shard
		case coldTenant:
			coldShard = shard
		}
	}

	// Verify hot tenant vector storage
	require.NotNil(t, hotShard)
	assert.Equal(t, int64(hotObjectCount), hotShard.ObjectsCount)
	assert.Equal(t, uint64(hotStorageSize), hotShard.ObjectsStorageBytes)
	assert.Equal(t, uint64(hotVectorStorageSize), hotShard.VectorStorageBytes) // Verify hot tenant vector storage
	assert.Len(t, hotShard.NamedVectors, 1)

	// Verify cold tenant vector storage
	require.NotNil(t, coldShard)
	assert.Equal(t, int64(coldObjectCount), coldShard.ObjectsCount)
	assert.Equal(t, uint64(coldStorageSize), coldShard.ObjectsStorageBytes)
	assert.Equal(t, uint64(coldVectorStorageSize), coldShard.VectorStorageBytes) // Verify cold tenant vector storage
	assert.Len(t, coldShard.NamedVectors, 0)                                     // Cold tenants don't have named vectors

	// Verify vector details for hot tenant
	vector := hotShard.NamedVectors[0]
	assert.Equal(t, vectorName, vector.Name)
	assert.Equal(t, vectorType, vector.VectorIndexType)
	assert.Equal(t, compression, vector.Compression)
	assert.Equal(t, compressionRatio, vector.VectorCompressionRatio)
	assert.Len(t, vector.Dimensionalities, 1)
	dim := vector.Dimensionalities[0]
	assert.Equal(t, dimensionality, dim.Dimensions)
	assert.Equal(t, dimensionCount, dim.Count)

	mockSchema.AssertExpectations(t)
	mockDB.AssertExpectations(t)
	mockIndex.AssertExpectations(t)
	mockHotShard.AssertExpectations(t)
	mockVectorIndex.AssertExpectations(t)
	mockCompressionStats.AssertExpectations(t)
	mockBackupProvider.AssertExpectations(t)
}
