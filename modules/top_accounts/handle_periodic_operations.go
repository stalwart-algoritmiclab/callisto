/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package top_accounts

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/modules/utils"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "top accounts").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(1).Day().At("00:00").Do(func() {
		utils.WatchMethod(m.RefreshTotalAccounts)
	}); err != nil {
		return fmt.Errorf("error while setting up refresh total top accounts periodic operation: %s", err)
	}
	utils.WatchMethod(m.RefreshTotalAccounts)

	if _, err := scheduler.Every(1).Day().At("00:01").Do(func() {
		utils.WatchMethod(m.RefreshTopAccountsList)
	}); err != nil {
		return fmt.Errorf("error while setting up refresh top accounts list periodic operation: %s", err)
	}
	utils.WatchMethod(m.RefreshTopAccountsList)

	if _, err := scheduler.Every(1).Day().At("00:02").Do(func() {
		utils.WatchMethod(m.RefreshAvailableBalance)
	}); err != nil {
		return fmt.Errorf("error while setting up refresh top accounts avail balance periodic operation: %s", err)
	}
	utils.WatchMethod(m.RefreshAvailableBalance)

	if _, err := scheduler.Every(1).Day().At("00:03").Do(func() {
		utils.WatchMethod(m.RefreshRewards)
	}); err != nil {
		return fmt.Errorf("error while setting up refresh top accounts rewards periodic operation: %s", err)
	}
	utils.WatchMethod(m.RefreshRewards)

	if _, err := scheduler.Every(1).Day().At("00:04").Do(func() {
		utils.WatchMethod(m.RefreshSSC)
	}); err != nil {
		return fmt.Errorf("error while setting up refresh top accounts SSC periodic operation: %s", err)
	}
	utils.WatchMethod(m.RefreshSSC)

	return nil
}

// RefreshTotalAccounts refreshes total number of accounts/wallets in database
func (m *Module) RefreshTotalAccounts() error {
	log.Trace().Str("module", "top accounts").Str("operation", "refresh total accounts").
		Msg("refreshing number of all wallets")

	height, err := m.node.LatestHeight()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	totalAccountsNumber, err := m.authSource.GetTotalNumberOfAccounts(height)
	if err != nil {
		return fmt.Errorf("error while getting total number of accounts: %s", err)
	}

	err = m.db.SaveTotalAccounts(int64(totalAccountsNumber), height)
	if err != nil {
		return fmt.Errorf("error while storing total number of accounts: %s", err)
	}

	return nil
}

// RefreshAvailableBalance refreshes latest available balance in db
func (m *Module) RefreshAvailableBalance() error {
	log.Trace().Str("module", "top accounts").Str("operation", "refresh available balance").
		Msg("refreshing available balance")

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	accounts, err := m.authModule.GetAllBaseAccounts(height)
	if err != nil {
		return fmt.Errorf("error while getting base accounts: %s", err)
	}

	if len(accounts) == 0 {
		return nil
	}

	// Store accounts
	err = m.db.SaveAccounts(accounts)
	if err != nil {
		return err
	}

	// Parse addresses to []string
	var addresses []string
	for _, a := range accounts {
		addresses = append(addresses, a.Address)
	}

	err = m.bankModule.UpdateBalances(addresses, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts balances, error: %s", err)
	}

	err = m.refreshTopAccountsSum(addresses, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum value: %s", err)
	}

	return nil
}

// RefreshRewards refreshes the rewards for all delegators
func (m *Module) RefreshRewards() error {
	time.Sleep(3 * time.Second)

	log.Trace().Str("module", "top accounts").Str("operation", "refresh rewards").
		Msg("refreshing delegators rewards")

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	// Get the delegators
	delegators, err := m.db.GetDelegators()
	if err != nil {
		return fmt.Errorf("error while getting delegators: %s", err)
	}

	if len(delegators) == 0 {
		return nil
	}

	// Refresh rewards
	err = m.distrModule.RefreshDelegatorRewards(delegators, height)
	if err != nil {
		return fmt.Errorf("error while refreshing delegators rewards: %s", err)
	}

	err = m.refreshTopAccountsSum(delegators, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum value: %s", err)
	}

	return nil
}

// RefreshSSC refreshes the token balances for all SSCs
func (m *Module) RefreshSSC() error {
	time.Sleep(3 * time.Second)

	log.Trace().Str("module", "top accounts").Str("operation ", "refresh SSC").
		Msg("refreshing SSC balances")

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	accounts, err := m.authModule.GetAllBaseAccounts(height)
	if err != nil {
		return fmt.Errorf("error while getting base accounts: %s", err)
	}

	if len(accounts) == 0 {
		return nil
	}

	// Store accounts
	err = m.db.SaveAccounts(accounts)
	if err != nil {
		return err
	}

	// Parse addresses to []string
	var addresses []string
	for _, a := range accounts {
		addresses = append(addresses, a.Address)
	}

	err = m.bankModule.UpdateSSCBalances(addresses, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts balances: %s", err)
	}

	return nil
}

// RefreshTopAccountsList refreshes top accounts list in db
func (m *Module) RefreshTopAccountsList() error {
	log.Trace().Str("module", "top accounts").Str("operation", "refresh top accounts list").
		Msg("refreshing top accounts list")

	height, err := m.db.GetLastBlockHeight()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	// Unpack all accounts into types.Account type
	_, err = m.authModule.RefreshTopAccountsList(height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts list: %s", err)
	}

	return nil
}
