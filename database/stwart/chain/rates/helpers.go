package rates

import (
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// toMsgCreateAddressesDomain - mapping db model to proto model
func toMsgCreateAddressesDomain(m MsgCreateAddresses) rates.MsgCreateAddresses {
	return rates.MsgCreateAddresses{
		Creator: m.Creator,
		Address: m.Address,
	}
}

// toMsgCreateAddressesDomainList - mapping func to a proto list.
func toMsgCreateAddressesDomainList(m []MsgCreateAddresses) []rates.MsgCreateAddresses {
	res := make([]rates.MsgCreateAddresses, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgCreateAddressesDomain(msg))
	}

	return res
}

// toMsgCreateAddressesDatabase - mapping func to a database model.
func toMsgCreateAddressesDatabase(txHash string, m *rates.MsgCreateAddresses) (MsgCreateAddresses, error) {
	return MsgCreateAddresses{
		Creator: m.Creator,
		Address: m.Address,
		TxHash:  txHash,
	}, nil
}

// toMsgCreateRatesDomain - mapping db model to proto model
func toMsgCreateRatesDomain(m MsgCreateRates) rates.MsgCreateRates {
	return rates.MsgCreateRates{
		Creator:  m.Creator,
		Denom:    m.Denom,
		Rate:     m.Rate,
		Decimals: m.Decimals,
	}
}

// toMsgCreateRatesDomainList - mapping func to a proto list.
func toMsgCreateRatesDomainList(m []MsgCreateRates) []rates.MsgCreateRates {
	res := make([]rates.MsgCreateRates, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgCreateRatesDomain(msg))
	}

	return res
}

// toMsgCreateAddressesDatabase - mapping func to a database model.
func toMsgCreateRatesDatabase(txHash string, m *rates.MsgCreateRates) (MsgCreateRates, error) {
	return MsgCreateRates{
		Creator:  m.Creator,
		Decimals: m.Decimals,
		Denom:    m.Denom,
		Rate:     m.Rate,
		TxHash:   txHash,
	}, nil
}

func toMsgUpdateRatesDomain(m MsgUpdateRates) rates.MsgUpdateRates {
	return rates.MsgUpdateRates{
		Creator: m.Creator,
		Denom:   m.Denom,
		Rate:    m.Rate,
	}
}

func toMsgUpdateRatesDomainList(m []MsgUpdateRates) []rates.MsgUpdateRates {
	res := make([]rates.MsgUpdateRates, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgUpdateRatesDomain(msg))
	}

	return res
}

func toMsgUpdateRatesDatabase(txHash string, m *rates.MsgUpdateRates) (MsgUpdateRates, error) {
	return MsgUpdateRates{
		Creator: m.Creator,
		Denom:   m.Denom,
		Rate:    m.Rate,
		TxHash:  txHash,
	}, nil
}

func toMsgDeleteRatesDomain(m MsgDeleteRates) rates.MsgDeleteRates {
	return rates.MsgDeleteRates{
		Creator: m.Creator,
		Denom:   m.Denom,
	}
}

func toMsgDeleteRatesDomainList(m []MsgDeleteRates) []rates.MsgDeleteRates {
	res := make([]rates.MsgDeleteRates, 0, len(m))
	for _, msg := range m {
		res = append(res, toMsgDeleteRatesDomain(msg))
	}

	return res
}

func toMsgDeleteRatesDatabase(txHash string, m *rates.MsgDeleteRates) (MsgDeleteRates, error) {
	return MsgDeleteRates{
		Creator: m.Creator,
		Denom:   m.Denom,
		TxHash:  txHash,
	}, nil
}
