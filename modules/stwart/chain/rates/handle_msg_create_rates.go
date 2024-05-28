package rates

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// handleMsgCreateRates allows to properly handle a MsgCreateRates
func (m *Module) handleMsgCreateRates(tx *juno.Tx, _ int, msg *rates.MsgCreateRates) error {

	return m.ratesRepo.InsertMsgCreateRates(tx.TxHash, &rates.MsgCreateRates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	})
}
