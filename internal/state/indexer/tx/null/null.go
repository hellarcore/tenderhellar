package null

import (
	"context"
	"errors"

	abci "github.com/hellarcore/tenderhellar/abci/types"
	"github.com/hellarcore/tenderhellar/internal/pubsub/query"
	"github.com/hellarcore/tenderhellar/internal/state/indexer"
)

var _ indexer.TxIndexer = (*TxIndex)(nil)

// TxIndex acts as a /dev/null.
type TxIndex struct{}

// Get on a TxIndex is disabled and panics when invoked.
func (txi *TxIndex) Get(_hash []byte) (*abci.TxResult, error) {
	return nil, errors.New(`indexing is disabled (set 'tx_index = "kv"' in config)`)
}

// AddBatch is a noop and always returns nil.
func (txi *TxIndex) AddBatch(_batch *indexer.Batch) error {
	return nil
}

// Index is a noop and always returns nil.
func (txi *TxIndex) Index(_results []*abci.TxResult) error {
	return nil
}

func (txi *TxIndex) Search(_ctx context.Context, _q *query.Query) ([]*abci.TxResult, error) {
	return []*abci.TxResult{}, nil
}
