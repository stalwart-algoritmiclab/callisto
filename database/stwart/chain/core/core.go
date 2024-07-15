/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	"database/sql"
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/jmoiron/sqlx"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

var _ chain.Core = &Repository{}

type (
	// Repository - defines a repository for allowed repository
	Repository struct {
		cdc codec.Codec
		db  *sqlx.DB
	}
)

// NewRepository constructor.
func NewRepository(db *sqlx.DB, cdc codec.Codec) *Repository {
	return &Repository{
		cdc: cdc,
		db:  db,
	}
}

func (r Repository) GetAllMsgIssue(filter filter.Filter) ([]types.MsgIssue, error) {
	query, args := filter.Build(tableIssue)

	var result []MsgIssue
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableIssue}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableIssue}
	}

	return toMsgIssueDomainList(result), nil

}

func (r Repository) InsertMsgIssue(hash string, msgs ...*types.MsgIssue) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_core_issue (
		 creator, denom, amount, address, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING
			id, tx_hash, creator, denom, amount, address
	`

	for _, msg := range msgs {
		m, err := toMsgIssueDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.Creator, m.Denom, m.Amount, m.Address, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgWithdraw(filter filter.Filter) ([]types.MsgWithdraw, error) {
	query, args := filter.Build(tableWithdraw)

	var result []MsgWithdraw
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableWithdraw}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableWithdraw}
	}

	return toMsgWithdrawDomainList(result), nil
}

func (r Repository) InsertMsgWithdraw(hash string, msgs ...*types.MsgWithdraw) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_core_withdraw (
		 creator, denom, amount, address, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING
			id, tx_hash, creator, denom, amount, address
	`

	for _, msg := range msgs {
		m, err := toMsgWithdrawDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.Creator, m.Denom, m.Amount, m.Address, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgRefund(filter filter.Filter) ([]types.MsgRefund, error) {
	query, args := filter.Build(tableRefund)

	var result []MsgRefund
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableRefund}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableRefund}
	}

	return toMsgRefundDomainList(result), nil
}

func (r Repository) InsertMsgRefund(hash string, msgs ...*types.MsgRefund) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_core_refund (
		 creator, amount, from_address, to_address, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING
			id, tx_hash, creator, amount, from_address, to_address
	`

	for _, msg := range msgs {
		m, err := toMsgRefundDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.Creator, m.Amount, m.From, m.To, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgFees(filter filter.Filter) ([]types.MsgFees, error) {
	query, args := filter.Build(tableFees)

	var result []MsgFees
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableFees}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableFees}
	}

	return toMsgFeesDomainList(result), nil
}

func (r Repository) InsertMsgFees(hash string, msgs ...*types.MsgFees) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_core_fees (
		 creator, commission, address, tx_hash
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, tx_hash, creator, commission, address
	`

	for _, msg := range msgs {
		m, err := toMsgFeesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.Creator, m.Commission, m.Address, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgRefReward(filter filter.Filter) ([]types.MsgRefReward, error) {
	query, args := filter.Build(tableRefReward)

	var result []MsgRefReward
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableRefReward}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableRefReward}
	}

	return toMsgRefRewardDomainList(result), nil
}

func (r Repository) InsertMsgRefReward(hash string, msgs ...*types.MsgRefReward) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_core_refreward (
		 creator, amount, referrer, tx_hash
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, tx_hash, creator, amount, referrer
	`

	for _, msg := range msgs {
		m, err := toMsgRefRewardDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.Creator, m.Amount, m.Referrer, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgSend(filter filter.Filter) ([]types.MsgSend, error) {
	query, args := filter.Build(tableSend)

	var result []MsgSend
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableSend}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableSend}
	}

	return toMsgSendDomainList(result), nil
}

func (r Repository) InsertMsgSend(hash string, msgs ...*types.MsgSend) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_core_send (
		 creator, from_address, to_address, amount, denom, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING
			id, tx_hash, creator, from_address, to_address, amount, denom
	`

	for _, msg := range msgs {
		m, err := toMsgSendDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.Creator, m.From, m.To, m.Amount, m.Denom, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
