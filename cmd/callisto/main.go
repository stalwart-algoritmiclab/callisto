/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package main

import (
	"github.com/forbole/juno/v6/cmd"
	initcmd "github.com/forbole/juno/v6/cmd/init"
	parsetypes "github.com/forbole/juno/v6/cmd/parse/types"
	startcmd "github.com/forbole/juno/v6/cmd/start"
	"github.com/forbole/juno/v6/modules/messages"

	migratecmd "github.com/stalwart-algoritmiclab/callisto/cmd/migrate"
	parsecmd "github.com/stalwart-algoritmiclab/callisto/cmd/parse"
	"github.com/stalwart-algoritmiclab/callisto/types/config"
	"github.com/stalwart-algoritmiclab/callisto/utils"

	vault "github.com/stalwart-algoritmiclab/callisto/config"
	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/modules"
)

func main() {
	initCfg := initcmd.NewConfig().
		WithConfigCreator(config.Creator)

	cdc := utils.GetCodec()
	parseCfg := parsetypes.NewConfig().
		WithDBBuilder(database.Builder(cdc)).
		WithRegistrar(modules.NewRegistrar(getAddressesParser(), cdc))

	cfg := cmd.NewConfig("callisto").
		WithInitConfig(initCfg).
		WithParseConfig(parseCfg)

	// Run the command
	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.NewInitCmd(cfg.GetInitConfig()),
		vault.CheckVaultConfig(cfg.GetName(), startcmd.NewStartCmd(cfg.GetParseConfig())),
		parsecmd.NewParseCmd(cfg.GetParseConfig()),
		migratecmd.NewMigrateCmd(cfg.GetName(), cfg.GetParseConfig()),
		startcmd.NewStartCmd(cfg.GetParseConfig()),
	)

	executor := cmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

// getAddressesParser returns the messages parser that should be used to get the users involved in
// a specific message.
// This should be edited by custom implementations if needed.
func getAddressesParser() messages.MessageAddressesParser {
	return messages.JoinMessageParsers(
		messages.CosmosMessageAddressesParser,
	)
}
