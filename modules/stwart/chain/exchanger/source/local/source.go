package local

import (
	"github.com/forbole/juno/v5/node/local"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryServer that works on a local node
type Source struct {
	*local.Source
	exchangerServer exchanger.QueryServer
}

// NewSource builds a new Source instance
func NewSource(source *local.Source, exchangerServer exchanger.QueryServer) *Source {
	return &Source{
		Source:          source,
		exchangerServer: exchangerServer,
	}
}