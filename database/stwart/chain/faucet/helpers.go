/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package faucet

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
)

// MsgIssue - mapping db model to model
func toMsgIssueDomain(m MsgIssue) types.MsgIssue {
	return types.MsgIssue{
		Creator: m.Creator,
		Address: m.Address,
	}
}

// toMsgIssueDomainList - mapping func to a list.
func toMsgIssueDomainList(m []MsgIssue) []types.MsgIssue {
	res := make([]types.MsgIssue, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgIssueDomain(msg))
	}

	return res
}

// toMsgIssueDatabase - mapping func to a database model.
func toMsgIssueDatabase(txHash string, m *types.MsgIssue) (MsgIssue, error) {
	return MsgIssue{
		Creator: m.Creator,
		Address: m.Address,
		TxHash:  txHash,
	}, nil
}
