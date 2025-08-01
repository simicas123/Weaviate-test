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

package modstgazure

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/weaviate/weaviate/entities/moduletools"
)

// Test user overrides
func TestUploadParams(t *testing.T) {
	defaultBlockSize := int64(40 * 1024 * 1024)
	defaultEnvironmentValue := int64(11)
	defaultHeaderValue := int64(13)
	testCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	azure := New()
	os.Setenv("BACKUP_AZURE_CONTAINER", "test")
	os.Setenv("AZURE_STORAGE_ACCOUNT", "test")

	params := moduletools.NewMockModuleInitParams(t)
	params.EXPECT().GetLogger().Return(logrus.New())
	params.EXPECT().GetStorageProvider().Return(&fakeStorageProvider{dataPath: t.TempDir()})
	err := azure.Init(testCtx, params)
	require.Nil(t, err)

	t.Run("getBlockSize with no inputs", func(t *testing.T) {
		blockSize := azure.getBlockSize(testCtx)
		assert.Equal(t, defaultBlockSize, blockSize)
	})

	t.Run("getBlockSize with environment variable", func(t *testing.T) {
		t.Setenv("AZURE_BLOCK_SIZE", "11")
		azure := New()
		params := moduletools.NewMockModuleInitParams(t)
		params.EXPECT().GetLogger().Return(logrus.New())
		params.EXPECT().GetStorageProvider().Return(&fakeStorageProvider{dataPath: t.TempDir()})
		err := azure.Init(testCtx, params)
		assert.Nil(t, err)

		blockSize := azure.getBlockSize(testCtx)
		assert.Equal(t, defaultEnvironmentValue, blockSize)
	})

	t.Run("getBlockSize with invalid environment variable", func(t *testing.T) {
		t.Setenv("AZURE_BLOCK_SIZE", "invalid")
		azure := New()
		params := moduletools.NewMockModuleInitParams(t)
		params.EXPECT().GetLogger().Return(logrus.New())
		params.EXPECT().GetStorageProvider().Return(&fakeStorageProvider{dataPath: t.TempDir()})
		err := azure.Init(testCtx, params)
		assert.Nil(t, err)

		blockSize := azure.getBlockSize(testCtx)
		assert.Equal(t, defaultBlockSize, blockSize)
	})

	t.Run("getBlockSize with header", func(t *testing.T) {
		ctxWithValue := context.WithValue(context.Background(),
			"X-Azure-Block-Size", []string{"13"})

		blockSize := azure.getBlockSize(ctxWithValue)
		assert.Equal(t, defaultHeaderValue, blockSize)
	})

	t.Run("getBlockSize with invalid header", func(t *testing.T) {
		ctxWithValue := context.WithValue(context.Background(),
			"X-Azure-Block-Size", []string{"invalid"})

		blockSize := azure.getBlockSize(ctxWithValue)
		assert.Equal(t, defaultBlockSize, blockSize)
	})

	t.Run("getBlockSize with environment variable and header", func(t *testing.T) {
		t.Setenv("AZURE_BLOCK_SIZE", "11")
		ctxWithValue := context.WithValue(context.Background(),
			"X-Azure-Block-Size", []string{"13"})

		blockSize := azure.getBlockSize(ctxWithValue)
		assert.Equal(t, defaultHeaderValue, blockSize)
	})

	t.Run("getConcurrency with no inputs", func(t *testing.T) {
		concurrency := azure.getConcurrency(testCtx)
		assert.Equal(t, 1, concurrency)
	})

	t.Run("getConcurrency with environment variable", func(t *testing.T) {
		t.Setenv("AZURE_CONCURRENCY", "11")
		azure := New()
		params := moduletools.NewMockModuleInitParams(t)
		params.EXPECT().GetLogger().Return(logrus.New())
		params.EXPECT().GetStorageProvider().Return(&fakeStorageProvider{dataPath: t.TempDir()})
		err := azure.Init(testCtx, params)
		assert.Nil(t, err)

		concurrency := azure.getConcurrency(testCtx)
		assert.Equal(t, defaultEnvironmentValue, int64(concurrency))
	})

	t.Run("getConcurrency with invalid environment variable", func(t *testing.T) {
		t.Setenv("AZURE_CONCURRENCY", "invalid")
		azure := New()
		params := moduletools.NewMockModuleInitParams(t)
		params.EXPECT().GetLogger().Return(logrus.New())
		params.EXPECT().GetStorageProvider().Return(&fakeStorageProvider{dataPath: t.TempDir()})
		err := azure.Init(testCtx, params)
		assert.Nil(t, err)

		concurrency := azure.getConcurrency(testCtx)
		assert.Equal(t, 1, concurrency)
	})

	t.Run("getConcurrency with header", func(t *testing.T) {
		ctxWithValue := context.WithValue(context.Background(),
			"X-Azure-Concurrency", []string{"13"})

		concurrency := azure.getConcurrency(ctxWithValue)
		assert.Equal(t, defaultHeaderValue, int64(concurrency))
	})

	t.Run("getConcurrency with invalid header", func(t *testing.T) {
		ctxWithValue := context.WithValue(context.Background(),
			"X-Azure-Concurrency", []string{"invalid"})

		concurrency := azure.getConcurrency(ctxWithValue)
		assert.Equal(t, 1, concurrency)
	})

	t.Run("getConcurrency with environment variable and header", func(t *testing.T) {
		t.Setenv("AZURE_CONCURRENCY", "11")
		ctxWithValue := context.WithValue(context.Background(),
			"X-Azure-Concurrency", []string{"13"})

		concurrency := azure.getConcurrency(ctxWithValue)
		assert.Equal(t, defaultHeaderValue, int64(concurrency))
	})
}

type fakeStorageProvider struct {
	dataPath string
}

func (f *fakeStorageProvider) Storage(name string) (moduletools.Storage, error) {
	return nil, nil
}

func (f *fakeStorageProvider) DataPath() string {
	return f.dataPath
}
