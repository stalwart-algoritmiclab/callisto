package secured

import (
	"github.com/lib/pq"
)

const (
	tableCreateAddresses = "stwart_secured_create_addresses"
	tableDeleteAddresses = "stwart_secured_delete_addresses"
	tableUpdateAddresses = "stwart_secured_update_addresses"
)

// MsgCreateAddresses - db model for 'stwart_secured_create_addresses'
type (
	MsgCreateAddresses struct {
		ID        uint64         `db:"id"`
		Creator   string         `db:"creator"`
		TxHash    string         `db:"tx_hash"`
		Addresses pq.StringArray `db:"addresses"`
	}

	// MsgDeleteAddresses - db model for 'stwart_secured_delete_addresses'
	MsgDeleteAddresses struct {
		ID      uint64 `db:"id"`
		Creator string `db:"creator"`
		TxHash  string `db:"tx_hash"`
	}

	// MsgUpdateAddresses - db model for 'stwart_secured_update_addresses'
	MsgUpdateAddresses struct {
		ID        uint64         `db:"id"`
		Creator   string         `db:"creator"`
		TxHash    string         `db:"tx_hash"`
		Addresses pq.StringArray `db:"addresses"`
	}
)
