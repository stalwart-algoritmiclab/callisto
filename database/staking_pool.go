/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database

import (
	"fmt"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// SaveStakingPool allows to save for the given height the given stakingtypes pool
func (db *Db) SaveStakingPool(pool *types.Pool) error {
	stmt := `
INSERT INTO staking_pool (bonded_tokens, not_bonded_tokens, unbonding_tokens, staked_not_bonded_tokens, height) 
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (one_row_id) DO UPDATE 
    SET bonded_tokens = excluded.bonded_tokens, 
        not_bonded_tokens = excluded.not_bonded_tokens, 
		unbonding_tokens = excluded.unbonding_tokens,
		staked_not_bonded_tokens = excluded.staked_not_bonded_tokens,
        height = excluded.height
WHERE staking_pool.height <= excluded.height`

	_, err := db.SQL.Exec(stmt,
		pool.BondedTokens.String(),
		pool.NotBondedTokens.String(),
		pool.UnbondingTokens.String(),
		pool.StakedNotBondedTokens.String(),
		pool.Height)
	if err != nil {
		return fmt.Errorf("error while storing staking pool: %s", err)
	}

	return nil
}
