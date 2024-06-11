/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/core"
)

func toMsgRefRewardDomain(m MsgRefReward) core.MsgRefReward {
	coin, err := sdk.ParseCoinNormalized(m.Amount)
	if err != nil {
		return core.MsgRefReward{}
	}
	return core.MsgRefReward{
		Creator:  m.Creator,
		Amount:   coin,
		Referrer: m.Referrer,
	}
}

func toMsgRefRewardDomainList(m []MsgRefReward) []core.MsgRefReward {
	res := make([]core.MsgRefReward, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgRefRewardDomain(msg))
	}

	return res

}

func toMsgRefRewardDatabase(txHash string, m *core.MsgRefReward) (MsgRefReward, error) {
	amount := fmt.Sprintf("{%s}{%s}", m.Amount.Amount.String(), m.Amount.Denom)
	return MsgRefReward{
		TxHash:   txHash,
		Creator:  m.Creator,
		Amount:   amount,
		Referrer: m.Referrer,
	}, nil
}

func toMsgFeesDomain(m MsgFees) core.MsgFees {
	coin, err := sdk.ParseCoinNormalized(m.Commission)
	if err != nil {
		return core.MsgFees{}
	}
	return core.MsgFees{
		Creator:   m.Creator,
		Comission: coin,
		AddressTo: m.Address,
	}
}

func toMsgFeesDomainList(m []MsgFees) []core.MsgFees {
	res := make([]core.MsgFees, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgFeesDomain(msg))
	}

	return res
}

func toMsgFeesDatabase(txHash string, m *core.MsgFees) (MsgFees, error) {
	commission := fmt.Sprintf("{%s}{%s}", m.Comission.Amount.String(), m.Comission.Denom)
	return MsgFees{
		TxHash:     txHash,
		Creator:    m.Creator,
		Commission: commission,
		Address:    m.AddressTo,
	}, nil

}

func toMsgSendDomain(m MsgSend) core.MsgSend {
	return core.MsgSend{
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
		Denom:   m.Denom,
	}
}

func toMsgSendDomainList(m []MsgSend) []core.MsgSend {
	res := make([]core.MsgSend, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgSendDomain(msg))
	}

	return res
}

func toMsgSendDatabase(txHash string, m *core.MsgSend) (MsgSend, error) {
	return MsgSend{
		TxHash:  txHash,
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
		Denom:   m.Denom,
	}, nil
}

func toMsgRefundDomain(m MsgRefund) core.MsgRefund {
	return core.MsgRefund{
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
	}
}

func toMsgRefundDomainList(m []MsgRefund) []core.MsgRefund {
	res := make([]core.MsgRefund, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgRefundDomain(msg))
	}

	return res
}

func toMsgRefundDatabase(txHash string, m *core.MsgRefund) (MsgRefund, error) {
	return MsgRefund{
		TxHash:  txHash,
		Creator: m.Creator,
		From:    m.From,
		To:      m.To,
		Amount:  m.Amount,
	}, nil
}

func toMsgIssueDomain(m MsgIssue) core.MsgIssue {
	return core.MsgIssue{
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}
}

func toMsgIssueDomainList(m []MsgIssue) []core.MsgIssue {
	res := make([]core.MsgIssue, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgIssueDomain(msg))
	}

	return res
}

func toMsgIssueDatabase(txHash string, m *core.MsgIssue) (MsgIssue, error) {
	return MsgIssue{
		TxHash:  txHash,
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}, nil
}

func toMsgWithdrawDomain(m MsgWithdraw) core.MsgWithdraw {
	return core.MsgWithdraw{
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}
}

func toMsgWithdrawDomainList(m []MsgWithdraw) []core.MsgWithdraw {
	res := make([]core.MsgWithdraw, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgWithdrawDomain(msg))
	}

	return res
}

func toMsgWithdrawDatabase(txHash string, m *core.MsgWithdraw) (MsgWithdraw, error) {
	return MsgWithdraw{
		TxHash:  txHash,
		Creator: m.Creator,
		Amount:  m.Amount,
		Denom:   m.Denom,
		Address: m.Address,
	}, nil
}
