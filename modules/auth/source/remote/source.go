/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package remote

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/forbole/juno/v6/node/remote"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/modules/auth/source"
)

var (
	_ source.Source = &Source{}
)

// Source implements authsource.Source by using a remote node
type Source struct {
	*remote.Source
	authClient authtypes.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, authClient authtypes.QueryClient) *Source {
	return &Source{
		Source:     source,
		authClient: authClient,
	}
}

// GetAllAnyAccounts returns all accounts
func (s Source) GetAllAnyAccounts(height int64) ([]*codectypes.Any, error) {
	log.Debug().Msg("getting all accounts")
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	var (
		accounts    []*codectypes.Any
		nextKey     []byte
		stop        bool = false
		counter     uint64
		totalCounts uint64
		pageLimit   uint64 = 1000 // Get 1000 accounts per query
	)

	for !stop {
		// Get accounts
		res, err := s.authClient.Accounts(
			ctx,
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
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	res, err := s.authClient.Accounts(
		ctx,
		&authtypes.QueryAccountsRequest{})
	if err != nil {
		return 0, fmt.Errorf("error while getting total number of accounts from source: %s", err)
	}

	return res.Pagination.Total, nil
}
