package rates

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateAddresses{}

func (msg *MsgCreateAddresses) ValidateBasic() error {
	return nil
}

func (msg *MsgCreateAddresses) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(msg.Creator)}
}
