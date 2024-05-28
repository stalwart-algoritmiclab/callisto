package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgFees{}

func NewMsgFees(comission sdk.Coin, addressTo string) *MsgFees { //nolint:misspell
	return &MsgFees{
		Comission: comission, //nolint:misspell
		AddressTo: addressTo,
	}
}

func (msg *MsgFees) ValidateBasic() error {
	return nil
}
