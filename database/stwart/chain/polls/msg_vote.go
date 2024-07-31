/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import (
	"github.com/lib/pq"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// InsertMsgVote inserts a new row into the msg_vote table.
func (r Repository) InsertMsgVote(hash string, msgs ...*types.MsgVote) error {
	if msgs == nil || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_polls_msg_vote (
		    creator, poll_id, option_id, amount, tx_hash
		) VALUES (	
		    $1, $2, $3, $4, $5
		) RETURNING
			id, creator, poll_id, option_id, amount, tx_hash
	`

	for _, msg := range msgs {
		m := toDatabaseMsgVote(hash, msg)

		if _, err := r.db.Exec(q, m.Creator, m.PollID, m.OptionID, pq.Array(m.Amount), m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// GetAllMsgVote returns all MsgVote from the database.
func (r Repository) GetAllMsgVote(filter filter.Filter) ([]types.MsgVote, error) {
	query, args := filter.Build(tableMsgVote)

	var result []MsgVote
	if err := r.db.Select(&result, query, args...); err != nil {
		return nil, errs.Internal{Cause: err.Error()}
	}

	return toMsgVoteDomainList(result), nil
}
