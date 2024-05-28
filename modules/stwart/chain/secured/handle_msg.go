package secured

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	//if len(tx.Logs) == 0 {
	//	return nil
	//}

	switch securedMsg := msg.(type) {
	case *secured.MsgCreateAddresses:
		return m.handleMsgCreateAddresses(tx, index, securedMsg)
	default:
		return nil
	}
}
