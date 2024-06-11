/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

// FeeAllowanceRow represents a single row inside the fee_grant_allowance table
type FeeAllowanceRow struct {
	ID        uint64 `db:"id"`
	Grantee   string `db:"grantee_address"`
	Granter   string `db:"granter_address"`
	Allowance string `db:"allowance"`
	Height    int64  `db:"height"`
}
