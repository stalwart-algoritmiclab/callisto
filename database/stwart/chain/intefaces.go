/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package chain

import (
	coretypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	exchangertypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
	faucettypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
	feepolicytypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
	ratestypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
	referraltypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"
	securedtypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"

	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

type (
	// LastBlock - describes an interface for working with last block database models.
	LastBlock interface {
		Get() (uint64, error)
		Update(id uint64) error
	}

	// Exchanger - describes an interface for working with exchanger database models.
	Exchanger interface {
		GetAllMsgExchange(filter filter.Filter) ([]exchangertypes.MsgExchange, error)
		InsertMsgExchange(hash string, msgs ...*exchangertypes.MsgExchange) error
	}

	// Faucet - describes an interface for working with faucet database models.
	Faucet interface {
		GetAllMsgIssue(filter filter.Filter) ([]faucettypes.MsgIssue, error)
		InsertMsgIssue(hash string, msgs ...*faucettypes.MsgIssue) error
	}

	// Feepolicy - describes an interface for working with feepolicy database models.
	Feepolicy interface {
		InsertMsgCreateTariffs(height uint64, hash string, msgs ...*feepolicytypes.MsgCreateTariffs) error
		InsertMsgUpdateTariffs(height uint64, hash string, msgs ...*feepolicytypes.MsgUpdateTariffs) error
		InsertDeleteMsgDeleteTariffs(height uint64, hash string, msgs ...*feepolicytypes.MsgDeleteTariffs) error
		InsertMsgCreateAddresses(hash string, msgs ...*feepolicytypes.MsgCreateAddresses) error
		InsertMsgUpdateAddresses(hash string, msgs ...*feepolicytypes.MsgUpdateAddresses) error
		InsertMsgDeleteAddresses(hash string, msgs ...*feepolicytypes.MsgDeleteAddresses) error
	}

	// Core - describes an interface for working with database models.
	Core interface {
		GetAllMsgIssue(filter filter.Filter) ([]coretypes.MsgIssue, error)
		InsertMsgIssue(hash string, msgs ...*coretypes.MsgIssue) error

		GetAllMsgWithdraw(filter filter.Filter) ([]coretypes.MsgWithdraw, error)
		InsertMsgWithdraw(hash string, msgs ...*coretypes.MsgWithdraw) error

		GetAllMsgRefund(filter filter.Filter) ([]coretypes.MsgRefund, error)
		InsertMsgRefund(hash string, msgs ...*coretypes.MsgRefund) error

		GetAllMsgFees(filter filter.Filter) ([]coretypes.MsgFees, error)
		InsertMsgFees(hash string, msgs ...*coretypes.MsgFees) error

		GetAllMsgRefReward(filter filter.Filter) ([]coretypes.MsgRefReward, error)
		InsertMsgRefReward(hash string, msgs ...*coretypes.MsgRefReward) error

		GetAllMsgSend(filter filter.Filter) ([]coretypes.MsgSend, error)
		InsertMsgSend(hash string, msgs ...*coretypes.MsgSend) error
	}

	// Rates - describes an interface for working with rates database models.
	Rates interface {
		GetAllMsgCreateAddresses(filter filter.Filter) ([]ratestypes.MsgCreateAddresses, error)
		InsertMsgCreateAddresses(hash string, msgs ...*ratestypes.MsgCreateAddresses) error

		GetAllMsgUpdateAddresses(filter filter.Filter) ([]ratestypes.MsgUpdateAddresses, error)
		InsertMsgUpdateAddresses(hash string, msgs ...*ratestypes.MsgUpdateAddresses) error

		GetAllMsgDeleteAddresses(filter filter.Filter) ([]ratestypes.MsgDeleteAddresses, error)
		InsertMsgDeleteAddresses(hash string, msgs ...*ratestypes.MsgDeleteAddresses) error

		GetAllMsgCreateRates(filter filter.Filter) ([]ratestypes.MsgCreateRates, error)
		InsertMsgCreateRates(hash string, msgs ...*ratestypes.MsgCreateRates) error

		GetAllMsgUpdateRates(filter filter.Filter) ([]ratestypes.MsgUpdateRates, error)
		InsertMsgUpdateRates(hash string, msgs ...*ratestypes.MsgUpdateRates) error

		GetAllMsgDeleteRates(filter filter.Filter) ([]ratestypes.MsgDeleteRates, error)
		InsertMsgDeleteRates(hash string, msgs ...*ratestypes.MsgDeleteRates) error
	}

	// Secured - describes an interface for working with secured database models.
	Secured interface {
		GetAllMsgCreateAddresses(filter filter.Filter) ([]securedtypes.MsgCreateAddresses, error)
		InsertMsgCreateAddresses(hash string, msgs ...*securedtypes.MsgCreateAddresses) error

		GetAllMsgDeleteAddresses(filter filter.Filter) ([]securedtypes.MsgDeleteAddresses, error)
		InsertMsgDeleteAddresses(hash string, msgs ...*securedtypes.MsgDeleteAddresses) error

		GetAllMsgUpdateAddresses(filter filter.Filter) ([]securedtypes.MsgUpdateAddresses, error)
		InsertMsgUpdateAddresses(hash string, msgs ...*securedtypes.MsgUpdateAddresses) error
	}

	// Referrals - describes an interface for working with referrals database models.
	Referrals interface {
		InsertMsgCreateUser(hash string, msgs ...*referraltypes.MsgCreateUser) error
		InsertMsgUpdateUser(hash string, msgs ...*referraltypes.MsgUpdateUser) error
		InsertMsgDeleteUser(hash string, msgs ...*referraltypes.MsgDeleteUser) error
		InsertMsgSetReferrer(hash string, msgs ...*referraltypes.MsgSetReferrer) error
	}
)
