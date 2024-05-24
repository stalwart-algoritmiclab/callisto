package faucet

const (
	tableFaucet = "stwart_faucet"
)

// MsgIssue - db model for 'stwart_faucet'
type MsgIssue struct {
	ID      uint64 `db:"id"`
	Creator string `db:"creator"`
	Address string `db:"address"`
	TxHash  string `db:"tx_hash"`
}
