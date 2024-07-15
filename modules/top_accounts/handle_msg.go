/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package top_accounts

import (
	"fmt"

	distritypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v6/types"

	moduleutils "github.com/stalwart-algoritmiclab/callisto/modules/utils"
	"github.com/stalwart-algoritmiclab/callisto/utils"
)

// msgFilter contains the list of messages that should be handled by the module
var msgFilter = map[string]bool{
	"/cosmos.staking.v1beta1.MsgDelegate":                     true,
	"/cosmos.staking.v1beta1.MsgUndelegate":                   true,
	"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation":    true,
	"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg types.Message, tx *types.Transaction) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	// Refresh x/bank available account balances
	addresses, err := m.messageParser(tx)
	if err != nil {
		return fmt.Errorf("error while parsing account addresses of message type %s: %s", msg.GetType(), err)
	}

	addresses = moduleutils.FilterNonAccountAddresses(addresses)
	err = m.bankModule.UpdateBalances(addresses, int64(tx.Height))
	if err != nil {
		return fmt.Errorf("error while updating account available balances: %s", err)
	}

	err = m.refreshTopAccountsSum(addresses, int64(tx.Height))
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum while refreshing balance: %s", err)
	}

	// Handle x/staking delegations and unbondings
	switch msg.GetType() {

	case "/cosmos.staking.v1beta1.MsgDelegate":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &stakingtypes.MsgDelegate{})
		return m.handleMsgDelegate(cosmosMsg.DelegatorAddress, int64(tx.Height))

	case "/cosmos.staking.v1beta1.MsgUndelegate":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &stakingtypes.MsgUndelegate{})
		return m.handleMsgUndelegate(cosmosMsg.DelegatorAddress, int64(tx.Height))

	case "/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &stakingtypes.MsgCancelUnbondingDelegation{})
		return m.handleMsgCancelUnbondingDelegation(cosmosMsg.DelegatorAddress, int64(tx.Height))

	// Handle x/distribution delegator rewards
	case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &distritypes.MsgWithdrawDelegatorReward{})
		return m.handleMsgWithdrawDelegatorReward(cosmosMsg.DelegatorAddress, int64(tx.Height))

	}

	return nil
}

func (m *Module) handleMsgDelegate(delAddr string, height int64) error {
	err := m.stakingModule.RefreshDelegations(delAddr, height)
	if err != nil {
		return fmt.Errorf("error while refreshing delegations while handling MsgDelegate: %s", err)
	}

	err = m.refreshTopAccountsSum([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum while handling MsgDelegate: %s", err)
	}

	return nil
}

// handleMsgUndelegate handles a MsgUndelegate storing the data inside the database
func (m *Module) handleMsgUndelegate(delAddr string, height int64) error {
	err := m.stakingModule.RefreshUnbondings(delAddr, height)
	if err != nil {
		return fmt.Errorf("error while refreshing undelegations while handling MsgUndelegate: %s", err)
	}

	err = m.refreshTopAccountsSum([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum while handling MsgUndelegate: %s", err)
	}

	return nil
}

// handleMsgCancelUnbondingDelegation handles a MsgCancelUnbondingDelegation storing the data inside the database
func (m *Module) handleMsgCancelUnbondingDelegation(delAddr string, height int64) error {
	err := m.stakingModule.RefreshDelegations(delAddr, height)
	if err != nil {
		return fmt.Errorf("error while refreshing delegations of account %s, error: %s", delAddr, err)
	}

	err = m.stakingModule.RefreshUnbondings(delAddr, height)
	if err != nil {
		return fmt.Errorf("error while refreshing unbonding delegations of account %s, error: %s", delAddr, err)
	}

	err = m.bankModule.UpdateBalances([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while refreshing balance of account %s, error: %s", delAddr, err)
	}

	err = m.refreshTopAccountsSum([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum %s, error: %s", delAddr, err)
	}

	return nil
}

// handleMsgWithdrawDelegatorReward handles a MsgWithdrawDelegatorReward storing the data inside the database
func (m *Module) handleMsgWithdrawDelegatorReward(delAddr string, height int64) error {
	err := m.distrModule.RefreshDelegatorRewards([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while refreshing delegator rewards: %s", err)
	}

	err = m.bankModule.UpdateBalances([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while updating account available balances with MsgWithdrawDelegatorReward: %s", err)
	}

	err = m.refreshTopAccountsSum([]string{delAddr}, height)
	if err != nil {
		return fmt.Errorf("error while refreshing top accounts sum while handling MsgWithdrawDelegatorReward: %s", err)
	}

	return nil
}
