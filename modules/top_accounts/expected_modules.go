/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package top_accounts

import (
	"github.com/stalwart-algoritmiclab/callisto/types"
)

// AuthModule represents the module that allows to interact with the auth module
type AuthModule interface {
	GetAllBaseAccounts(height int64) ([]types.Account, error)
	RefreshTopAccountsList(height int64) ([]types.Account, error)
}

// AuthSource represents the module that allows to interact with the auth module
type AuthSource interface {
	GetTotalNumberOfAccounts(height int64) (uint64, error)
}

// BankModule represents the module that allows to interact with the bank module
type BankModule interface {
	UpdateBalances(addresses []string, height int64) error
}

// DistrModule represents the module that allows to interact with the distribution module
type DistrModule interface {
	RefreshDelegatorRewards(delegators []string, height int64) error
}

// StakingModule represents the module that allows to interact with the staking module
type StakingModule interface {
	RefreshDelegations(delegatorAddr string, height int64) error
	RefreshUnbondings(delegatorAddr string, height int64) error
}
