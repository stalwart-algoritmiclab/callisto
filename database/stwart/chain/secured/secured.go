package secured

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/jmoiron/sqlx"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
)

var _ chain.Secured = &Repository{}

type (
	// Repository - defines a repository for secured repository
	Repository struct {
		cdc codec.Codec
		db  *sqlx.DB
	}
)

// NewRepository constructor.
func NewRepository(db *sqlx.DB, cdc codec.Codec) *Repository {
	return &Repository{
		cdc: cdc,
		db:  db,
	}
}
