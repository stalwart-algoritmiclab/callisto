package rates

import "github.com/lib/pq"

const (
	tableCreateAdrresses = "stwart_rates_create_addresses"
	tableCreateRates     = "stwart_rates_create_rates"
	tableUpdateRates     = "stwart_rates_update_rates"
	tableDeleteRates     = "stwart_rates_delete_rates"
)

// MsgCreateAddresses - db model for 'stwart_rates_create_addresses'
type MsgCreateAddresses struct {
	ID      uint64         `db:"id"`
	Creator string         `db:"creator"`
	Address pq.StringArray `db:"address"`
	TxHash  string         `db:"tx_hash"`
}

// MsgCreateRates - db model for 'stwart_rates_create_rates'
type MsgCreateRates struct {
	ID       uint64 `db:"id"`
	Creator  string `db:"creator"`
	Decimals int32  `db:"decimals"`
	Denom    string `db:"denom"`
	Rate     string `db:"rate"`
	TxHash   string `db:"tx_hash"`
}

// MsgUpdateRates - db model for 'stwart_rates_create_rates'
type MsgUpdateRates struct {
	ID      uint64 `db:"id"`
	Creator string `db:"creator"`
	Denom   string `db:"denom"`
	Rate    string `db:"rate"`
	TxHash  string `db:"tx_hash"`
}

// MsgDeleteRates - db model for 'stwart_rates_create_rates'
type MsgDeleteRates struct {
	ID      uint64 `db:"id"`
	Creator string `db:"creator"`
	Denom   string `db:"denom"`
	TxHash  string `db:"tx_hash"`
}
