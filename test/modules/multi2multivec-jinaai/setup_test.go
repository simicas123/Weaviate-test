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

package test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/weaviate/weaviate/test/docker"
)

func TestMulti2VecJinaAI_SingleNode(t *testing.T) {
	apiKey := os.Getenv("JINAAI_APIKEY")
	if apiKey == "" {
		t.Skip("skipping, JINAAI_APIKEY environment variable not present")
	}
	ctx := context.Background()
	compose, err := docker.New().
		WithWeaviate().
		WithMulti2MultivecJinaAI(apiKey).
		Start(ctx)
	require.NoError(t, err)
	defer func() {
		require.NoError(t, compose.Terminate(ctx))
	}()
	endpoint := compose.GetWeaviate().URI()

	t.Run("multi2multivec-jinaai", testMulti2MultivecJinaAI(endpoint))
}
