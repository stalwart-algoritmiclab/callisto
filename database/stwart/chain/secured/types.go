package secured

import (
	"github.com/lib/pq"
)

const (
	tableSecured = "stwart_secured_addresses"
)

// MsgCreateAddresses - db model for 'stwart_secured_addresses'
type MsgCreateAddresses struct {
	ID        uint64         `db:"id"`
	Creator   string         `db:"creator"`
	TxHash    string         `db:"tx_hash"`
	Addresses pq.StringArray `db:"addresses"`
}
