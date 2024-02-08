package light

import (
	"context"

	hellarcore "github.com/hellarcore/tenderhellar/hellar/core"
	"github.com/hellarcore/tenderhellar/light/provider"
	"github.com/hellarcore/tenderhellar/light/provider/http"
	"github.com/hellarcore/tenderhellar/light/store"
)

// NewHTTPClient initiates an instance of a light client using HTTP addresses
// for both the primary provider and witnesses of the light client. A trusted
// header and hash must be passed to initialize the client.
//
// See all Option(s) for the additional configuration.
// See NewClient.
func NewHTTPClient(
	ctx context.Context,
	chainID string,
	primaryAddress string,
	witnessesAddresses []string,
	trustedStore store.Store,
	hellarCoreRPCClient hellarcore.QuorumVerifier,
	options ...Option) (*Client, error) {

	providers, err := providersFromAddresses(append(witnessesAddresses, primaryAddress), chainID)
	if err != nil {
		return nil, err
	}

	return NewClient(
		ctx,
		chainID,
		providers[len(providers)-1],
		providers[:len(providers)-1],
		trustedStore,
		hellarCoreRPCClient,
		options...)
}

// NewHTTPClientFromTrustedStore initiates an instance of a light client using
// HTTP addresses for both the primary provider and witnesses and uses a
// trusted store as the root of trust.
//
// See all Option(s) for the additional configuration.
// See NewClientFromTrustedStore.
func NewHTTPClientFromTrustedStore(
	chainID string,
	primaryAddress string,
	witnessesAddresses []string,
	trustedStore store.Store,
	hellarCoreRPCClient hellarcore.Client,
	options ...Option) (*Client, error) {

	providers, err := providersFromAddresses(append(witnessesAddresses, primaryAddress), chainID)
	if err != nil {
		return nil, err
	}

	return NewClientFromTrustedStore(
		chainID,
		providers[len(providers)-1],
		providers[:len(providers)-1],
		trustedStore,
		hellarCoreRPCClient,
		options...)
}

func providersFromAddresses(addrs []string, chainID string) ([]provider.Provider, error) {
	providers := make([]provider.Provider, len(addrs))
	for idx, address := range addrs {
		p, err := http.New(chainID, address)
		if err != nil {
			return nil, err
		}
		providers[idx] = p
	}
	return providers, nil
}
