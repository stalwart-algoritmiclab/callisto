package core

const (
	tableIssue     = "stwart_core_issue"
	tableWithdraw  = "stwart_core_withdraw"
	tableRefReward = "stwart_core_refreward"
	tableRefund    = "stwart_core_refund"
	tableFees      = "stwart_core_fees"
	tableSend      = "stwart_core_send"
)

type (
	// MsgIssue - db model for 'stwart_core_issue'
	MsgIssue struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Denom   string `db:"denom"`
		Amount  string `db:"amount"`
		Address string `db:"address"`
	}

	// MsgFees - db model for 'stwart_core_fees'
	MsgFees struct {
		ID         uint64 `db:"id"`
		TxHash     string `db:"tx_hash"`
		Creator    string `db:"creator"`
		Commission string `db:"commission"`
		Address    string `db:"address"`
	}

	// MsgRefReward - db model for 'stwart_core_ref_reward'
	MsgRefReward struct {
		ID       uint64 `db:"id"`
		TxHash   string `db:"tx_hash"`
		Creator  string `db:"creator"`
		Amount   string `db:"amount"`
		Referrer string `db:"referrer"`
	}

	// MsgRefund - db model for 'stwart_core_refund'
	MsgRefund struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  string `db:"amount"`
		From    string `db:"from_address"`
		To      string `db:"to_address"`
	}

	// MsgSend - db model for 'stwart_core_send'
	MsgSend struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		From    string `db:"from_address"`
		To      string `db:"to_address"`
		Amount  string `db:"amount"`
		Denom   string `db:"denom"`
	}

	// MsgWithdraw - db model for 'stwart_core_withdraw'
	MsgWithdraw struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  string `db:"amount"`
		Denom   string `db:"denom"`
		Address string `db:"address"`
	}
)
