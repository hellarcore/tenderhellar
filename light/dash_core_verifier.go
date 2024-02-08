package light

import (
	"github.com/hellarcore/hellard-go/btcjson"
	rpc "github.com/hellarcore/hellard-go/rpcclient"
)

// HellarCoreVerifier is used to verify signatures of light blocks
type HellarCoreVerifier struct {
	endpoint          *rpc.Client
	host              string
	rpcUsername       string
	rpcPassword       string
	defaultQuorumType btcjson.LLMQType
}

// NewHellarCoreVerifierClient returns an instance of SignerClient.
// it will start the endpoint (if not already started)
func NewHellarCoreVerifierClient(
	host string,
	rpcUsername string,
	rpcPassword string,
	defaultQuorumType btcjson.LLMQType,
) (*HellarCoreVerifier, error) {
	// Connect to local hellar core RPC server using HTTP POST mode.
	connCfg := &rpc.ConnConfig{
		Host:         host,
		User:         rpcUsername,
		Pass:         rpcPassword,
		HTTPPostMode: true, // Hellar core only supports HTTP POST mode
		DisableTLS:   true, // Hellar core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpc.New(connCfg, nil)
	if err != nil {
		return nil, err
	}

	return &HellarCoreVerifier{
		endpoint:          client,
		host:              host,
		rpcUsername:       rpcUsername,
		rpcPassword:       rpcPassword,
		defaultQuorumType: defaultQuorumType,
	}, nil
}
