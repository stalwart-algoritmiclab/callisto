package remote

import (
	"github.com/forbole/juno/v5/node/remote"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet/source"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

var (
	_ source.Source = &Source{}
)

// Source represents the implementation of the QueryClient that works on a remote node
type Source struct {
	*remote.Source
	client faucet.QueryClient
}

// NewSource builds a new Source instance
func NewSource(source *remote.Source, faucetClient faucet.QueryClient) *Source {
	return &Source{
		Source: source,
		client: faucetClient,
	}
}
