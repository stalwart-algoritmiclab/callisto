/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database

import (
	"encoding/json"
	"fmt"

	dbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"

	"github.com/stalwart-algoritmiclab/callisto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lib/pq"
)

// SaveCommunityPool allows to save for the given height the given total amount of coins
func (db *Db) SaveCommunityPool(coin sdk.DecCoins, height int64) error {
	query := `
INSERT INTO community_pool(coins, height) 
VALUES ($1, $2) 
ON CONFLICT (one_row_id) DO UPDATE 
    SET coins = excluded.coins,
        height = excluded.height
WHERE community_pool.height <= excluded.height`
	_, err := db.SQL.Exec(query, pq.Array(dbtypes.NewDbDecCoins(coin)), height)
	if err != nil {
		return fmt.Errorf("error while storing community pool: %s", err)
	}

	return nil
}

// -------------------------------------------------------------------------------------------------------------------

// SaveDistributionParams allows to store the given distribution parameters inside the database
func (db *Db) SaveDistributionParams(params *types.DistributionParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling params: %s", err)
	}

	stmt := `
INSERT INTO distribution_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
      	height = excluded.height
WHERE distribution_params.height <= excluded.height`
	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing distribution params: %s", err)
	}

	return nil
}

// GetDelegators returns the current delegators set
func (db *Db) GetDelegators() ([]string, error) {
	var rows []string
	err := db.Sqlx.Select(&rows, `SELECT DISTINCT (address) FROM top_accounts `)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
