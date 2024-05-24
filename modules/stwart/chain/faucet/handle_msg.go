package faucet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch faucetMsg := msg.(type) {
	case *faucet.MsgIssue:
		return m.handleMsgIssue(tx, index, faucetMsg)
	default:
		return nil
	}
}
