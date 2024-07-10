/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package local

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/forbole/juno/v6/node/local"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/modules/auth/source"
)

var (
	_ source.Source = &Source{}
)

// Source implements authsource.Source by using a local node
type Source struct {
	*local.Source
	q authtypes.QueryServer
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, q authtypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      q,
	}
}

// GetAllAnyAccounts returns all accounts
func (s Source) GetAllAnyAccounts(height int64) ([]*codectypes.Any, error) {
	log.Debug().Msg("getting all accounts")
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return nil, fmt.Errorf("error while loading height: %s", err)
	}

	var (
		accounts    []*codectypes.Any
		nextKey     []byte
		stop        bool
		counter     uint64
		totalCounts uint64
		pageLimit   uint64 = 1000 // Get 1000 accounts per query
	)

	for !stop {
		// Get accounts
		res, err := s.q.Accounts(
			sdk.WrapSDKContext(ctx),
			&authtypes.QueryAccountsRequest{
				Pagination: &query.PageRequest{
					Key:        nextKey,
					Limit:      pageLimit,
					CountTotal: true,
				},
			})
		if err != nil {
			return nil, fmt.Errorf("error while getting any accounts from source: %s", err)
		}
		nextKey = res.Pagination.NextKey
		stop = len(res.Pagination.NextKey) == 0
		accounts = append(accounts, res.Accounts...)

		// Log getting accounts progress
		if res.Pagination.GetTotal() != 0 {
			totalCounts = res.Pagination.GetTotal()
		}
		counter += uint64(len(res.Accounts))
		log.Debug().Uint64("total accounts", totalCounts).Uint64("current counter", counter).Msg("getting accounts...")
	}

	return accounts, nil
}

// GetTotalNumberOfAccounts returns the total number of accounts
func (s Source) GetTotalNumberOfAccounts(height int64) (uint64, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return 0, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Accounts(
		ctx,
		&authtypes.QueryAccountsRequest{})
	if err != nil {
		return 0, fmt.Errorf("error while getting total number of accounts from source: %s", err)
	}

	return res.Pagination.Total, nil
}
