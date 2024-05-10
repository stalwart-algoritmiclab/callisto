package bank

import (
	"fmt"

	modulestypes "github.com/stalwart-algoritmiclab/callisto/modules/types"
	"github.com/stalwart-algoritmiclab/callisto/utils"

	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/types/config"
	"github.com/spf13/cobra"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/modules/bank"
)

// supplyCmd returns the Cobra command allowing to refresh x/bank total supply
func supplyCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "supply",
		Short: "Refresh total supply",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			sources, err := modulestypes.BuildSources(config.Cfg.Node, utils.GetCodec())
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build bank module
			bankModule := bank.NewModule(nil, sources.BankSource, utils.GetCodec(), db)

			err = bankModule.UpdateSupply()
			if err != nil {
				return fmt.Errorf("error while getting latest bank supply: %s", err)
			}

			return nil
		},
	}
}
