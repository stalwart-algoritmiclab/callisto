package exchanger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	if len(tx.Logs) == 0 {
		return nil
	}

	switch exchangerMsg := msg.(type) {
	case *exchanger.MsgExchange:
		return m.handleMsgExchange(tx, index, exchangerMsg)
	default:
		return nil
	}
}
