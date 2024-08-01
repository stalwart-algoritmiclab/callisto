/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package actions

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/stalwart-algoritmiclab/callisto/modules/actions/handlers"
	actionstypes "github.com/stalwart-algoritmiclab/callisto/modules/actions/types"
)

var (
	waitGroup sync.WaitGroup
)

func (m *Module) RunAdditionalOperations() error {
	// Build the worker
	context := actionstypes.NewContext(m.node, m.sources)
	worker := actionstypes.NewActionsWorker(context)

	// Register the endpoints

	// -- Bank --
	worker.RegisterHandler("/account_balance", handlers.AccountBalanceHandler)

	// -- Distribution --
	worker.RegisterHandler("/delegation_reward", handlers.DelegationRewardHandler)
	worker.RegisterHandler("/delegator_withdraw_address", handlers.DelegatorWithdrawAddressHandler)
	worker.RegisterHandler("/validator_commission_amount", handlers.ValidatorCommissionAmountHandler)

	// -- Staking Delegator --
	worker.RegisterHandler("/delegation", handlers.DelegationHandler)
	worker.RegisterHandler("/delegation_total", handlers.TotalDelegationAmountHandler)
	worker.RegisterHandler("/unbonding_delegation", handlers.UnbondingDelegationsHandler)
	worker.RegisterHandler("/unbonding_delegation_total", handlers.UnbondingDelegationsTotal)
	worker.RegisterHandler("/redelegation", handlers.RedelegationHandler)

	// -- Staking Validator --
	worker.RegisterHandler("/validator_delegations", handlers.ValidatorDelegation)
	worker.RegisterHandler("/validator_redelegations_from", handlers.ValidatorRedelegationsFromHandler)
	worker.RegisterHandler("/validator_unbonding_delegations", handlers.ValidatorUnbondingDelegationsHandler)

	// -- StalwartChain Polls --
	worker.RegisterHandler("/stalwart_polls", handlers.StalwartChainPollsHandler)

	// Listen for and trap any OS signal to gracefully shutdown and exit
	m.trapSignal()

	// Start the worker
	waitGroup.Add(1)
	go worker.Start(m.cfg.Host, m.cfg.Port)

	// Block main process (signal capture will call WaitGroup's Done)
	waitGroup.Wait()
	return nil
}

// trapSignal will listen for any OS signal and invoke Done on the main
// WaitGroup allowing the main process to gracefully exit.
func (m *Module) trapSignal() {
	var sigCh = make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGTERM)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		defer m.node.Stop()
		defer waitGroup.Done()
	}()
}
