/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package bank

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// UpdateBalances updates the balances of the accounts having the given addresses,
// taking the data at the provided height
func (m *Module) UpdateBalances(addresses []string, height int64) error {
	log.Debug().Str("module", "bank").Int64("height", height).Msg("updating balances")

	balances, err := m.keeper.GetBalances(addresses, height)
	if err != nil {
		return fmt.Errorf("error while getting account balances: %s", err)
	}

	var nativeTokenAmounts []types.NativeTokenAmount
	for _, balance := range balances {
		denomAmount := balance.Balance.AmountOf(types.TokenDenomSTW)
		if denomAmount.IsNil() {
			continue
		}

		nativeTokenAmounts = append(nativeTokenAmounts, types.NewNativeTokenAmount(balance.Address, denomAmount, height))
	}

	err = m.db.SaveTopAccountsBalance("available", nativeTokenAmounts)
	if err != nil {
		return fmt.Errorf("error while saving top accounts available balances: %s", err)
	}

	return nil
}

// UpdateSSCBalances updates the balances of the accounts having the given addresses,
// taking the data at the provided height
func (m *Module) UpdateSSCBalances(addresses []string, height int64) error {
	log.Debug().Str("module", "bank").Int64("height", height).Msg("updating SSC balances")

	balances, err := m.keeper.GetBalances(addresses, height)
	if err != nil {
		return fmt.Errorf("error while getting account balances: %s", err)
	}

	var nativeTokenAmounts []types.NativeTokenAmount
	for _, balance := range balances {
		denomAmount := balance.Balance.AmountOf(types.TokenDenomSSC)
		if denomAmount.IsNil() {
			continue
		}

		nativeTokenAmounts = append(nativeTokenAmounts, types.NewNativeTokenAmount(balance.Address, denomAmount, height))
	}

	err = m.db.SaveTopAccountsBalance("ssc", nativeTokenAmounts)
	if err != nil {
		return fmt.Errorf("error while saving top accounts SSC balances: %s", err)
	}

	return nil
}
