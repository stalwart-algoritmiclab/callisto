/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
)

// toMsgCreateTariffsDataBase - mapping model to db model
func toMsgCreateTariffsDataBase(tariff *types.MsgCreateTariffs) MsgCreateTariffs {
	return MsgCreateTariffs{
		Creator: tariff.Creator,
		Denom:   tariff.Denom,
		Tariffs: toTariffDatabase(tariff.Tariffs),
	}

}

// toMsgDeleteTariffsDataBase - mapping model to db model
func toMsgDeleteTariffsDataBase(tariff *types.MsgDeleteTariffs) MsgDeleteTariffs {
	return MsgDeleteTariffs{
		Creator:  tariff.Creator,
		Denom:    tariff.Denom,
		TariffID: tariff.TariffID,
		FeeID:    tariff.FeeID,
	}
}

// toMsgUpdateTariffsDataBase - mapping model to db model
func toMsgUpdateTariffsDataBase(tariff *types.MsgUpdateTariffs) MsgUpdateTariffs {
	return MsgUpdateTariffs{
		Creator: tariff.Creator,
		Denom:   tariff.Denom,
		Tariffs: toTariffDatabase(tariff.Tariffs),
	}
}

// toTariffDatabase - mapping model to db model
func toTariffDatabase(m *types.Tariff) Tariff {
	return Tariff{
		Denom:         m.Denom,
		TariffID:      m.Id,
		Amount:        m.Amount,
		MinRefBalance: m.MinRefBalance,
		Fees:          toFeesDatabaseList(m.Fees),
	}
}

// toFeesDatabaseList - mapping list to db list
func toFeesDatabaseList(m []*types.Fees) Fees {
	res := make(Fees, 0, len(m))
	for _, fee := range m {
		res = append(res, toFeesDatabase(fee))
	}

	return res
}

// toFeesDatabase - mapping model to db model
func toFeesDatabase(m *types.Fees) Fee {
	return Fee{
		AmountFrom:  m.AmountFrom,
		Fee:         m.Fee,
		RefReward:   m.RefReward,
		StakeReward: m.StakeReward,
		MinAmount:   m.MinAmount,
		NoRefReward: m.NoRefReward,
		Creator:     m.Creator,
		ID:          m.Id,
	}
}

// toMsgCreateAddressesDatabase - mapping model to db model
func toMsgCreateAddressesDatabase(hash string, m *types.MsgCreateAddresses) (MsgCreateAddresses, error) {
	return MsgCreateAddresses{
		Creator: m.Creator,
		Address: m.Address,
		TxHash:  hash,
	}, nil
}

// toMsgUpdateAddressesDatabase - mapping model to db model
func toMsgUpdateAddressesDatabase(hash string, m *types.MsgUpdateAddresses) (MsgUpdateAddresses, error) {
	return MsgUpdateAddresses{
		Creator:   m.Creator,
		Address:   m.Address,
		AddressID: m.Id,
		TxHash:    hash,
	}, nil
}

// toMsgDeleteAddressesDatabase - mapping model to db model
func toMsgDeleteAddressesDatabase(hash string, m *types.MsgDeleteAddresses) (MsgDeleteAddresses, error) {
	return MsgDeleteAddresses{
		Creator:   m.Creator,
		AddressID: m.Id,
		TxHash:    hash,
	}, nil
}
