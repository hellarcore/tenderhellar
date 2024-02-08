package types

import tmbytes "github.com/hellarcore/tenderhellar/libs/bytes"

// SetHellarParams sets hellar's some parameters to a block
// this method should call if we need to provide specific hellar data
func (b *Block) SetHellarParams(
	lastCoreChainLockedBlockHeight uint32,
	coreChainLock *CoreChainLock,
	proposedAppVersion uint64,
	nextValidatorsHash tmbytes.HexBytes,
) {
	if coreChainLock == nil {
		b.CoreChainLockedHeight = lastCoreChainLockedBlockHeight
	} else {
		b.CoreChainLockedHeight = coreChainLock.CoreBlockHeight
	}
	b.ProposedAppVersion = proposedAppVersion
	b.CoreChainLock = coreChainLock

	if len(nextValidatorsHash) > 0 {
		b.NextValidatorsHash = nextValidatorsHash
	}
}
