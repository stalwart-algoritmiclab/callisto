/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package distribution

import (
	"fmt"

	"cosmossdk.io/math"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// RefreshDelegatorRewards refreshes the rewards for the given delegators at the given height
func (m *Module) RefreshDelegatorRewards(delegators []string, height int64) error {
	log.Debug().
		Str("module", "distribution").
		Int64("height", height).Msg("updating rewards")

	var nativeTokenAmounts []types.NativeTokenAmount
	for _, del := range delegators {
		rews, err := m.source.DelegatorTotalRewards(del, height)
		if err != nil {
			return fmt.Errorf("error while getting delegator rewards: %s", err)
		}

		amount := math.LegacyNewDec(0)
		for _, r := range rews {
			decCoinAmount := r.Reward.AmountOf(types.TokenDenomSTW)
			amount = amount.Add(decCoinAmount)

			nativeTokenAmounts = append(nativeTokenAmounts, types.NewNativeTokenAmount(del, amount.RoundInt(), height))
		}
	}

	err := m.db.SaveTopAccountsBalance("reward", nativeTokenAmounts)
	if err != nil {
		return fmt.Errorf("error while saving delegators rewards amounts: %s", err)
	}

	return nil
}
