/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package remote

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/forbole/juno/v6/node/remote"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/polls/source"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryClient that works on a remote node
type Source struct {
	*remote.Source
	client types.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, pollsClient types.QueryClient) *Source {
	return &Source{
		Source: source,
		client: pollsClient,
	}
}

// GetAllPoll implements source.Source
func (s Source) GetAllPoll(height int64, pagination *query.PageRequest) (*types.QueryAllPollsResponse, error) {
	ctx := s.Ctx
	res, err := s.client.PollsAll(
		remote.GetHeightRequestContext(ctx, height),
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
	ctx := s.Ctx
	res, err := s.client.Polls(
		remote.GetHeightRequestContext(ctx, height),
		&types.QueryGetPollsRequest{
			Id: pollID,
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
