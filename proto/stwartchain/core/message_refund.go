package core

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRefund{}

func NewMsgRefund(creator string, from string, to string, amount string) *MsgRefund {
	return &MsgRefund{
		Creator: creator,
		From:    from,
		To:      to,
		Amount:  amount,
	}
}

func (msg *MsgRefund) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
