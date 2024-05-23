package exchanger

const (
	tableExchange = "stwart_exchanger"
)

// MsgExchange - db model for 'stwart_exchanger'
type MsgExchange struct {
	ID      uint64 `db:"id"`
	Creator string `db:"creator"`
	Denom   string `db:"denom"`
	Amount  string `db:"amount"`
	DenomTo string `db:"denom_to"`
	TxHash  string `db:"tx_hash"`
}
