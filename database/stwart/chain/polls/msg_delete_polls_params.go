/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// InsertMsgDeletePollsParams inserts a new row into the msg_delete_polls_params table.
func (r Repository) InsertMsgDeletePollsParams(hash string, msgs ...*types.MsgDeletePollsParams) error {
	if msgs == nil || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_polls_msg_delete_polls_params (
		    creator, tx_hash
		) VALUES (	
		    $1, $2
		) RETURNING
			id, creator, tx_hash
	`

	for _, msg := range msgs {
		m := toDatabaseMsgDeletePollsParams(hash, msg)

		if _, err := r.db.Exec(q, m.Creator, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// GetAllMsgDeletePollsParams returns all MsgDeletePollsParams from the database.
func (r Repository) GetAllMsgDeletePollsParams(filter filter.Filter) ([]types.MsgDeletePollsParams, error) {
	query, args := filter.Build(tableMsgDeletePollsParams)

	var result []MsgDeletePollsParams
	if err := r.db.Select(&result, query, args...); err != nil {
		return nil, errs.Internal{Cause: err.Error()}
	}

	return toMsgDeletePollsParamsDomainList(result), nil
}
