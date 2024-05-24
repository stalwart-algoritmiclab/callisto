package exchanger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// implement missing methods for sdk.Msg
var _ sdk.Msg = &MsgExchange{}

func (msg *MsgExchange) ValidateBasic() error {
	return nil
}

func (msg *MsgExchange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
