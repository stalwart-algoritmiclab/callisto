/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package faucet

import (
	"database/sql"
	"errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"

	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
)

// GetAllMsgIssue - method that get data from a db (stwartchain_faucet).
func (r Repository) GetAllMsgIssue(filter filter.Filter) ([]types.MsgIssue, error) {
	query, args := filter.Build(tableFaucet)

	var result []MsgIssue
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableFaucet}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableFaucet}
	}

	return toMsgIssueDomainList(result), nil
}

// InsertMsgIssue - insert a new MsgIssue in a database (stwartchain_faucet).
func (r Repository) InsertMsgIssue(hash string, msgs ...*types.MsgIssue) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_faucet (
			 creator, address, tx_hash
		) VALUES (
			$1, $2, $3
		) RETURNING
			id,  creator, address, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgIssueDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.TxHash, m.Creator, m.Address); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
