package privval

import (
	"github.com/go-pkgz/jrpc"

	"github.com/hellarcore/tenderhellar/crypto"
	"github.com/hellarcore/tenderhellar/types"
)

type HellarCoreMockSignerServer struct {
	server     *jrpc.Server
	chainID    string
	quorumHash crypto.QuorumHash
	privVal    types.PrivValidator

	// handlerMtx tmsync.Mutex
}

func NewHellarCoreMockSignerServer(
	_endpoint *SignerDialerEndpoint,
	chainID string,
	quorumHash crypto.QuorumHash,
	privVal types.PrivValidator,
) *HellarCoreMockSignerServer {
	// create plugin (jrpc server)
	mockServer := &HellarCoreMockSignerServer{
		server: &jrpc.Server{
			API:        "/command",     // base url for rpc calls
			AuthUser:   "user",         // basic auth user name
			AuthPasswd: "password",     // basic auth password
			AppName:    "hellarcoremock", // plugin name for headers
		},
		chainID:    chainID,
		quorumHash: quorumHash,
		privVal:    privVal,
	}

	return mockServer
}

// OnStart implements service.Service.
func (ss *HellarCoreMockSignerServer) Run(port int) error {
	return ss.server.Run(port)
}
