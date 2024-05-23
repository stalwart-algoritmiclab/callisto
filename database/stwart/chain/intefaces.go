package chain

import (
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

// LastBlock - describes an interface for working with database models.
type (
	LastBlock interface {
		Get() (uint64, error)
		Update(id uint64) error
	}

	// Exchanger - describes an interface for working with database models.
	Exchanger interface {
		GetAllMsgExchange(filter filter.Filter) ([]exchanger.MsgExchange, error)
		InsertMsgExchange(hash string, msgs ...*exchanger.MsgExchange) error
	}
)
