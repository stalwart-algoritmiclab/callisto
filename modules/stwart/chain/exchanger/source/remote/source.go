package remote

import (
	"github.com/forbole/juno/v5/node/remote"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryClient that works on a remote node
type Source struct {
	*remote.Source
	client exchanger.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, exchangerClient exchanger.QueryClient) *Source {
	return &Source{
		Source: source,
		client: exchangerClient,
	}
}
