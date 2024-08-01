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

// insertOptions inserts a new row into the options table.
func (r Repository) insertOptions(options []types.Options, pollID uint64) error {
	if options == nil {
		return nil
	}

	q := `
		INSERT INTO stwart_polls_options (
		    id, poll_id, voters_count, tokens_amount, is_veto, text, is_winner
		) VALUES (
		    $1, $2, $3, $4, $5, $6, $7
		) RETURNING
			id, poll_id, voters_count, tokens_amount, is_veto, text, is_winner
	`

	for _, option := range options {
		m := toDatabaseOptions(option)

		if _, err := r.db.Exec(q, m.ID, pollID, m.VotersCount, pq.Array(m.TokensAmount), m.IsVeto, m.Text, m.IsWinner); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// getAllOptions returns all Options from the database.
func (r Repository) getAllOptions(filter filter.Filter) ([]Options, error) {
	query, args := filter.Build(tablePollsOptions)

	var result []Options
	if err := r.db.Select(&result, query, args...); err != nil {
		return nil, errs.Internal{Cause: err.Error()}
	}

	return result, nil
}

// InsertMsgCreatePoll inserts a new row into the msg_create_poll table.
func (r Repository) InsertMsgCreatePoll(hash string, msgs ...*types.MsgCreatePoll) error {
	if msgs == nil || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_polls_msg_create_poll (
		    creator, title, description, voting_start_time, voting_period, min_vote_amount, min_adresses_count, min_vote_coins_amount, tx_hash
		) VALUES (	
		    $1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING
			id, creator, title, description, voting_start_time, voting_period, min_vote_amount, min_adresses_count, min_vote_coins_amount, tx_hash
	`

	for _, msg := range msgs {
		m := toDatabaseMsgCreatePolls(hash, msg)

		row := r.db.QueryRow(q, m.Creator, m.Title, m.Description, m.VotingStartTime, m.VotingPeriod, m.MinVoteAmount, m.MinAdressesCount, m.MinVoteCoinsAmount, m.TxHash)

		var result MsgCreatePolls

		if err := row.Scan(&result.ID, &result.Creator, &result.Title, &result.Description, &result.VotingStartTime, &result.VotingPeriod, &result.MinVoteAmount, &result.MinAdressesCount, &result.MinVoteCoinsAmount, &result.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}

		if err := r.insertOptions(msg.Options, result.ID); err != nil {
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
