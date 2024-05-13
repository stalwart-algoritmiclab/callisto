/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database_test

import (
	"github.com/forbole/callisto/database/types"
)

func (suite *DbTestSuite) TestBigDipperDb_InsertEnableModules() {
	modules := []string{"auth", "bank", "consensus", "distribution", "gov", "mint", "pricefeed", "staking", "supply"}
	err := suite.database.InsertEnableModules(modules)
	suite.Require().NoError(err)

	var results types.ModuleRows
	err = suite.database.Sqlx.Select(&results, "SELECT * FROM modules")
	suite.Require().NoError(err)

	expected := types.NewModuleRows(modules)
	suite.Require().True(results.Equal(&expected))

}
