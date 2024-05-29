package rates

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress { return m.GetSigners() }

func (msg *MsgCreateAddresses) GetSigners() []sdk.AccAddress { return msg.GetSigners() }

func (msg *MsgUpdateAddresses) GetSigners() []sdk.AccAddress { return msg.GetSigners() }

func (msg *MsgDeleteAddresses) GetSigners() []sdk.AccAddress { return msg.GetSigners() }

func (msg *MsgCreateRates) GetSigners() []sdk.AccAddress { return msg.GetSigners() }

func (msg *MsgUpdateRates) GetSigners() []sdk.AccAddress { return msg.GetSigners() }

func (msg *MsgDeleteRates) GetSigners() []sdk.AccAddress { return msg.GetSigners() }
