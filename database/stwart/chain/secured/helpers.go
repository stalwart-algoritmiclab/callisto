/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

// MsgCreateAddresses - mapping db model to model
func toMsgCreateAddressesDomain(m MsgCreateAddresses) types.MsgCreateAddresses {
	return types.MsgCreateAddresses{
		Creator: m.Creator,
		Address: m.Addresses,
	}
}

// toMsgCreateAddressesDomainList - mapping func to a list.
func toMsgCreateAddressesDomainList(m []MsgCreateAddresses) []types.MsgCreateAddresses {
	res := make([]types.MsgCreateAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgCreateAddressesDomain(msg))
	}

	return res
}

// toMsgCreateAddressesDatabase - mapping func to a database model.
func toMsgCreateAddressesDatabase(txHash string, m *types.MsgCreateAddresses) (MsgCreateAddresses, error) {
	return MsgCreateAddresses{
		Creator:   m.Creator,
		Addresses: m.Address,
		TxHash:    txHash,
	}, nil
}

// MsgUpdateAddresses - mapping db model to model
func toMsgUpdateAddressesDomain(m MsgUpdateAddresses) types.MsgUpdateAddresses {
	return types.MsgUpdateAddresses{
		Creator: m.Creator,
		Id:      m.AddressID,
		Address: m.Addresses,
	}
}

// toMsgUpdateAddressesDomainList - mapping func to a list.
func toMsgUpdateAddressesDomainList(m []MsgUpdateAddresses) []types.MsgUpdateAddresses {
	res := make([]types.MsgUpdateAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgUpdateAddressesDomain(msg))
	}

	return res
}

// toMsgUpdateAddressesDatabase - mapping func to a database model.
func toMsgUpdateAddressesDatabase(txHash string, m *types.MsgUpdateAddresses) (MsgUpdateAddresses, error) {
	return MsgUpdateAddresses{
		AddressID: m.Id,
		Creator:   m.Creator,
		Addresses: m.Address,
		TxHash:    txHash,
	}, nil
}

// MsgDeleteAddresses - mapping db model to model
func toMsgDeleteAddressesDomain(m MsgDeleteAddresses) types.MsgDeleteAddresses {
	return types.MsgDeleteAddresses{
		Id:      m.AddressID,
		Creator: m.Creator,
	}
}

// toMsgDeleteAddressesDomainList - mapping func to a list.
func toMsgDeleteAddressesDomainList(m []MsgDeleteAddresses) []types.MsgDeleteAddresses {
	res := make([]types.MsgDeleteAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgDeleteAddressesDomain(msg))
	}

	return res
}

// toMsgDeleteAddressesDatabase - mapping func to a database model.
func toMsgDeleteAddressesDatabase(txHash string, m *types.MsgDeleteAddresses) (MsgDeleteAddresses, error) {
	return MsgDeleteAddresses{
		AddressID: m.Id,
		Creator:   m.Creator,
		TxHash:    txHash,
	}, nil
}
