/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

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

// MsgUpdateAddresses - mapping db model to proto model
func toMsgUpdateAddressesDomain(m MsgUpdateAddresses) secured.MsgUpdateAddresses {
	return secured.MsgUpdateAddresses{
		Creator: m.Creator,
		Id:      m.AddressID,
		Address: m.Addresses,
	}
}

// toMsgUpdateAddressesDomainList - mapping func to a proto list.
func toMsgUpdateAddressesDomainList(m []MsgUpdateAddresses) []secured.MsgUpdateAddresses {
	res := make([]secured.MsgUpdateAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgUpdateAddressesDomain(msg))
	}

	return res
}

// toMsgUpdateAddressesDatabase - mapping func to a database model.
func toMsgUpdateAddressesDatabase(txHash string, m *secured.MsgUpdateAddresses) (MsgUpdateAddresses, error) {
	return MsgUpdateAddresses{
		AddressID: m.Id,
		Creator:   m.Creator,
		Addresses: m.Address,
		TxHash:    txHash,
	}, nil
}

// MsgDeleteAddresses - mapping db model to proto model
func toMsgDeleteAddressesDomain(m MsgDeleteAddresses) secured.MsgDeleteAddresses {
	return secured.MsgDeleteAddresses{
		Id:      m.AddressID,
		Creator: m.Creator,
	}
}

// toMsgDeleteAddressesDomainList - mapping func to a proto list.
func toMsgDeleteAddressesDomainList(m []MsgDeleteAddresses) []secured.MsgDeleteAddresses {
	res := make([]secured.MsgDeleteAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgDeleteAddressesDomain(msg))
	}

	return res
}

// toMsgDeleteAddressesDatabase - mapping func to a database model.
func toMsgDeleteAddressesDatabase(txHash string, m *secured.MsgDeleteAddresses) (MsgDeleteAddresses, error) {
	return MsgDeleteAddresses{
		AddressID: m.Id,
		Creator:   m.Creator,
		TxHash:    txHash,
	}, nil
}
