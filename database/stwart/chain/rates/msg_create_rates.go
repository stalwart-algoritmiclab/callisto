package rates

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
)

// GetAllMsgCreateRates - method that get data from a db (stwartchain_rates_create_rates).
func (r Repository) GetAllMsgCreateRates(filter filter.Filter) ([]rates.MsgCreateRates, error) {
	query, args := filter.Build(tableCreateRates)

	var result []MsgCreateRates
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableCreateRates}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableCreateRates}
	}

	return toMsgCreateRatesDomainList(result), nil
}

// InsertMsgCreateRates - insert a new MsgCreateRates in a database (stwartchain_rates_create_rates).
func (r Repository) InsertMsgCreateRates(hash string, msgs ...*rates.MsgCreateRates) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_create_rates (
			creator, decimals, denom, rate, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING
			id, creator, decimals, denom, rate, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgCreateRatesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Decimals, m.Denom, m.Rate, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
