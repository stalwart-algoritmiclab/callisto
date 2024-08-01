/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package local

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
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

// GetAllPoll implements source.Source
func (s Source) GetAllPoll(height int64, pagination *query.PageRequest) (*types.QueryAllPollsResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, err
	}

	res, err := s.pollsServer.PollsAll(
		sdk.WrapSDKContext(ctx),
		&types.QueryAllPollsRequest{
			Pagination: pagination,
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetPoll implements source.Source
func (s Source) GetPoll(pollID uint64, height int64) (*types.QueryGetPollsResponse, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, err
	}

	res, err := s.pollsServer.Polls(sdk.WrapSDKContext(ctx), &types.QueryGetPollsRequest{
		Id: pollID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
