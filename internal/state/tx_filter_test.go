package state_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sm "github.com/hellarcore/tenderhellar/internal/state"
	"github.com/hellarcore/tenderhellar/internal/test/factory"
	tmrand "github.com/hellarcore/tenderhellar/libs/rand"
	"github.com/hellarcore/tenderhellar/types"
)

func TestTxFilter(t *testing.T) {
	const maxBlockBytes = 3241
	maxTxSize := maxBlockBytes - 1131
	genDoc := factory.MinimalGenesisDoc()
	genDoc.ConsensusParams.Block.MaxBytes = maxBlockBytes
	genDoc.ConsensusParams.Evidence.MaxBytes = 1500

	// Max size of Txs is much smaller than size of block,
	// since we need to account for commits and evidence.

	testCases := []struct {
		bytes int
		isErr bool
	}{
		{0, false},
		{maxTxSize - 1, false},
		{maxTxSize, false},
		{maxTxSize + 1, true},
		{3000, true},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			state, err := sm.MakeGenesisState(&genDoc)
			require.NoError(t, err)
			tx := types.Tx(tmrand.Bytes(tc.bytes))
			// Read MaxDataBytes, for debugging/logging only
			maxDataBytes, err := types.MaxDataBytes(state.ConsensusParams.Block.MaxBytes, nil, 0)
			require.NoError(t, err)

			f := sm.TxPreCheckForState(state)
			if tc.isErr {
				assert.NotNil(t, f(tx), "%+v, maxDataBytes:%d", tc, maxDataBytes)
			} else {
				assert.Nil(t, f(tx), "%+v, maxDataBytes:%d", tc, maxDataBytes)
			}
		})
	}
}
