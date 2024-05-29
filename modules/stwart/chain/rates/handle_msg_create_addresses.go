package rates

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// handleMsgCreateAddresses allows to properly handle a MsgCreateAddresses
func (m *Module) handleMsgCreateAddresses(tx *juno.Tx, msg *rates.MsgCreateAddresses) error {
	return m.ratesRepo.InsertMsgCreateAddresses(tx.TxHash, &rates.MsgCreateAddresses{
		Creator: msg.Creator,
		Address: msg.Address,
	})
}
