/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

func toMsgRefRewardDomain(m MsgRefReward) types.MsgRefReward {
	coin, err := sdk.ParseCoinNormalized(m.Amount)
	if err != nil {
		return types.MsgRefReward{}
	}
	return types.MsgRefReward{
		Creator:  m.Creator,
		Amount:   coin,
		Referrer: m.Referrer,
	}
}

func toMsgRefRewardDomainList(m []MsgRefReward) []types.MsgRefReward {
	res := make([]types.MsgRefReward, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgRefRewardDomain(msg))
	}

	return res

}

func toMsgRefRewardDatabase(txHash string, m *types.MsgRefReward) (MsgRefReward, error) {
	amount := fmt.Sprintf("{%s}{%s}", m.Amount.Amount.String(), m.Amount.Denom)
	return MsgRefReward{
		TxHash:   txHash,
		Creator:  m.Creator,
		Amount:   amount,
		Referrer: m.Referrer,
	}, nil
}

func toMsgFeesDomain(m MsgFees) types.MsgFees {
	coin, err := sdk.ParseCoinNormalized(m.Commission)
	if err != nil {
		return types.MsgFees{}
	}
	return types.MsgFees{
		Creator:   m.Creator,
		Comission: coin,
		AddressTo: m.Address,
	}
}

func toMsgFeesDomainList(m []MsgFees) []types.MsgFees {
	res := make([]types.MsgFees, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgFeesDomain(msg))
	}

	return res
}

func toMsgFeesDatabase(txHash string, m *types.MsgFees) (MsgFees, error) {
	commission := fmt.Sprintf("{%s}{%s}", m.Comission.Amount.String(), m.Comission.Denom)
	return MsgFees{
		TxHash:     txHash,
		Creator:    m.Creator,
		Commission: commission,
		Address:    m.AddressTo,
	}, nil

}

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

func toMsgRefundDomain(m MsgRefund) types.MsgRefund {
	return types.MsgRefund{
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
	}
}

func toMsgRefundDomainList(m []MsgRefund) []types.MsgRefund {
	res := make([]types.MsgRefund, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgRefundDomain(msg))
	}

	return res
}

func toMsgRefundDatabase(txHash string, m *types.MsgRefund) (MsgRefund, error) {
	return MsgRefund{
		TxHash:  txHash,
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
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
