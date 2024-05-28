package secured

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
)

// GetAllMsgCreateAddresses - method that get data from a db (stwart_secured_addresses).
func (r Repository) GetAllMsgCreateAddresses(filter filter.Filter) ([]secured.MsgCreateAddresses, error) {
	query, args := filter.Build(tableSecured)

	var result []MsgCreateAddresses
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableSecured}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableSecured}
	}

	return toMsgCreateAddressesDomainList(result), nil
}

// InsertMsgCreateAddresses - insert a new MsgCreateAddresses in a database (stwart_secured_addresses).
func (r Repository) InsertMsgCreateAddresses(hash string, msgs ...*secured.MsgCreateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_secured_addresses (
			 tx_hash, creator, addresses
		) VALUES (
			$1, $2, $3
		) RETURNING
			id, tx_hash, creator, addresses
	`

	for _, msg := range msgs {
		m, err := toMsgCreateAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.TxHash, m.Creator, m.Addresses); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
