/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package staking

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// RefreshDelegations refreshes the delegations for the given delegator at the given height
func (m *Module) RefreshDelegations(delegatorAddr string, height int64) error {
	log.Debug().
		Str("module", "staking").
		Int64("height", height).Msg("updating delegations")

	var (
		coin    = sdk.Coin{Denom: types.TokenDenomSTW, Amount: math.NewInt(0)}
		nextKey []byte
		stop    bool
	)

	for !stop {
		res, err := m.source.GetDelegationsWithPagination(
			height,
			delegatorAddr,
			&query.PageRequest{
				Key:   nextKey,
				Limit: 100,
			},
		)
		if err != nil {
			return fmt.Errorf("error while getting delegations: %s", err)
		}

		nextKey = res.Pagination.NextKey
		stop = len(res.Pagination.NextKey) == 0

		for _, r := range res.DelegationResponses {
			coin = coin.Add(r.Balance)
		}
	}

	err := m.db.SaveTopAccountsBalance("delegation",
		[]types.NativeTokenAmount{
			types.NewNativeTokenAmount(delegatorAddr, coin.Amount, height),
		})
	if err != nil {
		return fmt.Errorf("error while saving top accounts delegation from MsgDelegate: %s", err)
	}

	return nil
}
