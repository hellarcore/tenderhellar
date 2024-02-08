package core

import (
	"context"

	"github.com/hellarcore/tenderhellar/rpc/coretypes"
)

// UnsafeFlushMempool removes all transactions from the mempool.
func (env *Environment) UnsafeFlushMempool(_ctx context.Context) (*coretypes.ResultUnsafeFlushMempool, error) {
	env.Mempool.Flush()
	return &coretypes.ResultUnsafeFlushMempool{}, nil
}
