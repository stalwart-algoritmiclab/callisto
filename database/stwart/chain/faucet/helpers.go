package faucet

import (
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

// MsgIssue - mapping db model to proto model
func toMsgIssueDomain(m MsgIssue) faucet.MsgIssue {
	return faucet.MsgIssue{
		Creator: m.Creator,
		Address: m.Address,
	}
}

// toMsgIssueDomainList - mapping func to a proto list.
func toMsgIssueDomainList(m []MsgIssue) []faucet.MsgIssue {
	res := make([]faucet.MsgIssue, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgIssueDomain(msg))
	}

	return res
}

// toMsgIssueDatabase - mapping func to a database model.
func toMsgIssueDatabase(txHash string, m *faucet.MsgIssue) (MsgIssue, error) {
	return MsgIssue{
		Creator: m.Creator,
		Address: m.Address,
		TxHash:  txHash,
	}, nil
}
