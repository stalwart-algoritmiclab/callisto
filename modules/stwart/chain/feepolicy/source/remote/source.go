/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package remote

import (
	"github.com/forbole/juno/v5/node/remote"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/feepolicy/source"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryClient that works on a remote node
type Source struct {
	*remote.Source
	client feepolicy.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, feepolicyClient feepolicy.QueryClient) *Source {
	return &Source{
		Source: source,
		client: feepolicyClient,
	}
}