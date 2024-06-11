/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package source

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

type Source interface {
	ValidatorCommission(valOperAddr string, height int64) (sdk.DecCoins, error)
	DelegatorTotalRewards(delegator string, height int64) ([]distrtypes.DelegationDelegatorReward, error)
	DelegatorWithdrawAddress(delegator string, height int64) (string, error)
	CommunityPool(height int64) (sdk.DecCoins, error)
	Params(height int64) (distrtypes.Params, error)
}
