package exchanger

import sdk "github.com/cosmos/cosmos-sdk/types"

func (msg *MsgExchange) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}

func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
