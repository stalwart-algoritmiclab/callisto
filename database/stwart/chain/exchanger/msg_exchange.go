package exchanger

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

// GetAllMsgExchange - method that get data from a db (stwartchain_exchanger).
func (r Repository) GetAllMsgExchange(filter filter.Filter) ([]exchanger.MsgExchange, error) {
	query, args := filter.Build(tableExchange)

	var result []MsgExchange
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableExchange}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableExchange}
	}

	return toMsgExchangeDomainList(result), nil
}

// InsertMsgExchange - insert a new MsgExchange in a database (stwartchain_exchanger).
func (r Repository) InsertMsgExchange(hash string, msgs ...*exchanger.MsgExchange) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_exchanger (
			 creator, denom, amount, denom_to, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING
			id,  creator, denom, amount, denom_to, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgExchangeDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.TxHash, m.Creator, m.Amount); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
