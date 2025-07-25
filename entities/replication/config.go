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

package replication

import "github.com/weaviate/weaviate/usecases/config/runtime"

// GlobalConfig represents system-wide config that may restrict settings of an
// individual class
type GlobalConfig struct {
	AsyncReplicationDisabled *runtime.DynamicValue[bool] `json:"async_replication_disabled" yaml:"async_replication_disabled"`
	// MinimumFactor can enforce replication. For example, with MinimumFactor set
	// to 2, users can no longer create classes with a factor of 1, therefore
	// forcing them to have replicated classes.
	MinimumFactor int `json:"minimum_factor" yaml:"minimum_factor"`

	DeletionStrategy string `json:"deletion_strategy" yaml:"deletion_strategy"`
}
