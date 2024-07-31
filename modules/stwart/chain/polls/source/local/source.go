/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package local

import (
	"github.com/forbole/juno/v6/node/local"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/polls/source"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryServer that works on a local node
type Source struct {
	*local.Source
	pollsServer types.QueryServer
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, pollsServer types.QueryServer) *Source {
	return &Source{
		Source:      source,
		pollsServer: pollsServer,
	}
}
