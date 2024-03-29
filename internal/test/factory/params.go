package factory

import (
	"time"

	"github.com/hellarcore/tenderhellar/types"
)

// ConsensusParams returns a default set of ConsensusParams that are suitable
// for use in testing
func ConsensusParams(opts ...func(*types.ConsensusParams)) *types.ConsensusParams {
	c := types.DefaultConsensusParams()
	c.Timeout = types.TimeoutParams{
		Commit:              10 * time.Millisecond,
		Propose:             40 * time.Millisecond,
		ProposeDelta:        1 * time.Millisecond,
		Vote:                10 * time.Millisecond,
		VoteDelta:           1 * time.Millisecond,
		BypassCommitTimeout: true,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
