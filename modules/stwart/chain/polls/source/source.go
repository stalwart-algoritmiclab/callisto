/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package source

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

// Source represents the implementation of the polls keeper that works on a local node
type Source interface {
	// GetAllPoll returns all polls
	GetAllPoll(height int64, pagination *query.PageRequest) (*types.QueryAllPollsResponse, error)
	// GetPoll returns a poll by its ID
	GetPoll(pollID uint64, height int64) (*types.QueryGetPollsResponse, error)
}
