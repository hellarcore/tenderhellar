package p2ptest

import (
	"github.com/hellarcore/tenderhellar/proto/tendermint/p2p"
	"github.com/hellarcore/tenderhellar/types"
)

// Message is a simple message containing a string-typed Value field.
type Message = p2p.Echo

func NodeInSlice(id types.NodeID, ids []types.NodeID) bool {
	for _, n := range ids {
		if id == n {
			return true
		}
	}
	return false
}
