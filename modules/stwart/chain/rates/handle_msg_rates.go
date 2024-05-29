package rates

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// handleMsgCreateRates allows to properly handle a MsgCreateRates
func (m *Module) handleMsgCreateRates(tx *juno.Tx, msg *rates.MsgCreateRates) error {

	return m.ratesRepo.InsertMsgCreateRates(tx.TxHash, &rates.MsgCreateRates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	})
}

// handleMsgUpdateRates allows to properly handle a MsgUpdateRates
func (m *Module) handleMsgUpdateRates(tx *juno.Tx, msg *rates.MsgUpdateRates) error {
	return m.ratesRepo.InsertMsgUpdateRates(tx.TxHash, &rates.MsgUpdateRates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	})
}

func (m *Module) handleMsgDeleteRates(tx *juno.Tx, msg *rates.MsgDeleteRates) error {
	return m.ratesRepo.InsertMsgDeleteRates(tx.TxHash, &rates.MsgDeleteRates{
		Creator: msg.Creator,
		Denom:   msg.Denom,
	})
}
