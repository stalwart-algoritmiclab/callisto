/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

// StakingPoolRow represents a single row inside the staking_pool table
type StakingPoolRow struct {
	OneRowID              bool  `db:"one_row_id"`
	BondedTokens          int64 `db:"bonded_tokens"`
	NotBondedTokens       int64 `db:"not_bonded_tokens"`
	UnbondingTokens       int64 `db:"unbonding_tokens"`
	StakedNotBondedTokens int64 `db:"staked_not_bonded_tokens"`
	Height                int64 `db:"height"`
}

// NewStakingPoolRow allows to easily create a new StakingPoolRow
func NewStakingPoolRow(bondedTokens, notBondedTokens, unbondingTokens, stakedNotBondedTokens int64, height int64) StakingPoolRow {
	return StakingPoolRow{
		OneRowID:              true,
		BondedTokens:          bondedTokens,
		NotBondedTokens:       notBondedTokens,
		UnbondingTokens:       unbondingTokens,
		StakedNotBondedTokens: stakedNotBondedTokens,
		Height:                height,
	}
}

// Equal allows to tells whether r and as represent the same rows
func (r StakingPoolRow) Equal(s StakingPoolRow) bool {
	return r.BondedTokens == s.BondedTokens &&
		r.NotBondedTokens == s.NotBondedTokens &&
		r.UnbondingTokens == s.UnbondingTokens &&
		r.StakedNotBondedTokens == s.StakedNotBondedTokens &&
		r.Height == s.Height
}
