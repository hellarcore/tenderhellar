// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/hellarcore/tenderhellar/internal/consensus/types"
	mock "github.com/stretchr/testify/mock"
)

// Gossiper is an autogenerated mock type for the Gossiper type
type Gossiper struct {
	mock.Mock
}

// GossipBlockPartsForCatchup provides a mock function with given fields: ctx, rs, prs
func (_m *Gossiper) GossipBlockPartsForCatchup(ctx context.Context, rs types.RoundState, prs *types.PeerRoundState) {
	_m.Called(ctx, rs, prs)
}

// GossipCommit provides a mock function with given fields: ctx, rs, prs
func (_m *Gossiper) GossipCommit(ctx context.Context, rs types.RoundState, prs *types.PeerRoundState) {
	_m.Called(ctx, rs, prs)
}

// GossipProposal provides a mock function with given fields: ctx, rs, prs
func (_m *Gossiper) GossipProposal(ctx context.Context, rs types.RoundState, prs *types.PeerRoundState) {
	_m.Called(ctx, rs, prs)
}

// GossipProposalBlockParts provides a mock function with given fields: ctx, rs, prs
func (_m *Gossiper) GossipProposalBlockParts(ctx context.Context, rs types.RoundState, prs *types.PeerRoundState) {
	_m.Called(ctx, rs, prs)
}

// GossipVote provides a mock function with given fields: ctx, rs, prs
func (_m *Gossiper) GossipVote(ctx context.Context, rs types.RoundState, prs *types.PeerRoundState) {
	_m.Called(ctx, rs, prs)
}

// GossipVoteSetMaj23 provides a mock function with given fields: ctx, rs, prs
func (_m *Gossiper) GossipVoteSetMaj23(ctx context.Context, rs types.RoundState, prs *types.PeerRoundState) {
	_m.Called(ctx, rs, prs)
}

// NewGossiper creates a new instance of Gossiper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGossiper(t interface {
	mock.TestingT
	Cleanup(func())
}) *Gossiper {
	mock := &Gossiper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
