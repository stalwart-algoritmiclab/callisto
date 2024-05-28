package rates

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateRates{}

func (msg *MsgCreateRates) ValidateBasic() error {
	return nil
}

func (msg *MsgCreateRates) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}
