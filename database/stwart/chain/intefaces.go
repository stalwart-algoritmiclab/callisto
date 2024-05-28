package chain

import (
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

type (
	// LastBlock - describes an interface for working with last block database models.
	LastBlock interface {
		Get() (uint64, error)
		Update(id uint64) error
	}

	// Exchanger - describes an interface for working with exchanger database models.
	Exchanger interface {
		GetAllMsgExchange(filter filter.Filter) ([]exchanger.MsgExchange, error)
		InsertMsgExchange(hash string, msgs ...*exchanger.MsgExchange) error
	}

	// Faucet - describes an interface for working with faucet database models.
	Faucet interface {
		GetAllMsgIssue(filter filter.Filter) ([]faucet.MsgIssue, error)
		InsertMsgIssue(hash string, msgs ...*faucet.MsgIssue) error
	}

	// Secured - describes an interface for working with secured database models.
	Secured interface {
		GetAllMsgCreateAddresses(filter filter.Filter) ([]secured.MsgCreateAddresses, error)
		InsertMsgCreateAddresses(hash string, msgs ...*secured.MsgCreateAddresses) error
	}
)
