/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package gov

import (
	"github.com/spf13/cobra"

	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/types/config"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/modules/distribution"
	"github.com/stalwart-algoritmiclab/callisto/modules/gov"
	"github.com/stalwart-algoritmiclab/callisto/modules/mint"
	"github.com/stalwart-algoritmiclab/callisto/modules/slashing"
	"github.com/stalwart-algoritmiclab/callisto/modules/staking"
	modulestypes "github.com/stalwart-algoritmiclab/callisto/modules/types"
	"github.com/stalwart-algoritmiclab/callisto/utils"
)

func paramsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Get the current parameters of the gov module",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			cdc := utils.GetCodec()
			sources, err := modulestypes.BuildSources(config.Cfg.Node, cdc)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build expected modules of gov modules
			distrModule := distribution.NewModule(sources.DistrSource, cdc, db)
			mintModule := mint.NewModule(sources.MintSource, cdc, db)
			slashingModule := slashing.NewModule(sources.SlashingSource, cdc, db)
			stakingModule := staking.NewModule(sources.StakingSource, cdc, db)

			// Build the gov module
			govModule := gov.NewModule(sources.GovSource, distrModule, mintModule, slashingModule, stakingModule, cdc, db)

			height, err := parseCtx.Node.LatestHeight()
			if err != nil {
				return err
			}

			return govModule.UpdateParams(height)
		},
	}
}
