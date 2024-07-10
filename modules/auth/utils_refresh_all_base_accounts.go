/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package auth

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// GetAllBaseAccounts returns all the base accounts at the given height
func (m *Module) GetAllBaseAccounts(height int64) ([]types.Account, error) {
	anyAccounts, err := m.source.GetAllAnyAccounts(height)
	if err != nil {
		return nil, fmt.Errorf("error while getting any accounts: %s", err)
	}
	unpacked, err := m.unpackAnyAccounts(anyAccounts)
	if err != nil {
		return nil, err
	}

	return unpacked, nil

}

// RefreshTopAccountsList refreshes the top accounts list at the given height
func (m *Module) RefreshTopAccountsList(height int64) ([]types.Account, error) {
	// Get all accounts from the node
	anyAccounts, err := m.source.GetAllAnyAccounts(height)
	if err != nil {
		return nil, fmt.Errorf("error while getting any accounts: %s", err)
	}

	// Unpack all accounts into types.Account type
	accounts, err := m.unpackAnyAccounts(anyAccounts)
	if err != nil {
		return nil, fmt.Errorf("error while unpacking accounts: %s", err)
	}

	log.Debug().Int("total", len(accounts)).Msg("saving accounts...")
	// Store accounts
	err = m.db.SaveAccounts(accounts)
	if err != nil {
		return nil, err
	}

	// Unpack all accounts into types.TopAccount type
	accountsWithTypes, err := m.unpackAnyAccountsWithTypes(anyAccounts)
	if err != nil {
		return nil, fmt.Errorf("error while unpacking top accounts with account types: %s", err)
	}

	log.Debug().Int("total", len(accounts)).Msg("saving top accounts addresses...")
	// Store all top accounts addresses with account type
	err = m.db.SaveTopAccounts(accountsWithTypes, height)
	if err != nil {
		return nil, fmt.Errorf("error while storing top accounts with account types: %s", err)
	}

	return accounts, nil
}

// unpackAnyAccounts unpacks the given any accounts into types.Account
func (m *Module) unpackAnyAccounts(anyAccounts []*codectypes.Any) ([]types.Account, error) {
	accounts := []types.Account{}
	for _, account := range anyAccounts {
		var accountI authtypes.AccountI
		err := m.cdc.UnpackAny(account, &accountI)
		if err != nil {
			return nil, fmt.Errorf("error while unpacking any account: %s", err)
		}

		accounts = append(accounts, types.NewAccount(accountI.GetAddress().String()))
	}

	return accounts, nil
}

// unpackAnyAccountsWithTypes unpacks the given any accounts into types.TopAccount
func (m *Module) unpackAnyAccountsWithTypes(anyAccounts []*codectypes.Any) ([]types.TopAccount, error) {
	accounts := []types.TopAccount{}
	for _, account := range anyAccounts {
		var accountI authtypes.AccountI
		err := m.cdc.UnpackAny(account, &accountI)
		if err != nil {
			return nil, fmt.Errorf("error while unpacking any account: %s", err)
		}

		accounts = append(accounts, types.NewTopAccount(accountI.GetAddress().String(), account.TypeUrl))
	}

	return accounts, nil
}
