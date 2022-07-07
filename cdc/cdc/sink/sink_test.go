// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package sink

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tikv/migration/cdc/pkg/config"
	"github.com/tikv/migration/cdc/pkg/util/testleak"
)

func TestValidateSink(t *testing.T) {
	defer testleak.AfterTestT(t)()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	replicateConfig := config.GetDefaultReplicaConfig()
	opts := make(map[string]string)

	// test sink uri right
	sinkURI := "tikv://127.0.0.1:3306/"
	err := Validate(ctx, sinkURI, replicateConfig, opts)
	require.Nil(t, err)

	sinkURI = "tikv://127.0.0.1:3306/?concurrency=4"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.Nil(t, err)

	sinkURI = "tikv://127.0.0.1:3306,127.0.0.1:3307/?concurrency=4"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.Nil(t, err)

	sinkURI = "blackhole://"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.Nil(t, err)

	// test sink uri wrong
	sinkURI = "tikv://http://127.0.0.1:3306/"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.NotNil(t, err)

	sinkURI = "tikv://127.0.0.1:3306a/"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.NotNil(t, err)

	sinkURI = "tikv://a127.0.0.1:3306/"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.NotNil(t, err)

	sinkURI = "tikv://127.0.0.1:3306, tikv://127.0.0.1:3307/"
	err = Validate(ctx, sinkURI, replicateConfig, opts)
	require.NotNil(t, err)
}
