/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database_test

import (
	"encoding/json"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	dbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"

	"github.com/stalwart-algoritmiclab/callisto/types"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	bddbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"
)

func (suite *DbTestSuite) TestBigDipperDb_SaveCommunityPool() {
	// Save data
	original := sdk.NewDecCoins(sdk.NewDecCoin("uatom", math.NewInt(100)))
	err := suite.database.SaveCommunityPool(original, 10)
	suite.Require().NoError(err)

	// Verify data
	expected := bddbtypes.NewCommunityPoolRow(dbtypes.NewDbDecCoins(original), 10)
	var rows []bddbtypes.CommunityPoolRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]))

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with lower height
	coins := sdk.NewDecCoins(sdk.NewDecCoin("uatom", math.NewInt(50)))
	err = suite.database.SaveCommunityPool(coins, 5)
	suite.Require().NoError(err)

	// Verify data
	expected = bddbtypes.NewCommunityPoolRow(dbtypes.NewDbDecCoins(original), 10)
	rows = []bddbtypes.CommunityPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]), "updating with lower height should not modify the data")

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with equal height
	coins = sdk.NewDecCoins(sdk.NewDecCoin("uatom", math.NewInt(120)))
	err = suite.database.SaveCommunityPool(coins, 10)
	suite.Require().NoError(err)

	// Verify data
	expected = bddbtypes.NewCommunityPoolRow(dbtypes.NewDbDecCoins(coins), 10)
	rows = []bddbtypes.CommunityPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]), "updating with same height should modify the data")

	// ---------------------------------------------------------------------------------------------------------------

	// Try updating with higher height
	coins = sdk.NewDecCoins(sdk.NewDecCoin("uatom", math.NewInt(200)))
	err = suite.database.SaveCommunityPool(coins, 11)
	suite.Require().NoError(err)

	// Verify data
	expected = bddbtypes.NewCommunityPoolRow(dbtypes.NewDbDecCoins(coins), 11)
	rows = []bddbtypes.CommunityPoolRow{}
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM community_pool`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1, "community_pool table should contain only one row")
	suite.Require().True(expected.Equals(rows[0]), "updating with higher height should modify the data")
}

func (suite *DbTestSuite) TestBigDipperDb_SaveDistributionParams() {
	distrParams := distrtypes.Params{
		CommunityTax:        math.LegacyNewDecWithPrec(2, 2),
		BaseProposerReward:  math.LegacyNewDecWithPrec(1, 2),
		BonusProposerReward: math.LegacyNewDecWithPrec(4, 2),
		WithdrawAddrEnabled: true,
	}
	err := suite.database.SaveDistributionParams(types.NewDistributionParams(distrParams, 10))
	suite.Require().NoError(err)

	var rows []dbtypes.DistributionParamsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM distribution_params`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)

	var stored distrtypes.Params
	err = json.Unmarshal([]byte(rows[0].Params), &stored)
	suite.Require().NoError(err)
	suite.Require().Equal(distrParams, stored)
	suite.Require().Equal(int64(10), rows[0].Height)
}
