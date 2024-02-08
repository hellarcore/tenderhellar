package blocksync

import (
	p2pproto "github.com/hellarcore/tenderhellar/proto/tendermint/p2p"
	"github.com/hellarcore/tenderhellar/types"
)

const (
	MaxMsgSize = types.MaxBlockSizeBytes +
		p2pproto.BlockResponseMessagePrefixSize +
		p2pproto.BlockResponseMessageFieldKeySize
)
