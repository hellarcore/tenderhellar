// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	state "github.com/hellarcore/tenderhellar/internal/state"
	mock "github.com/stretchr/testify/mock"

	types "github.com/hellarcore/tenderhellar/types"
)

// Executor is an autogenerated mock type for the Executor type
type Executor struct {
	mock.Mock
}

// ApplyBlock provides a mock function with given fields: ctx, _a1, blockID, block, commit
func (_m *Executor) ApplyBlock(ctx context.Context, _a1 state.State, blockID types.BlockID, block *types.Block, commit *types.Commit) (state.State, error) {
	ret := _m.Called(ctx, _a1, blockID, block, commit)

	var r0 state.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.State, types.BlockID, *types.Block, *types.Commit) (state.State, error)); ok {
		return rf(ctx, _a1, blockID, block, commit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.State, types.BlockID, *types.Block, *types.Commit) state.State); ok {
		r0 = rf(ctx, _a1, blockID, block, commit)
	} else {
		r0 = ret.Get(0).(state.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.State, types.BlockID, *types.Block, *types.Commit) error); ok {
		r1 = rf(ctx, _a1, blockID, block, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProposalBlock provides a mock function with given fields: ctx, height, round, _a3, commit, proposerProTxHash, proposedAppVersion
func (_m *Executor) CreateProposalBlock(ctx context.Context, height int64, round int32, _a3 state.State, commit *types.Commit, proposerProTxHash []byte, proposedAppVersion uint64) (*types.Block, state.CurrentRoundState, error) {
	ret := _m.Called(ctx, height, round, _a3, commit, proposerProTxHash, proposedAppVersion)

	var r0 *types.Block
	var r1 state.CurrentRoundState
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int32, state.State, *types.Commit, []byte, uint64) (*types.Block, state.CurrentRoundState, error)); ok {
		return rf(ctx, height, round, _a3, commit, proposerProTxHash, proposedAppVersion)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int32, state.State, *types.Commit, []byte, uint64) *types.Block); ok {
		r0 = rf(ctx, height, round, _a3, commit, proposerProTxHash, proposedAppVersion)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int32, state.State, *types.Commit, []byte, uint64) state.CurrentRoundState); ok {
		r1 = rf(ctx, height, round, _a3, commit, proposerProTxHash, proposedAppVersion)
	} else {
		r1 = ret.Get(1).(state.CurrentRoundState)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int64, int32, state.State, *types.Commit, []byte, uint64) error); ok {
		r2 = rf(ctx, height, round, _a3, commit, proposerProTxHash, proposedAppVersion)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ExtendVote provides a mock function with given fields: ctx, vote
func (_m *Executor) ExtendVote(ctx context.Context, vote *types.Vote) {
	_m.Called(ctx, vote)
}

// FinalizeBlock provides a mock function with given fields: ctx, _a1, uncommittedState, blockID, block, commit
func (_m *Executor) FinalizeBlock(ctx context.Context, _a1 state.State, uncommittedState state.CurrentRoundState, blockID types.BlockID, block *types.Block, commit *types.Commit) (state.State, error) {
	ret := _m.Called(ctx, _a1, uncommittedState, blockID, block, commit)

	var r0 state.State
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, state.State, state.CurrentRoundState, types.BlockID, *types.Block, *types.Commit) (state.State, error)); ok {
		return rf(ctx, _a1, uncommittedState, blockID, block, commit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, state.State, state.CurrentRoundState, types.BlockID, *types.Block, *types.Commit) state.State); ok {
		r0 = rf(ctx, _a1, uncommittedState, blockID, block, commit)
	} else {
		r0 = ret.Get(0).(state.State)
	}

	if rf, ok := ret.Get(1).(func(context.Context, state.State, state.CurrentRoundState, types.BlockID, *types.Block, *types.Commit) error); ok {
		r1 = rf(ctx, _a1, uncommittedState, blockID, block, commit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProcessProposal provides a mock function with given fields: ctx, block, round, _a3, verify
func (_m *Executor) ProcessProposal(ctx context.Context, block *types.Block, round int32, _a3 state.State, verify bool) (state.CurrentRoundState, error) {
	ret := _m.Called(ctx, block, round, _a3, verify)

	var r0 state.CurrentRoundState
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Block, int32, state.State, bool) (state.CurrentRoundState, error)); ok {
		return rf(ctx, block, round, _a3, verify)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.Block, int32, state.State, bool) state.CurrentRoundState); ok {
		r0 = rf(ctx, block, round, _a3, verify)
	} else {
		r0 = ret.Get(0).(state.CurrentRoundState)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.Block, int32, state.State, bool) error); ok {
		r1 = rf(ctx, block, round, _a3, verify)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateBlock provides a mock function with given fields: ctx, _a1, block
func (_m *Executor) ValidateBlock(ctx context.Context, _a1 state.State, block *types.Block) error {
	ret := _m.Called(ctx, _a1, block)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.State, *types.Block) error); ok {
		r0 = rf(ctx, _a1, block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateBlockWithRoundState provides a mock function with given fields: ctx, _a1, uncommittedState, block
func (_m *Executor) ValidateBlockWithRoundState(ctx context.Context, _a1 state.State, uncommittedState state.CurrentRoundState, block *types.Block) error {
	ret := _m.Called(ctx, _a1, uncommittedState, block)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, state.State, state.CurrentRoundState, *types.Block) error); ok {
		r0 = rf(ctx, _a1, uncommittedState, block)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyVoteExtension provides a mock function with given fields: ctx, vote
func (_m *Executor) VerifyVoteExtension(ctx context.Context, vote *types.Vote) error {
	ret := _m.Called(ctx, vote)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Vote) error); ok {
		r0 = rf(ctx, vote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewExecutor creates a new instance of Executor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExecutor(t interface {
	mock.TestingT
	Cleanup(func())
}) *Executor {
	mock := &Executor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
