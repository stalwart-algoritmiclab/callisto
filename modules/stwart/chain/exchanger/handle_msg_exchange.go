package exchanger

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

// handleMsgExchange handles the MsgExchange message
func (m *Module) handleMsgExchange(tx *juno.Tx, _ int, msg *exchanger.MsgExchange) error {
	return m.exchangerRepo.InsertMsgExchange(tx.TxHash, msg)
}
