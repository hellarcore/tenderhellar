package e2e_test

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hellarcore/tenderhellar/abci/example/code"
	tmrand "github.com/hellarcore/tenderhellar/libs/rand"
	"github.com/hellarcore/tenderhellar/rpc/client/http"
	e2e "github.com/hellarcore/tenderhellar/test/e2e/pkg"
	"github.com/hellarcore/tenderhellar/types"
)

const (
	randomSeed = 4827085738
)

// Tests that any initial state given in genesis has made it into the app.
func TestApp_InitialState(t *testing.T) {
	testNode(t, func(ctx context.Context, t *testing.T, node e2e.Node) {
		if len(node.Testnet.InitialState.Items) == 0 {
			return
		}

		client, err := node.Client()
		require.NoError(t, err)
		for k, v := range node.Testnet.InitialState.Items {
			resp, err := client.ABCIQuery(ctx, "", []byte(k))
			require.NoError(t, err)
			assert.Equal(t, k, string(resp.Response.Key))
			assert.Equal(t, v, string(resp.Response.Value))
		}
	})
}

// Tests that the app hash (as reported by the app) matches the last
// block and the node sync status.
func TestApp_Hash(t *testing.T) {
	testNode(t, func(ctx context.Context, t *testing.T, node e2e.Node) {
		client, err := node.Client()
		require.NoError(t, err)

		info, err := client.ABCIInfo(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, info.Response.LastBlockAppHash, "expected app to return app hash")

		// In same-block execution, the app hash is stored in the same block
		blockHeight := info.Response.LastBlockHeight

		require.Eventually(t, func() bool {
			status, err := client.Status(ctx)
			require.NoError(t, err)
			require.NotZero(t, status.SyncInfo.LatestBlockHeight)
			return status.SyncInfo.LatestBlockHeight >= blockHeight
		}, 60*time.Second, 500*time.Millisecond)

		block, err := client.Block(ctx, &blockHeight)
		require.NoError(t, err)
		require.Equal(t, blockHeight, block.Block.Height)
		require.EqualValues(t,
			info.Response.LastBlockAppHash,
			block.Block.AppHash.Bytes(),
			"app hash does not match last block's app hash at height %d", blockHeight)
	})
}

// Tests that the app and blockstore have and report the same height.
func TestApp_Height(t *testing.T) {
	testNode(t, func(ctx context.Context, t *testing.T, node e2e.Node) {
		client, err := node.Client()
		require.NoError(t, err)
		info, err := client.ABCIInfo(ctx)
		require.NoError(t, err)
		require.NotZero(t, info.Response.LastBlockHeight)

		status, err := client.Status(ctx)
		require.NoError(t, err)
		require.NotZero(t, status.SyncInfo.LatestBlockHeight)

		block, err := client.Block(ctx, &info.Response.LastBlockHeight)
		require.NoError(t, err)

		require.Equal(t, info.Response.LastBlockHeight, block.Block.Height)

		require.True(t, status.SyncInfo.LatestBlockHeight >= info.Response.LastBlockHeight,
			"status out of sync with application")
	})
}

// Tests that we can set a value and retrieve it.
func TestApp_Tx(t *testing.T) {
	type broadcastFunc func(context.Context, types.Tx) error

	testCases := []struct {
		Name        string
		WaitTime    time.Duration
		BroadcastTx func(client *http.HTTP) broadcastFunc
		ShouldSkip  bool
	}{
		{
			Name:     "sync",
			WaitTime: time.Minute,
			BroadcastTx: func(client *http.HTTP) broadcastFunc {
				return func(ctx context.Context, tx types.Tx) error {
					_, err := client.BroadcastTxSync(ctx, tx)
					return err
				}
			},
		},
		{
			Name:     "flushMempool",
			WaitTime: 15 * time.Second,
			// TODO: turn this check back on if it can
			// return reliably. Currently these calls have
			// a hard timeout of 10s (server side
			// configured). The Sync check is probably
			// safe.
			ShouldSkip: true,
			BroadcastTx: func(client *http.HTTP) broadcastFunc {
				return func(ctx context.Context, tx types.Tx) error {
					_, err := client.BroadcastTxCommit(ctx, tx)
					return err
				}
			},
		},
		{
			Name:     "Async",
			WaitTime: 90 * time.Second,
			// TODO: turn this check back on if there's a
			// way to avoid failures in the case that the
			// transaction doesn't make it into the
			// mempool. (retries?)
			ShouldSkip: true,
			BroadcastTx: func(client *http.HTTP) broadcastFunc {
				return func(ctx context.Context, tx types.Tx) error {
					_, err := client.BroadcastTxAsync(ctx, tx)
					return err
				}
			},
		},
	}

	r := rand.New(rand.NewSource(randomSeed))
	for idx, test := range testCases {
		if test.ShouldSkip {
			continue
		}
		t.Run(test.Name, func(t *testing.T) {
			test := testCases[idx]
			testNode(t, func(ctx context.Context, t *testing.T, node e2e.Node) {
				client, err := node.Client()
				require.NoError(t, err)

				key := fmt.Sprintf("testapp-tx-%v", node.Name)
				value := tmrand.StrFromSource(r, 32)
				tx := types.Tx(fmt.Sprintf("%v=%v", key, value))

				err = test.BroadcastTx(client)(ctx, tx)
				require.NoError(t, err)

				hash := tx.Hash()

				require.Eventuallyf(t, func() bool {
					txResp, err := client.Tx(ctx, hash, false)
					return err == nil && bytes.Equal(txResp.Tx, tx)
				},
					test.WaitTime, // timeout
					time.Second,   // interval
					"submitted tx %X wasn't committed after %v",
					hash, test.WaitTime,
				)

				abciResp, err := client.ABCIQuery(ctx, "", []byte(key))
				require.NoError(t, err)
				assert.Equal(t, code.CodeTypeOK, abciResp.Response.Code)
				assert.Equal(t, key, string(abciResp.Response.Key))
				assert.Equal(t, value, string(abciResp.Response.Value))
			})

		})

	}

}

// Given transactions which take more than the block size,
// when I submit them to the node,
// then the first transaction should be committed before the last one.
func TestApp_TxTooBig(t *testing.T) {
	const timeout = 60 * time.Second

	testNode(t, func(ctx context.Context, t *testing.T, node e2e.Node) {
		session := rand.Int63()

		client, err := node.Client()
		require.NoError(t, err)

		// FIXME: ConsensusParams is broken for last height, this is just workaround
		status, err := client.Status(ctx)
		assert.NoError(t, err)
		cp, err := client.ConsensusParams(ctx, &status.SyncInfo.LatestBlockHeight)
		assert.NoError(t, err)

		// ensure we have more txs than fits the block
		TxPayloadSize := int(cp.ConsensusParams.Block.MaxBytes / 100) // 1% of block size
		numTxs := 101

		tx := make(types.Tx, TxPayloadSize) // first tx is just zeros

		var firstTxHash []byte
		var key string

		for i := 0; i < numTxs; i++ {
			key = fmt.Sprintf("testapp-big-tx-%v-%08x-%d=", node.Name, session, i)
			copy(tx, key)

			payloadOffset := len(tx) - 8 // where we put the `i` into the payload
			assert.Greater(t, payloadOffset, len(key))

			big.NewInt(int64(i)).FillBytes(tx[payloadOffset:])
			assert.Len(t, tx, TxPayloadSize)

			if i == 0 {
				firstTxHash = tx.Hash()
			}

			_, err = client.BroadcastTxAsync(ctx, tx)

			assert.NoError(t, err, "failed to broadcast tx %06x", i)
		}

		lastTxHash := tx.Hash()

		require.Eventuallyf(t, func() bool {
			// last tx should be committed later than first
			lastTxResp, err := client.Tx(ctx, lastTxHash, false)
			if err == nil {
				assert.Equal(t, lastTxHash, lastTxResp.Tx.Hash())

				// fetch first tx
				firstTxResp, err := client.Tx(ctx, firstTxHash, false)
				assert.NoError(t, err, "first tx should be committed before second")
				assert.EqualValues(t, firstTxHash, firstTxResp.Tx.Hash())

				firstTxBlock, err := client.Header(ctx, &firstTxResp.Height)
				assert.NoError(t, err)
				lastTxBlock, err := client.Header(ctx, &lastTxResp.Height)
				assert.NoError(t, err)

				t.Logf("first tx in block %d, last tx in block %d, time diff %s",
					firstTxResp.Height,
					lastTxResp.Height,
					lastTxBlock.Header.Time.Sub(firstTxBlock.Header.Time).String(),
				)

				assert.Less(t, firstTxResp.Height, lastTxResp.Height, "first tx should in block before last tx")
				return true
			}

			return false
		},
			timeout,     // timeout
			time.Second, // interval
			"submitted tx %X wasn't committed after %v",
			lastTxHash, timeout,
		)

		// abciResp, err := client.ABCIQuery(ctx, "", []byte(lastTxKey))
		// require.NoError(t, err)
		// assert.Equal(t, code.CodeTypeOK, abciResp.Response.Code)
		// assert.Equal(t, lastTxKey, string(abciResp.Response.Key))
		// assert.Equal(t, lastTxHash, types.Tx(abciResp.Response.Value).Hash())
	})

}
