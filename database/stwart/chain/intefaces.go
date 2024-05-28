package chain

import (
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
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

	// Faucet - describes an interface for working with faucet database models.
	Faucet interface {
		GetAllMsgIssue(filter filter.Filter) ([]faucet.MsgIssue, error)
		InsertMsgIssue(hash string, msgs ...*faucet.MsgIssue) error
	}
)
