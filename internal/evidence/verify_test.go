package evidence_test

import (
	"context"
	"testing"
	"time"

	"github.com/hellarcore/hellard-go/btcjson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	dbm "github.com/tendermint/tm-db"

	"github.com/hellarcore/tenderhellar/crypto"
	"github.com/hellarcore/tenderhellar/internal/eventbus"
	"github.com/hellarcore/tenderhellar/internal/evidence"
	"github.com/hellarcore/tenderhellar/internal/evidence/mocks"
	sm "github.com/hellarcore/tenderhellar/internal/state"
	smmocks "github.com/hellarcore/tenderhellar/internal/state/mocks"
	"github.com/hellarcore/tenderhellar/libs/log"
	tmproto "github.com/hellarcore/tenderhellar/proto/tendermint/types"
	"github.com/hellarcore/tenderhellar/types"
)

type voteData struct {
	vote1 *types.Vote
	vote2 *types.Vote
	valid bool
}

func TestVerifyDuplicateVoteEvidence(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := log.NewTestingLogger(t)

	quorumType := crypto.SmallQuorumType()
	quorumHash := crypto.RandQuorumHash()
	val := types.NewMockPVForQuorum(quorumHash)
	val2 := types.NewMockPVForQuorum(quorumHash)
	validator1 := val.ExtractIntoValidator(context.Background(), quorumHash)
	valSet := types.NewValidatorSet([]*types.Validator{validator1}, validator1.PubKey, quorumType, quorumHash, true)
	stateID := types.RandStateID()
	blockID := makeBlockID([]byte("blockhash"), 1000, []byte("partshash"), stateID)
	blockID2 := makeBlockID([]byte("blockhash2"), 1000, []byte("partshash"), stateID)
	blockID3 := makeBlockID([]byte("blockhash"), 10000, []byte("partshash"), stateID)
	blockID4 := makeBlockID([]byte("blockhash"), 10000, []byte("partshash2"), stateID)

	const chainID = "mychain"

	vote1 := makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID, quorumType, quorumHash)
	v1 := vote1.ToProto()
	err := val.SignVote(ctx, chainID, quorumType, quorumHash, v1, nil)
	require.NoError(t, err)
	badVote := makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID, quorumType, quorumHash)
	bv := badVote.ToProto()
	err = val2.SignVote(ctx, chainID, crypto.SmallQuorumType(), quorumHash, bv, nil)
	require.NoError(t, err)

	vote1.BlockSignature = v1.BlockSignature
	badVote.BlockSignature = bv.BlockSignature

	cases := []voteData{
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID2, quorumType, quorumHash), true}, // different block ids
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID3, quorumType, quorumHash), true},
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID4, quorumType, quorumHash), true},
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID, quorumType, quorumHash), false},     // wrong block id
		{vote1, makeVote(ctx, t, val, "mychain2", 0, 10, 2, 1, blockID2, quorumType, quorumHash), false}, // wrong chain id
		{vote1, makeVote(ctx, t, val, chainID, 0, 11, 2, 1, blockID2, quorumType, quorumHash), false},    // wrong height
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 3, 1, blockID2, quorumType, quorumHash), false},    // wrong round
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 2, 2, blockID2, quorumType, quorumHash), false},    // wrong step
		{vote1, makeVote(ctx, t, val2, chainID, 0, 10, 2, 1, blockID2, quorumType, quorumHash), false},   // wrong validator
		// a different vote time doesn't matter
		{vote1, makeVote(ctx, t, val, chainID, 0, 10, 2, 1, blockID2, quorumType, quorumHash), true},
		{vote1, badVote, false}, // signed by wrong key
	}

	require.NoError(t, err)
	for _, c := range cases {
		ev := &types.DuplicateVoteEvidence{
			VoteA:            c.vote1,
			VoteB:            c.vote2,
			ValidatorPower:   1,
			TotalVotingPower: 1,
			Timestamp:        defaultEvidenceTime,
		}
		if c.valid {
			assert.Nil(t, evidence.VerifyDuplicateVote(ev, chainID, valSet), "evidence should be valid")
		} else {
			assert.NotNil(t, evidence.VerifyDuplicateVote(ev, chainID, valSet), "evidence should be invalid")
		}
	}

	// create good evidence and correct validator power
	goodEv, err := types.NewMockDuplicateVoteEvidenceWithValidator(ctx, 10, defaultEvidenceTime, val, chainID, crypto.SmallQuorumType(), quorumHash)
	require.NoError(t, err)
	goodEv.ValidatorPower = types.DefaultHellarVotingPower
	goodEv.TotalVotingPower = types.DefaultHellarVotingPower
	badEv, err := types.NewMockDuplicateVoteEvidenceWithValidator(ctx, 10, defaultEvidenceTime, val, chainID, crypto.SmallQuorumType(), quorumHash)
	require.NoError(t, err)
	badEv.ValidatorPower = types.DefaultHellarVotingPower + 1
	badEv.TotalVotingPower = types.DefaultHellarVotingPower
	badTimeEv, err := types.NewMockDuplicateVoteEvidenceWithValidator(ctx, 10, defaultEvidenceTime.Add(1*time.Minute), val, chainID, crypto.SmallQuorumType(), quorumHash)
	require.NoError(t, err)
	badTimeEv.ValidatorPower = types.DefaultHellarVotingPower
	badTimeEv.TotalVotingPower = types.DefaultHellarVotingPower
	state := sm.State{
		ChainID:         chainID,
		LastBlockTime:   defaultEvidenceTime.Add(1 * time.Minute),
		LastBlockHeight: 11,
		ConsensusParams: *types.DefaultConsensusParams(),
	}
	stateStore := &smmocks.Store{}
	stateStore.On("LoadValidators", int64(10)).Return(valSet, nil)
	stateStore.On("Load").Return(state, nil)
	blockStore := &mocks.BlockStore{}
	blockStore.On("LoadBlockMeta", int64(10)).Return(&types.BlockMeta{Header: types.Header{Time: defaultEvidenceTime}})

	eventBus := eventbus.NewDefault(logger)
	require.NoError(t, eventBus.Start(ctx))

	pool := evidence.NewPool(logger, dbm.NewMemDB(), stateStore, blockStore, evidence.NopMetrics(), eventBus)
	startPool(t, pool, stateStore)

	evList := types.EvidenceList{goodEv}
	err = pool.CheckEvidence(ctx, evList)
	assert.NoError(t, err)

	// evidence with a different validator power should fail
	evList = types.EvidenceList{badEv}
	err = pool.CheckEvidence(ctx, evList)
	assert.Error(t, err)

	// evidence with a different timestamp should fail
	evList = types.EvidenceList{badTimeEv}
	err = pool.CheckEvidence(ctx, evList)
	assert.Error(t, err)
}

func makeVote(
	ctx context.Context,
	t *testing.T, val types.PrivValidator, chainID string, valIndex int32, height int64,
	round int32, step int, blockID types.BlockID, quorumType btcjson.LLMQType, quorumHash crypto.QuorumHash) *types.Vote {
	proTxHash, err := val.GetProTxHash(ctx)
	require.NoError(t, err)
	v := &types.Vote{
		ValidatorProTxHash: proTxHash,
		ValidatorIndex:     valIndex,
		Height:             height,
		Round:              round,
		Type:               tmproto.SignedMsgType(step),
		BlockID:            blockID,
	}

	vpb := v.ToProto()
	err = val.SignVote(ctx, chainID, quorumType, quorumHash, vpb, nil)
	require.NoError(t, err)
	v.BlockSignature = vpb.BlockSignature
	return v
}

func makeBlockID(hash []byte, partSetSize uint32, partSetHash []byte, stateID tmproto.StateID) types.BlockID {
	var (
		h   = make([]byte, crypto.HashSize)
		psH = make([]byte, crypto.HashSize)
	)
	copy(h, hash)
	copy(psH, partSetHash)
	return types.BlockID{
		Hash: h,
		PartSetHeader: types.PartSetHeader{
			Total: partSetSize,
			Hash:  psH,
		},
		StateID: stateID.Hash(),
	}
}
