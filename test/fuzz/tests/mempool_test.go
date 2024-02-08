//go:build gofuzz || go1.18

package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	abciclient "github.com/hellarcore/tenderhellar/abci/client"
	"github.com/hellarcore/tenderhellar/abci/example/kvstore"
	"github.com/hellarcore/tenderhellar/config"
	"github.com/hellarcore/tenderhellar/internal/mempool"
	"github.com/hellarcore/tenderhellar/libs/log"
)

func FuzzMempool(f *testing.F) {
	app, err := kvstore.NewMemoryApp()
	require.NoError(f, err)

	logger := log.NewNopLogger()
	conn := abciclient.NewLocalClient(logger, app)
	err = conn.Start(context.TODO())
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mp := mempool.NewTxMempool(logger, cfg, conn)

	f.Fuzz(func(t *testing.T, data []byte) {
		_ = mp.CheckTx(context.Background(), data, nil, mempool.TxInfo{})
	})
}
