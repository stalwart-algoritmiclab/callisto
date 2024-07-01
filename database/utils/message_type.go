/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package utils

import (
	"github.com/stalwart-algoritmiclab/callisto/modules/utils"
	"github.com/stalwart-algoritmiclab/callisto/types"
)

// Stalwart module types
const (
	// Stalwart feepolicy module event types
	Stwartchain_feepolicy_MsgCreateTariffs = "/stwartchain.feepolicy.MsgCreateTariffs"
	Stwartchain_feepolicy_MsgUpdateTariffs = "/stwartchain.feepolicy.MsgUpdateTariffs"
	Stwartchain_feepolicy_MsgDeleteTariffs = "/stwartchain.feepolicy.MsgDeleteTariffs"

	// Stalwart stats module event types

	// Stalwart core module event types
	Stwartchain_core_MsgWithdraw = "/stwartchain.core.MsgWithdraw"
	Stwartchain_core_MsgIssue    = "/stwartchain.core.MsgIssue"

	// Stalwart exchanger module event types
	Stwartchain_exchanger_MsgExchange = "/stwartchain.exchanger.MsgExchange"

	// Stalwart faucet module event types
	Stwartchain_faucet_MsgIssue = "/stwartchain.faucet.MsgIssue"

	// Stalwart rates module event types

	// Stalwart referrals module event types

	// Stalwart secured module event types

	// Stalwart stake module event types

	// Stalwart module event types

	// Stalwart system rewards module event types

	// Stalwart users module event types
)

// Cosmos module types
const (
	// Cosmos slashing module event types
	Cosmos_slashing_MsgUnjail = "/cosmos.slashing.v1beta1.MsgUnjail"

	// Cosmos bank module event types
	Cosmos_bank_MsgSend = "/cosmos.bank.v1beta1.MsgSend"

	// IBC core module event types
	Ibc_core_client_v1_MsgUpdateClient = "/ibc.core.client.v1.MsgUpdateClient"
)

// MessageTypeLists - return list of message types
func MessageTypeLists() []types.MessageType {
	return []types.MessageType{
		{
			Type:   Cosmos_slashing_MsgUnjail,
			Module: utils.GetModuleNameFromTypeURL(Cosmos_slashing_MsgUnjail),
			Label:  utils.GetMsgFromTypeURL(Cosmos_slashing_MsgUnjail),
			Height: 129,
		},
		{
			Type:   Stwartchain_exchanger_MsgExchange,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_exchanger_MsgExchange),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_exchanger_MsgExchange),
			Height: 253,
		},
		{
			Type:   Cosmos_bank_MsgSend,
			Module: utils.GetModuleNameFromTypeURL(Cosmos_bank_MsgSend),
			Label:  utils.GetMsgFromTypeURL(Cosmos_bank_MsgSend),
			Height: 247,
		},
		{
			Type:   Stwartchain_faucet_MsgIssue,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_faucet_MsgIssue),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_faucet_MsgIssue),
			Height: 251,
		},
		{
			Type:   Stwartchain_core_MsgWithdraw,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_core_MsgWithdraw),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_core_MsgWithdraw),
			Height: 114047,
		},
		{
			Type:   Ibc_core_client_v1_MsgUpdateClient,
			Module: utils.GetModuleNameFromTypeURL(Ibc_core_client_v1_MsgUpdateClient),
			Label:  utils.GetMsgFromTypeURL(Ibc_core_client_v1_MsgUpdateClient),
			Height: 18734,
		},
		{
			Type:   Stwartchain_feepolicy_MsgCreateTariffs,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_feepolicy_MsgCreateTariffs),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_feepolicy_MsgCreateTariffs),
			Height: 61445,
		},
		{
			Type:   Stwartchain_feepolicy_MsgUpdateTariffs,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_feepolicy_MsgUpdateTariffs),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_feepolicy_MsgUpdateTariffs),
			Height: 71648,
		},
		{
			Type:   Stwartchain_feepolicy_MsgDeleteTariffs,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_feepolicy_MsgDeleteTariffs),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_feepolicy_MsgDeleteTariffs),
			Height: 71682,
		},
		{
			Type:   Stwartchain_core_MsgIssue,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_core_MsgIssue),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_core_MsgIssue),
			Height: 81413,
		},
	}
}
