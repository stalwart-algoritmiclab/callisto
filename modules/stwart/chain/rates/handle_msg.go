package rates

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch ratesMsg := msg.(type) {
	case *rates.MsgCreateRates:
		return m.handleMsgCreateRates(tx, index, ratesMsg)
	case *rates.MsgCreateAddresses:
		return m.handleMsgCreateAddresses(tx, index, ratesMsg)
	default:
		return nil
	}
}
