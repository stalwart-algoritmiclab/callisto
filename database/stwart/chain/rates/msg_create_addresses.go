package rates

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
)

// GetAllMsgCreateAddresses - method that get data from a db (stwartchain_rates_create_addresses).
func (r Repository) GetAllMsgCreateAddresses(filter filter.Filter) ([]rates.MsgCreateAddresses, error) {
	query, args := filter.Build(tableCreateAdrresses)

	var result []MsgCreateAddresses
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableCreateAdrresses}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableCreateAdrresses}
	}

	return toMsgCreateAddressesDomainList(result), nil
}

// InsertMsgCreateAddresses - insert a new MsgCreateAddresses in a database (stwartchain_rates_create_addresses).
func (r Repository) InsertMsgCreateAddresses(hash string, msgs ...*rates.MsgCreateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_create_addresses (
			 creator, address, tx_hash
		) VALUES (
			$1, $2, $3
		) RETURNING
			id,  creator, address, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgCreateAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Address, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
