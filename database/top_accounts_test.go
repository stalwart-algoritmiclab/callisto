/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database_test

import (
	"cosmossdk.io/math"

	dbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"
	"github.com/stalwart-algoritmiclab/callisto/types"
)

const address = "stwart1kfkpkya7tmwltv2gc799exrerm2hrg73phkzm4"

// TestSaveTopAccounts tests saving top accounts
func (suite *DbTestSuite) TestSaveTopAccountsBalance() {
	suite.getAccount(address)

	// Test saving balances
	amount := types.NewNativeTokenAmount(
		address,
		math.NewInt(100),
		100,
	)

	account := types.NewAccount(address)
	err := suite.database.SaveAccounts([]types.Account{account})
	suite.Require().NoError(err)

	topAccount := types.NewTopAccount(address, "/cosmos.auth.v1beta1.BaseAccount")
	err = suite.database.SaveTopAccounts([]types.TopAccount{topAccount}, 100)
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("available", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("delegation", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("unbonding", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("reward", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.UpdateTopAccountsSum(address, "400", 100)
	suite.Require().NoError(err)

	// Verify data
	expected := dbtypes.NewTopAccountsRow(address, "cosmos.auth.v1beta1.BaseAccount", 100, 100, 100, 100, 400, 100)

	var rows []dbtypes.TopAccountsRow
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM top_accounts`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(expected.Equals(rows[0]))

	// Test saving higher values
	newAmount := types.NewNativeTokenAmount(
		address,
		math.NewInt(100),
		300,
	)

	err = suite.database.SaveTopAccounts([]types.TopAccount{topAccount}, 200)
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("available", []types.NativeTokenAmount{newAmount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("delegation", []types.NativeTokenAmount{newAmount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("unbonding", []types.NativeTokenAmount{newAmount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("reward", []types.NativeTokenAmount{newAmount})
	suite.Require().NoError(err)

	err = suite.database.UpdateTopAccountsSum(address, "800", 300)
	suite.Require().NoError(err)

	// Verify data
	expected = dbtypes.NewTopAccountsRow(address, "cosmos.auth.v1beta1.BaseAccount", 200, 200, 200, 200, 800, 300)
	err = suite.database.Sqlx.Select(&rows, `SELECT * FROM top_accounts`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().True(expected.Equals(rows[0]))

}

// TestGetAccountBalanceSum tests getting account balance sum
func (suite *DbTestSuite) TestGetAccountBalanceSum() {
	suite.getAccount(address)

	// Store balances
	amount := types.NewNativeTokenAmount(
		address,
		math.NewInt(100),
		10,
	)

	err := suite.database.SaveTopAccountsBalance("available", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("delegation", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("unbonding", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	err = suite.database.SaveTopAccountsBalance("reward", []types.NativeTokenAmount{amount})
	suite.Require().NoError(err)

	// Verify Data
	expectedSum := "400"
	sum, err := suite.database.GetAccountBalanceSum(address)
	suite.Require().NoError(err)
	suite.Require().Equal(expectedSum, sum)

	// Verify getting 0 amount
	expectedSum = "0"
	sum, err = suite.database.GetAccountBalanceSum("")
	suite.Require().NoError(err)
	suite.Require().Equal(expectedSum, sum)
}

// TestUpdateTopAccountsSum tests updating top accounts sum
func (suite *DbTestSuite) TestUpdateTopAccountsSum() {
	suite.getAccount(address)

	// Store top accounts sum
	amount := "100"
	err := suite.database.UpdateTopAccountsSum(address, amount, 500)
	suite.Require().NoError(err)

	// Verify data
	var rows []string
	err = suite.database.Sqlx.Select(&rows, `SELECT sum FROM top_accounts`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(amount, rows[0])

	// Store different amount
	amount = "200"
	err = suite.database.UpdateTopAccountsSum(address, amount, 500)
	suite.Require().NoError(err)

	// Verify data
	err = suite.database.Sqlx.Select(&rows, `SELECT sum FROM top_accounts`)
	suite.Require().NoError(err)
	suite.Require().Len(rows, 1)
	suite.Require().Equal(amount, rows[0])
}
