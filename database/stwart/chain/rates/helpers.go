/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

// toMsgCreateAddressesDomain - mapping db model to model
func toMsgCreateAddressesDomain(m MsgCreateAddresses) types.MsgCreateAddresses {
	return types.MsgCreateAddresses{
		Creator: m.Creator,
		Address: m.Address,
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
		Creator: m.Creator,
		Address: m.Address,
		TxHash:  txHash,
	}, nil
}

// toMsgUpdateAddressesDomain - mapping db model to model
func toMsgUpdateAddressesDomain(m MsgUpdateAddresses) types.MsgUpdateAddresses {
	return types.MsgUpdateAddresses{
		Creator: m.Creator,
		Id:      m.ID,
		Address: m.Address,
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
		Creator:   m.Creator,
		Address:   m.Address,
		AddressID: m.Id,
		TxHash:    txHash,
	}, nil
}

// toMsgDeleteAddressesDomain - mapping db model to model
func toMsgDeleteAddressesDomain(m MsgDeleteAddresses) types.MsgDeleteAddresses {
	return types.MsgDeleteAddresses{
		Creator: m.Creator,
		Id:      m.AddressID,
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
		Creator:   m.Creator,
		AddressID: m.Id,
		TxHash:    txHash,
	}, nil
}

// toMsgCreateRatesDomain - mapping db model to model
func toMsgCreateRatesDomain(m MsgCreateRates) types.MsgCreateRates {
	return types.MsgCreateRates{
		Creator:  m.Creator,
		Denom:    m.Denom,
		Rate:     m.Rate,
		Decimals: m.Decimals,
	}
}

// toMsgCreateRatesDomainList - mapping func to a list.
func toMsgCreateRatesDomainList(m []MsgCreateRates) []types.MsgCreateRates {
	res := make([]types.MsgCreateRates, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgCreateRatesDomain(msg))
	}

	return res
}

// toMsgCreateAddressesDatabase - mapping func to a database model.
func toMsgCreateRatesDatabase(txHash string, m *types.MsgCreateRates) (MsgCreateRates, error) {
	return MsgCreateRates{
		Creator:  m.Creator,
		Decimals: m.Decimals,
		Denom:    m.Denom,
		Rate:     m.Rate,
		TxHash:   txHash,
	}, nil
}

func toMsgUpdateRatesDomain(m MsgUpdateRates) types.MsgUpdateRates {
	return types.MsgUpdateRates{
		Creator: m.Creator,
		Denom:   m.Denom,
		Rate:    m.Rate,
	}
}

func toMsgUpdateRatesDomainList(m []MsgUpdateRates) []types.MsgUpdateRates {
	res := make([]types.MsgUpdateRates, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgUpdateRatesDomain(msg))
	}

	return res
}

func toMsgUpdateRatesDatabase(txHash string, m *types.MsgUpdateRates) (MsgUpdateRates, error) {
	return MsgUpdateRates{
		Creator: m.Creator,
		Denom:   m.Denom,
		Rate:    m.Rate,
		TxHash:  txHash,
	}, nil
}

func toMsgDeleteRatesDomain(m MsgDeleteRates) types.MsgDeleteRates {
	return types.MsgDeleteRates{
		Creator: m.Creator,
		Denom:   m.Denom,
	}
}

func toMsgDeleteRatesDomainList(m []MsgDeleteRates) []types.MsgDeleteRates {
	res := make([]types.MsgDeleteRates, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgDeleteRatesDomain(msg))
	}

	return res
}

func toMsgDeleteRatesDatabase(txHash string, m *types.MsgDeleteRates) (MsgDeleteRates, error) {
	return MsgDeleteRates{
		Creator: m.Creator,
		Denom:   m.Denom,
		TxHash:  txHash,
	}, nil
}
