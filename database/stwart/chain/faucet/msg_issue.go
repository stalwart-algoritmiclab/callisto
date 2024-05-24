package faucet

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

// GetAllMsgIssue - method that get data from a db (stwartchain_faucet).
func (r Repository) GetAllMsgIssue(filter filter.Filter) ([]faucet.MsgIssue, error) {
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
func (r Repository) InsertMsgIssue(hash string, msgs ...*faucet.MsgIssue) error {
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
