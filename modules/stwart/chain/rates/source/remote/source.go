package remote

import (
	"github.com/forbole/juno/v5/node/remote"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/rates/source"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryClient that works on a remote node
type Source struct {
	*remote.Source
	client rates.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, ratesClient rates.QueryClient) *Source {
	return &Source{
		Source: source,
		client: ratesClient,
	}
}
