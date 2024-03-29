package types

import (
	"bytes"
	"fmt"

	bls "github.com/hellarcore/bls-signatures/go-bindings"
	"github.com/hellarcore/hellard-go/btcjson"
	"github.com/hellarcore/tenderhellar/crypto"
	tmbytes "github.com/hellarcore/tenderhellar/libs/bytes"
	"github.com/hellarcore/tenderhellar/proto/tendermint/types"
	"github.com/rs/zerolog"
)

// QuorumSignData holds data which is necessary for signing and verification block, state, and each vote-extension in a list
type QuorumSignData struct {
	Block                  SignItem
	VoteExtensionSignItems []SignItem
}

// Signs items inside QuorumSignData using a given private key.
//
// Mainly for testing.
func (q QuorumSignData) SignWithPrivkey(key crypto.PrivKey) (QuorumSigns, error) {
	var err error
	var signs QuorumSigns
	if signs.BlockSign, err = key.SignDigest(q.Block.SignHash); err != nil {
		return signs, err
	}

	signs.VoteExtensionSignatures = make([][]byte, 0, len(q.VoteExtensionSignItems))
	for _, item := range q.VoteExtensionSignItems {
		var sign []byte
		if sign, err = key.SignDigest(item.SignHash); err != nil {
			return signs, err
		}
		signs.VoteExtensionSignatures = append(signs.VoteExtensionSignatures, sign)
	}

	return signs, nil
}

// Verify verifies a quorum signatures: block, state and vote-extensions
func (q QuorumSignData) Verify(pubKey crypto.PubKey, signs QuorumSigns) error {
	return NewQuorumSignsVerifier(q).Verify(pubKey, signs)
}

// MakeQuorumSignsWithVoteSet creates and returns QuorumSignData struct built with a vote-set and an added vote
func MakeQuorumSignsWithVoteSet(voteSet *VoteSet, vote *types.Vote) (QuorumSignData, error) {
	return MakeQuorumSigns(
		voteSet.chainID,
		voteSet.valSet.QuorumType,
		voteSet.valSet.QuorumHash,
		vote,
	)
}

// MakeQuorumSigns builds signing data for block, state and vote-extensions
// each a sign-id item consist of request-id, raw data, hash of raw and id
func MakeQuorumSigns(
	chainID string,
	quorumType btcjson.LLMQType,
	quorumHash crypto.QuorumHash,
	protoVote *types.Vote,
) (QuorumSignData, error) {
	quorumSign := QuorumSignData{
		Block: MakeBlockSignItem(chainID, protoVote, quorumType, quorumHash),
	}
	var err error
	quorumSign.VoteExtensionSignItems, err =
		VoteExtensionsFromProto(protoVote.VoteExtensions...).
			Filter(func(ext VoteExtensionIf) bool {
				return ext.IsThresholdRecoverable()
			}).
			SignItems(chainID, quorumType, quorumHash, protoVote.Height, protoVote.Round)
	if err != nil {
		return QuorumSignData{}, err
	}
	return quorumSign, nil
}

// MakeBlockSignItem creates SignItem struct for a block
func MakeBlockSignItem(chainID string, vote *types.Vote, quorumType btcjson.LLMQType, quorumHash []byte) SignItem {
	reqID := BlockRequestID(vote.Height, vote.Round)
	raw, err := vote.SignBytes(chainID)
	if err != nil {
		panic(fmt.Errorf("block sign item: %w", err))
	}
	return NewSignItem(quorumType, quorumHash, reqID, raw)
}

// BlockRequestID returns a block request ID
func BlockRequestID(height int64, round int32) []byte {
	return heightRoundRequestID("dpbvote", height, round)
}

// SignItem represents signing session data (in field SignItem.ID) that will be signed to get threshold signature share.
// Field names are the same as in Hellar Core, but the meaning is different.
// See DIP-0007
type SignItem struct {
	LlmqType   btcjson.LLMQType // Quorum type for which this sign item is created
	ID         []byte           // Request ID for quorum signing
	MsgHash    []byte           // Checksum of Raw
	QuorumHash []byte           // Quorum hash for which this sign item is created

	SignHash []byte // Hash of llmqType, quorumHash, id, and msgHash - as provided to crypto sign/verify functions

	Msg []byte // Raw data to be signed, before any transformations; optional
}

// Validate validates prepared data for signing
func (i *SignItem) Validate() error {
	if len(i.ID) != crypto.DefaultHashSize {
		return fmt.Errorf("invalid request ID size: %X", i.ID)
	}
	if len(i.MsgHash) != crypto.DefaultHashSize {
		return fmt.Errorf("invalid hash size %d: %X", len(i.MsgHash), i.MsgHash)
	}
	if len(i.QuorumHash) != crypto.DefaultHashSize {
		return fmt.Errorf("invalid quorum hash size %d: %X", len(i.QuorumHash), i.QuorumHash)
	}
	// Msg is optional
	if len(i.Msg) > 0 {
		if !bytes.Equal(crypto.Checksum(i.Msg), i.MsgHash) {
			return fmt.Errorf("invalid hash %X for raw data: %X", i.MsgHash, i.Msg)
		}
	}
	return nil
}

func (i SignItem) MarshalZerologObject(e *zerolog.Event) {
	e.Hex("msg", i.Msg)
	e.Hex("signRequestID", i.ID)
	e.Hex("signID", i.SignHash)
	e.Hex("msgHash", i.MsgHash)
	e.Hex("quorumHash", i.QuorumHash)
	e.Uint8("llmqType", uint8(i.LlmqType))

}

// NewSignItem creates a new instance of SignItem with calculating a hash for a raw and creating signID
//
// Arguments:
// - quorumType: quorum type
// - quorumHash: quorum hash
// - reqID: sign request ID
// - msg: raw data to be signed; it will be hashed with crypto.Checksum()
func NewSignItem(quorumType btcjson.LLMQType, quorumHash, reqID, msg []byte) SignItem {
	msgHash := crypto.Checksum(msg) // FIXME: shouldn't we use sha256(sha256(raw)) here?
	item := NewSignItemFromHash(quorumType, quorumHash, reqID, msgHash)
	item.Msg = msg

	return item
}

// Create a new sign item without raw value, using provided hash.
func NewSignItemFromHash(quorumType btcjson.LLMQType, quorumHash, reqID, msgHash []byte) SignItem {
	item := SignItem{
		ID:         reqID,
		MsgHash:    msgHash,
		LlmqType:   quorumType,
		QuorumHash: quorumHash,
		Msg:        nil, // Raw is empty, as we don't have it
	}

	// By default, reverse fields when calculating SignHash
	item.UpdateSignHash(true)

	return item
}

// UpdateSignHash recalculates signHash field
// If reverse is true, then all []byte elements will be reversed before
// calculating signID
func (i *SignItem) UpdateSignHash(reverse bool) {
	if err := i.Validate(); err != nil {
		panic("invalid sign item: " + err.Error())
	}
	llmqType := i.LlmqType

	quorumHash := i.QuorumHash
	requestID := i.ID
	messageHash := i.MsgHash

	if reverse {
		quorumHash = tmbytes.Reverse(quorumHash)
		requestID = tmbytes.Reverse(requestID)
		messageHash = tmbytes.Reverse(messageHash)
	}

	var blsQuorumHash bls.Hash
	copy(blsQuorumHash[:], quorumHash)

	var blsRequestID bls.Hash
	copy(blsRequestID[:], requestID)

	var blsMessageHash bls.Hash
	copy(blsMessageHash[:], messageHash)

	// fmt.Printf("LlmqType: %x + ", llmqType)
	// fmt.Printf("QuorumHash: %x + ", blsQuorumHash)
	// fmt.Printf("RequestID: %x + ", blsRequestID)
	// fmt.Printf("MsgHash: %x\n", blsMessageHash)

	blsSignHash := bls.BuildSignHash(uint8(llmqType), blsQuorumHash, blsRequestID, blsMessageHash)

	signHash := make([]byte, 32)
	copy(signHash, blsSignHash[:])

	i.SignHash = signHash
}
