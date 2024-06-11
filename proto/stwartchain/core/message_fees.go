/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

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
