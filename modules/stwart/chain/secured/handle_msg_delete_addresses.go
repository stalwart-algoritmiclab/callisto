package secured

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

// handleMsgIssue allows to properly handle a MsgIssue
func (m *Module) handleMsgDeleteAddresses(tx *juno.Tx, _ int, msg *secured.MsgDeleteAddresses) error {
	return m.securedRepo.InsertMsgDeleteAddresses(tx.TxHash, msg)
}
