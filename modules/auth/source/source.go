/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package source

import codectypes "github.com/cosmos/cosmos-sdk/codec/types"

// Source represents the source from which the accounts are retrieved
type Source interface {
	GetAllAnyAccounts(height int64) ([]*codectypes.Any, error)
	GetTotalNumberOfAccounts(height int64) (uint64, error)
}
