/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const (
	tableMsgCreateTariffs   = "stwart_feepolicy_msg_create_tariffs"
	tableMsgUpdateTariffs   = "stwart_feepolicy_msg_update_tariffs"
	tableMsgDeleteTariffs   = "stwart_feepolicy_msg_delete_tariffs"
	tableTariffs            = "stwart_feepolicy_tariffs"
	tableMsgCreateAddresses = "stwart_feepolicy_create_addresses"
	tableMsgUpdateAddresses = "stwart_feepolicy_update_addresses"
	tableMsgDeleteAddresses = "stwart_feepolicy_delete_addresses"
)

type (
	// MsgCreateAddresses - db model for 'stwart_feepolicy_create_addresses'
	MsgCreateAddresses struct {
		ID      uint64 `db:"id"`
		Creator string `db:"creator"`
		Address string `db:"address"`
		TxHash  string `db:"tx_hash"`
	}

	// MsgUpdateAddresses - db model for 'stwart_feepolicy_update_addresses'
	MsgUpdateAddresses struct {
		ID        uint64 `db:"id"`
		Creator   string `db:"creator"`
		Address   string `db:"address"`
		AddressID uint64 `db:"address_id"`
		TxHash    string `db:"tx_hash"`
	}

	// MsgDeleteAddresses - db model for 'stwart_feepolicy_delete_addresses'
	MsgDeleteAddresses struct {
		ID        uint64 `db:"id"`
		Creator   string `db:"creator"`
		AddressID uint64 `db:"address_id"`
		TxHash    string `db:"tx_hash"`
	}

	// MsgCreateTariffs - db model for 'stwart_feepolicy_msg_create_tariffs'
	MsgCreateTariffs struct {
		ID      int    `db:"id"`
		Denom   string `db:"denom"`
		Creator string `db:"creator"`
		Tariffs Tariff `db:"tariff_detail"`
	}

	// MsgUpdateTariffs - db model for 'stwart_feepolicy_msg_update_tariffs'
	MsgUpdateTariffs struct {
		ID      int    `db:"id"`
		Denom   string `db:"denom"`
		Creator string `db:"creator"`
		Tariffs Tariff `db:"tariff_detail"`
	}

	// MsgDeleteTariffs - db model for 'stwart_feepolicy_msg_delete_tariffs'
	MsgDeleteTariffs struct {
		ID       int    `db:"id"`
		Creator  string `db:"creator"`
		Denom    string `db:"denom"`
		TariffID string `db:"tariff_id"`
		FeeID    string `db:"fee_id"`
	}

	// Tariff - db model for 'stwart_feepolicy_tariffs'
	Tariff struct {
		ID            int    `db:"id"`
		TariffID      uint64 `db:"tariff_id"`
		Amount        string `db:"amount"`
		Denom         string `db:"denom"`
		MinRefBalance string `db:"min_ref_balance"`
		Fees          Fees   `db:"fees"` // Используем тип Fees для хранения JSON данных
	}

	// Fee json struct for Fee
	Fee struct {
		AmountFrom  string `json:"amountFrom"`
		Fee         string `json:"fee"`
		RefReward   string `json:"refReward"`
		StakeReward string `json:"stakeReward"`
		MinAmount   uint64 `json:"minAmount"`
		NoRefReward bool   `json:"noRefReward"`
		Creator     string `json:"creator"`
		ID          uint64 `json:"id"`
	}

	// Fees json struct for Fees
	Fees []Fee

	// Tariffs json struct for Tariffs
	Tariffs []Tariff
)

// Value realizes the driver.Valuer interface
func (f Fees) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan realizes the sql.Scanner interface
func (f *Fees) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not a byte slice")
	}

	return json.Unmarshal(bytes, f)
}
