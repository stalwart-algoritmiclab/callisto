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
	Stwartchain_secured_MsgCreateAddresses = "/stwartchain.secured.MsgCreateAddresses"

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
			Type:   Stwartchain_secured_MsgCreateAddresses,
			Module: utils.GetModuleNameFromTypeURL(Stwartchain_secured_MsgCreateAddresses),
			Label:  utils.GetMsgFromTypeURL(Stwartchain_secured_MsgCreateAddresses),
			Height: 76,
		},
	}
}
