/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package exchanger

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
)

// MsgExchange - mapping db model to model
func toMsgExchangeDomain(m MsgExchange) types.MsgExchange {
	return types.MsgExchange{
		Creator: m.Creator,
		Denom:   m.Denom,
		Amount:  m.Amount,
		DenomTo: m.DenomTo,
	}
}

// toMsgExchangeDomainList - mapping func to a list.
func toMsgExchangeDomainList(m []MsgExchange) []types.MsgExchange {
	res := make([]types.MsgExchange, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgExchangeDomain(msg))
	}

	return res
}

// toMsgExchangeDatabase - mapping func to a database model.
func toMsgExchangeDatabase(txHash string, m *types.MsgExchange) (MsgExchange, error) {
	return MsgExchange{
		TxHash:  txHash,
		Creator: m.Creator,
		Denom:   m.Denom,
		Amount:  m.Amount,
		DenomTo: m.DenomTo,
	}, nil
}
