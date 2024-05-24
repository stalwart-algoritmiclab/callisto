package faucet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// implement missing methods for sdk.Msg
var _ sdk.Msg = &MsgIssue{}

func (msg *MsgIssue) ValidateBasic() error {
	return nil
}

func (msg *MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
