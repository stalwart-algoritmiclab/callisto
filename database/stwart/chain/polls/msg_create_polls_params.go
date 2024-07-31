/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// InsertMsgCreatePollsParams - insert msg create polls params
func (r Repository) InsertMsgCreatePollsParams(hash string, msgs ...*types.MsgCreatePollsParams) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_polls_msg_create_polls_params (
		    creator, min_days_duration, max_days_duration, max_days_pending, proposer_deposit, burn_veto, tx_hash
		) VALUES (
		    $1, $2, $3, $4, $5, $6, $7
		) RETURNING
		  	id, creator, min_days_duration, max_days_duration, max_days_pending, proposer_deposit, burn_veto, tx_hash
	`

	for _, msg := range msgs {
		m := toDatabaseMsgCreatePollsParams(hash, msg)

		if _, err := r.db.Exec(q, m.Creator, m.MinDaysDuration, m.MaxDaysDuration, m.MaxDaysPending, m.ProposerDeposit, m.BurnVeto, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// GetAllMsgCreatePollsParams - get all msg create polls params
func (r Repository) GetAllMsgCreatePollsParams(filter filter.Filter) ([]types.MsgCreatePollsParams, error) {
	query, args := filter.Build(tableMsgCreatePollsParams)

	var result []MsgCreatePollsParams
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableMsgCreatePollsParams}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableMsgCreatePollsParams}
	}

	return toMsgCreatePollsParamsDomainList(result), nil
}
