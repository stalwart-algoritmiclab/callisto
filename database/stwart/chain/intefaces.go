/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package chain

import (
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/core"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/referrals"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

type (
	// LastBlock - describes an interface for working with last block database models.
	LastBlock interface {
		Get() (uint64, error)
		Update(id uint64) error
	}

	// Exchanger - describes an interface for working with exchanger database models.
	Exchanger interface {
		GetAllMsgExchange(filter filter.Filter) ([]exchanger.MsgExchange, error)
		InsertMsgExchange(hash string, msgs ...*exchanger.MsgExchange) error
	}

	// Faucet - describes an interface for working with faucet database models.
	Faucet interface {
		GetAllMsgIssue(filter filter.Filter) ([]faucet.MsgIssue, error)
		InsertMsgIssue(hash string, msgs ...*faucet.MsgIssue) error
	}

	// Feepolicy - describes an interface for working with feepolicy database models.
	Feepolicy interface {
		InsertMsgCreateTariffs(height int64, hash string, msgs ...*feepolicy.MsgCreateTariffs) error
		InsertMsgUpdateTariffs(height int64, hash string, msgs ...*feepolicy.MsgUpdateTariffs) error
		InsertDeleteMsgDeleteTariffs(height int64, hash string, msgs ...*feepolicy.MsgDeleteTariffs) error
		InsertMsgCreateAddresses(hash string, msgs ...*feepolicy.MsgCreateAddresses) error
		InsertMsgUpdateAddresses(hash string, msgs ...*feepolicy.MsgUpdateAddresses) error
		InsertMsgDeleteAddresses(hash string, msgs ...*feepolicy.MsgDeleteAddresses) error
	}

	// Core - describes an interface for working with database models.
	Core interface {
		GetAllMsgIssue(filter filter.Filter) ([]core.MsgIssue, error)
		InsertMsgIssue(hash string, msgs ...*core.MsgIssue) error

		GetAllMsgWithdraw(filter filter.Filter) ([]core.MsgWithdraw, error)
		InsertMsgWithdraw(hash string, msgs ...*core.MsgWithdraw) error

		GetAllMsgRefund(filter filter.Filter) ([]core.MsgRefund, error)
		InsertMsgRefund(hash string, msgs ...*core.MsgRefund) error

		GetAllMsgFees(filter filter.Filter) ([]core.MsgFees, error)
		InsertMsgFees(hash string, msgs ...*core.MsgFees) error

		GetAllMsgRefReward(filter filter.Filter) ([]core.MsgRefReward, error)
		InsertMsgRefReward(hash string, msgs ...*core.MsgRefReward) error

		GetAllMsgSend(filter filter.Filter) ([]core.MsgSend, error)
		InsertMsgSend(hash string, msgs ...*core.MsgSend) error
	}

	// Rates - describes an interface for working with rates database models.
	Rates interface {
		GetAllMsgCreateAddresses(filter filter.Filter) ([]rates.MsgCreateAddresses, error)
		InsertMsgCreateAddresses(hash string, msgs ...*rates.MsgCreateAddresses) error
		GetAllMsgCreateRates(filter filter.Filter) ([]rates.MsgCreateRates, error)
		InsertMsgCreateRates(hash string, msgs ...*rates.MsgCreateRates) error
	}

	// Secured - describes an interface for working with secured database models.
	Secured interface {
		GetAllMsgCreateAddresses(filter filter.Filter) ([]secured.MsgCreateAddresses, error)
		InsertMsgCreateAddresses(hash string, msgs ...*secured.MsgCreateAddresses) error

		GetAllMsgDeleteAddresses(filter filter.Filter) ([]secured.MsgDeleteAddresses, error)
		InsertMsgDeleteAddresses(hash string, msgs ...*secured.MsgDeleteAddresses) error

		GetAllMsgUpdateAddresses(filter filter.Filter) ([]secured.MsgUpdateAddresses, error)
		InsertMsgUpdateAddresses(hash string, msgs ...*secured.MsgUpdateAddresses) error
	}

	// Referrals - describes an interface for working with referrals database models.
	Referrals interface {
		InsertMsgCreateUser(hash string, msgs ...*referrals.MsgCreateUser) error
		InsertMsgUpdateUser(hash string, msgs ...*referrals.MsgUpdateUser) error
		InsertMsgDeleteUser(hash string, msgs ...*referrals.MsgDeleteUser) error
		InsertMsgSetReferrer(hash string, msgs ...*referrals.MsgSetReferrer) error
	}
)
