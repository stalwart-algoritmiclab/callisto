/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package handlers

import (
	"fmt"

	"github.com/stalwart-algoritmiclab/callisto/modules/actions/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/rs/zerolog/log"
)

func ValidatorDelegation(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Int64("height", payload.Input.Height).
		Msg("executing validator delegation action")

	// Get latest node height
	height, err := ctx.GetHeight(payload)
	if err != nil {
		return nil, err
	}

	// Get validator's total delegations
	res, err := ctx.Sources.StakingSource.GetValidatorDelegationsWithPagination(height, payload.GetAddress(), payload.GetPagination())
	if err != nil {
		return nil, fmt.Errorf("error while getting validator delegations: %s", err)
	}

	delegations := make([]types.Delegation, len(res.DelegationResponses))
	for index, del := range res.DelegationResponses {
		delegations[index] = types.Delegation{
			DelegatorAddress: del.Delegation.DelegatorAddress,
			ValidatorAddress: del.Delegation.ValidatorAddress,
			Coins:            types.ConvertCoins([]sdk.Coin{del.Balance}),
		}
	}

	return types.DelegationResponse{
		Delegations: delegations,
		Pagination:  res.Pagination,
	}, nil
}
