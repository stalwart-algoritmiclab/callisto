/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"

// toMsgCreateTariffsDataBase - mapping proto model to db model
func toMsgCreateTariffsDataBase(tariff *feepolicy.MsgCreateTariffs) MsgCreateTariffs {
	return MsgCreateTariffs{
		Creator: tariff.Creator,
		Denom:   tariff.Denom,
		Tariffs: toTariffDatabase(tariff.Tariffs),
	}

}

// toMsgDeleteTariffsDataBase - mapping proto model to db model
func toMsgDeleteTariffsDataBase(tariff *feepolicy.MsgDeleteTariffs) MsgDeleteTariffs {
	return MsgDeleteTariffs{
		Creator:  tariff.Creator,
		Denom:    tariff.Denom,
		TariffID: tariff.TariffID,
		FeeID:    tariff.FeeID,
	}
}

// toMsgUpdateTariffsDataBase - mapping proto model to db model
func toMsgUpdateTariffsDataBase(tariff *feepolicy.MsgUpdateTariffs) MsgUpdateTariffs {
	return MsgUpdateTariffs{
		Creator: tariff.Creator,
		Denom:   tariff.Denom,
		Tariffs: toTariffDatabase(tariff.Tariffs),
	}
}

// toTariffDatabase - mapping proto model to db model
func toTariffDatabase(m *feepolicy.Tariff) Tariff {
	return Tariff{
		Denom:         m.Denom,
		TariffID:      m.Id,
		Amount:        m.Amount,
		MinRefBalance: m.MinRefBalance,
		Fees:          toFeesDatabaseList(m.Fees),
	}
}

// toFeesDatabaseList - mapping proto list to db list
func toFeesDatabaseList(m []*feepolicy.Fees) Fees {
	res := make(Fees, 0, len(m))
	for _, fee := range m {
		res = append(res, toFeesDatabase(fee))
	}

	return res
}

// toFeesDatabase - mapping proto model to db model
func toFeesDatabase(m *feepolicy.Fees) Fee {
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
