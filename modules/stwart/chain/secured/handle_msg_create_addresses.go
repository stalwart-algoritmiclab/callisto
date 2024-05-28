package secured

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

// handleMsgIssue allows to properly handle a MsgIssue
func (m *Module) handleMsgCreateAddresses(tx *juno.Tx, _ int, msg *secured.MsgCreateAddresses) error {
	return m.securedRepo.InsertMsgCreateAddresses(tx.TxHash, msg)
}
