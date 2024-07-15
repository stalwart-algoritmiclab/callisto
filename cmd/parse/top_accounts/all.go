/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package top_accounts

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/parser"
	"github.com/forbole/juno/v6/types/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/modules/auth"
	"github.com/stalwart-algoritmiclab/callisto/modules/bank"
	"github.com/stalwart-algoritmiclab/callisto/modules/distribution"
	"github.com/stalwart-algoritmiclab/callisto/modules/staking"
	topaccounts "github.com/stalwart-algoritmiclab/callisto/modules/top_accounts"
	modulestypes "github.com/stalwart-algoritmiclab/callisto/modules/types"
	"github.com/stalwart-algoritmiclab/callisto/types"
	"github.com/stalwart-algoritmiclab/callisto/utils"
)

var (
	waitGroup sync.WaitGroup
)

const (
	flagWorker = "worker"
)

// allCmd returns the command that allows to parse all the top accounts
func allCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "all",
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

			// Build modules
			authModule := auth.NewModule(sources.AuthSource, nil, cdc, db)
			bankModule := bank.NewModule(nil, sources.BankSource, cdc, db)
			distriModule := distribution.NewModule(sources.DistrSource, cdc, db)
			stakingModule := staking.NewModule(sources.StakingSource, cdc, db)
			topaccountsModule := topaccounts.NewModule(authModule, sources.AuthSource, bankModule, distriModule, stakingModule, nil, cdc, parseCtx.Node, db)

			// Get workers
			exportQueue := NewQueue(5)
			workerCount, _ := cmd.Flags().GetInt64(flagWorker)
			workers := make([]Worker, workerCount)
			for i := range workers {
				workers[i] = NewWorker(exportQueue, topaccountsModule)
			}

			waitGroup.Add(1)

			// Query the latest chain height
			latestHeight, err := parseCtx.Node.LatestHeight()
			if err != nil {
				return fmt.Errorf("error while getting chain latest block height: %s", err)
			}

			// Set the height 5 blocks lower to avoid error
			// codespace sdk code 26: invalid height: cannot query with height in the future
			height := latestHeight - 5

			// Store all addresses in database
			accounts, err := authModule.RefreshTopAccountsList(height)
			if err != nil {
				return fmt.Errorf("error while unpacking accounts: %s", err)
			}

			for i, w := range workers {
				log.Debug().Int("number", i+1).Msg("starting worker...")
				go w.start()
			}

			trapSignal(parseCtx)

			go enqueueAddresses(exportQueue, accounts)

			waitGroup.Wait()
			return nil
		},
	}

	cmd.Flags().Int64(flagWorker, 1, "worker count")

	return cmd
}

// enqueueAddresses enqueues the given accounts' addresses into the given export queue
func enqueueAddresses(exportQueue AddressQueue, accounts []types.Account) {
	for _, account := range accounts {
		exportQueue <- account.Address
	}
}

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func trapSignal(ctx *parser.Context) {
	var sigCh = make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		sig := <-sigCh
		log.Info().Str("signal", sig.String()).Msg("caught signal; shutting down...")
		defer ctx.Node.Stop()
		defer ctx.Database.Close()
		defer waitGroup.Done()
	}()
}
