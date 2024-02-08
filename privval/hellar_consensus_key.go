package privval

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"

	"github.com/hellarcore/hellard-go/btcjson"

	tmcrypto "github.com/hellarcore/tenderhellar/crypto"
	"github.com/hellarcore/tenderhellar/types"
)

type hellarConsensusPrivateKey struct {
	quorumHash tmcrypto.QuorumHash
	privval    HellarPrivValidator
	quorumType btcjson.LLMQType
}

var _ tmcrypto.PrivKey = &hellarConsensusPrivateKey{}

func (key hellarConsensusPrivateKey) Bytes() []byte {
	quorumType := make([]byte, 8)
	binary.LittleEndian.PutUint64(quorumType, uint64(key.quorumType))
	ourType := key.TypeTag()

	bytes := make([]byte, len(key.quorumHash)+len(ourType)+8)
	bytes = append(bytes, quorumType...)
	bytes = append(bytes, []byte(ourType)...)
	bytes = append(bytes, key.quorumHash...)

	return bytes
}

// Sign implements tmcrypto.PrivKey
func (key hellarConsensusPrivateKey) Sign(messageBytes []byte) ([]byte, error) {
	messageHash := tmcrypto.Checksum(messageBytes)

	return key.SignDigest(messageHash)
}

// SignDigest implements tmcrypto.PrivKey
func (key hellarConsensusPrivateKey) SignDigest(messageHash []byte) ([]byte, error) {
	requestIDhash := messageHash
	decodedSignature, _, err := key.privval.QuorumSign(context.TODO(), messageHash, requestIDhash, key.quorumType, key.quorumHash)
	return decodedSignature, err
}

// PubKey implements tmcrypto.PrivKey
func (key hellarConsensusPrivateKey) PubKey() tmcrypto.PubKey {
	pubkey, err := key.privval.GetPubKey(context.TODO(), key.quorumHash)
	if err != nil {
		panic("cannot retrieve public key: " + err.Error()) // not nice, but this iface doesn;t support error handling
	}

	// return NewHellarConsensusPublicKey(pubkey, key.quorumHash, key.quorumType)
	return pubkey
}

// Equals implements tmcrypto.PrivKey
func (key hellarConsensusPrivateKey) Equals(other tmcrypto.PrivKey) bool {
	return bytes.Equal(key.Bytes(), other.Bytes())
}

// Type implements tmcrypto.PrivKey
func (key hellarConsensusPrivateKey) Type() string {
	return "hellarCoreRPCPrivateKey"
}

// TypeTag implements jsontypes.Tagged interface.
func (key hellarConsensusPrivateKey) TypeTag() string { return key.Type() }

func (key hellarConsensusPrivateKey) String() string {
	return fmt.Sprintf("%s(quorumHash:%s,quorumType:%d)", key.Type(), key.quorumHash.ShortString(), key.quorumType)
}

// HellarConesensusPublicKey is a public key that constructs SignID in the background, to avoid this additional step
// when verifying signatures.
type HellarConsensusPublicKey struct {
	tmcrypto.PubKey

	quorumHash tmcrypto.QuorumHash
	quorumType btcjson.LLMQType
}

var _ tmcrypto.PubKey = HellarConsensusPublicKey{}

// NewHellarConsensusPublicKey wraps a public key with transparent handling of SignID according to DIP-7
func NewHellarConsensusPublicKey(baseKey tmcrypto.PubKey, quorumHash tmcrypto.QuorumHash, quorumType btcjson.LLMQType) *HellarConsensusPublicKey {
	if key, ok := baseKey.(*HellarConsensusPublicKey); ok {
		// don't wrap ourselves, but allow change of quorum hash and type
		baseKey = key.PubKey
	}

	return &HellarConsensusPublicKey{
		PubKey:     baseKey,
		quorumHash: quorumHash,
		quorumType: quorumType,
	}
}

func (pub HellarConsensusPublicKey) VerifySignature(msg []byte, sig []byte) bool {
	hash := tmcrypto.Checksum(msg)
	return pub.VerifySignatureDigest(hash, sig)
}
func (pub HellarConsensusPublicKey) VerifySignatureDigest(hash []byte, sig []byte) bool {
	signID := types.NewSignItemFromHash(pub.quorumType, pub.quorumHash, hash, hash).SignHash

	return pub.PubKey.VerifySignatureDigest(signID, sig)
}
