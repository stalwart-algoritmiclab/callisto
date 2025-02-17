/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package auth

import (
	"encoding/json"
	"fmt"

	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/types/config"
	"github.com/spf13/cobra"

	"github.com/stalwart-algoritmiclab/callisto/database"
	authutils "github.com/stalwart-algoritmiclab/callisto/modules/auth"
	"github.com/stalwart-algoritmiclab/callisto/utils"
)

// vestingCmd returns a Cobra command that allows to fix the vesting data for the accounts
func vestingCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "vesting-accounts",
		Short: "Fix the vesting accounts stored by removing duplicated vesting periods",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Get the genesis
			genesis, err := utils.ReadGenesis(config.Cfg, parseCtx.Node)
			if err != nil {
				return fmt.Errorf("error while reading the genesis: %s", err)
			}

			var appState map[string]json.RawMessage
			if err := json.Unmarshal(genesis.AppState, &appState); err != nil {
				return fmt.Errorf("error unmarshalling genesis doc: %s", err)
			}

			vestingAccounts, err := authutils.GetGenesisVestingAccounts(appState, utils.GetCodec())
			if err != nil {
				return fmt.Errorf("error while gestting vesting accounts: %s", err)
			}

			err = db.SaveVestingAccounts(vestingAccounts)
			if err != nil {
				return fmt.Errorf("error while storing vesting accounts: %s", err)
			}

			return nil
		},
	}
}
