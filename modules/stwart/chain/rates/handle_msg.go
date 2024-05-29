package rates

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	switch ratesMsg := msg.(type) {
	case *rates.MsgCreateRates:
		return m.handleMsgCreateRates(tx, ratesMsg)
	case *rates.MsgCreateAddresses:
		return m.handleMsgCreateAddresses(tx, ratesMsg)
	case *rates.MsgUpdateRates:
		return m.handleMsgUpdateRates(tx, ratesMsg)
	case *rates.MsgDeleteRates:
		return m.handleMsgDeleteRates(tx, ratesMsg)
	default:
		return nil
	}
}
