package quorum

import (
	"github.com/hellarcore/tenderhellar/internal/p2p"
	"github.com/hellarcore/tenderhellar/types"
)

// nodeAddress converts ValidatorAddress to a NodeAddress object
func nodeAddress(va types.ValidatorAddress) p2p.NodeAddress {
	nodeAddress := p2p.NodeAddress{
		NodeID:   va.NodeID,
		Protocol: p2p.TCPProtocol,
		Hostname: va.Hostname,
		Port:     va.Port,
	}
	return nodeAddress
}
