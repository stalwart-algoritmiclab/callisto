/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// InsertMsgUpdatePollsParams inserts a new MsgUpdatePollsParams into the database.
func (r Repository) InsertMsgUpdatePollsParams(hash string, msgs ...*types.MsgUpdatePollsParams) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_polls_msg_update_polls_params (
		    creator, min_days_duration, max_days_duration, max_days_pending, proposer_deposit, burn_veto, tx_hash
		) VALUES (	
		    $1, $2, $3, $4, $5, $6, $7
		) RETURNING
			id, creator, min_days_duration, max_days_duration, max_days_pending, proposer_deposit, burn_veto, tx_hash
	`

	for _, msg := range msgs {
		m := toDatabaseMsgUpdatePollsParams(hash, msg)

		if _, err := r.db.Exec(q, m.Creator, m.MinDaysDuration, m.MaxDaysDuration, m.MaxDaysPending, pq.Array(m.ProposerDeposit), m.BurnVeto, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// GetAllMsgUpdatePollsParams returns all MsgUpdatePollsParams from the database.
func (r Repository) GetAllMsgUpdatePollsParams(filter filter.Filter) ([]types.MsgUpdatePollsParams, error) {
	query, args := filter.Build(tableMsgUpdatePollsParams)

	var result []MsgUpdatePollsParams
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableMsgUpdatePollsParams}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}

	return toMsgUpdatePollsParamsDomainList(result), nil
}
