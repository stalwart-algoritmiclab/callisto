package secured

import (
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

// MsgCreateAddresses - mapping db model to proto model
func toMsgCreateAddressesDomain(m MsgCreateAddresses) secured.MsgCreateAddresses {
	return secured.MsgCreateAddresses{
		Creator: m.Creator,
		Address: m.Addresses,
	}
}

// toMsgCreateAddressesDomainList - mapping func to a proto list.
func toMsgCreateAddressesDomainList(m []MsgCreateAddresses) []secured.MsgCreateAddresses {
	res := make([]secured.MsgCreateAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgCreateAddressesDomain(msg))
	}

	return res
}

// toMsgCreateAddressesDatabase - mapping func to a database model.
func toMsgCreateAddressesDatabase(txHash string, m *secured.MsgCreateAddresses) (MsgCreateAddresses, error) {
	return MsgCreateAddresses{
		Creator:   m.Creator,
		Addresses: m.Address,
		TxHash:    txHash,
	}, nil
}
