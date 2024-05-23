package exchanger

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/jmoiron/sqlx"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
)

var _ chain.Exchanger = &Repository{}

type (
	// Repository - defines a repository for allowed repository
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
