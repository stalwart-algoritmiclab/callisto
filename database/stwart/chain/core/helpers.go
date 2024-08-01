/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

func toMsgSendDomain(m MsgSend) types.MsgSend {
	return types.MsgSend{
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
		Denom:   m.Denom,
	}
}

func toMsgSendDomainList(m []MsgSend) []types.MsgSend {
	res := make([]types.MsgSend, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgSendDomain(msg))
	}

	return res
}

func toMsgSendDatabase(txHash string, m *types.MsgSend) (MsgSend, error) {
	return MsgSend{
		TxHash:  txHash,
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
		Denom:   m.Denom,
	}, nil
}

func toMsgIssueDomain(m MsgIssue) types.MsgIssue {
	return types.MsgIssue{
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}
}

func toMsgIssueDomainList(m []MsgIssue) []types.MsgIssue {
	res := make([]types.MsgIssue, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgIssueDomain(msg))
	}

	return res
}

func toMsgIssueDatabase(txHash string, m *types.MsgIssue) (MsgIssue, error) {
	return MsgIssue{
		TxHash:  txHash,
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}, nil
}

func toMsgWithdrawDomain(m MsgWithdraw) types.MsgWithdraw {
	return types.MsgWithdraw{
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}
}

func toMsgWithdrawDomainList(m []MsgWithdraw) []types.MsgWithdraw {
	res := make([]types.MsgWithdraw, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgWithdrawDomain(msg))
	}

	return res
}

func toMsgWithdrawDatabase(txHash string, m *types.MsgWithdraw) (MsgWithdraw, error) {
	return MsgWithdraw{
		TxHash:  txHash,
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}, nil
}
