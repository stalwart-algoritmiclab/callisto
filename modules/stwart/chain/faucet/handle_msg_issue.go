package faucet

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

// handleMsgIssue allows to properly handle a MsgIssue
func (m *Module) handleMsgIssue(tx *juno.Tx, _ int, msg *faucet.MsgIssue) error {
	return m.faucetRepo.InsertMsgIssue(tx.TxHash, msg)
}
