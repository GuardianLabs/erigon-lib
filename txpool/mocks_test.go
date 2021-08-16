// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package txpool

import (
	"github.com/ledgerwatch/erigon-lib/kv"
	"sync"
)

// Ensure, that PoolMock does implement Pool.
// If this is not the case, regenerate this file with moq.
var _ Pool = &PoolMock{}

// PoolMock is a mock implementation of Pool.
//
// 	func TestSomethingThatUsesPool(t *testing.T) {
//
// 		// make and configure a mocked Pool
// 		mockedPool := &PoolMock{
// 			AddFunc: func(db kv.Tx, newTxs TxSlots) error {
// 				panic("mock out the Add method")
// 			},
// 			AddNewGoodPeerFunc: func(peerID PeerID)  {
// 				panic("mock out the AddNewGoodPeer method")
// 			},
// 			GetRlpFunc: func(hash []byte) []byte {
// 				panic("mock out the GetRlp method")
// 			},
// 			IdHashKnownFunc: func(hash []byte) bool {
// 				panic("mock out the IdHashKnown method")
// 			},
// 			OnNewBlockFunc: func(db kv.Tx, stateChanges map[string]senderInfo, unwindTxs TxSlots, minedTxs TxSlots, protocolBaseFee uint64, pendingBaseFee uint64, blockHeight uint64) error {
// 				panic("mock out the OnNewBlock method")
// 			},
// 		}
//
// 		// use mockedPool in code that requires Pool
// 		// and then make assertions.
//
// 	}
type PoolMock struct {
	// AddFunc mocks the Add method.
	AddFunc func(db kv.Tx, newTxs TxSlots) error

	// AddNewGoodPeerFunc mocks the AddNewGoodPeer method.
	AddNewGoodPeerFunc func(peerID PeerID)

	// GetRlpFunc mocks the GetRlp method.
	GetRlpFunc func(hash []byte) []byte

	// IdHashKnownFunc mocks the IdHashKnown method.
	IdHashKnownFunc func(hash []byte) bool

	// OnNewBlockFunc mocks the OnNewBlock method.
	OnNewBlockFunc func(db kv.Tx, stateChanges map[string]senderInfo, unwindTxs TxSlots, minedTxs TxSlots, protocolBaseFee uint64, pendingBaseFee uint64, blockHeight uint64) error

	// calls tracks calls to the methods.
	calls struct {
		// Add holds details about calls to the Add method.
		Add []struct {
			// Db is the db argument value.
			Db kv.Tx
			// NewTxs is the newTxs argument value.
			NewTxs TxSlots
		}
		// AddNewGoodPeer holds details about calls to the AddNewGoodPeer method.
		AddNewGoodPeer []struct {
			// PeerID is the peerID argument value.
			PeerID PeerID
		}
		// GetRlp holds details about calls to the GetRlp method.
		GetRlp []struct {
			// Hash is the hash argument value.
			Hash []byte
		}
		// IdHashKnown holds details about calls to the IdHashKnown method.
		IdHashKnown []struct {
			// Hash is the hash argument value.
			Hash []byte
		}
		// OnNewBlock holds details about calls to the OnNewBlock method.
		OnNewBlock []struct {
			// Db is the db argument value.
			Db kv.Tx
			// StateChanges is the stateChanges argument value.
			StateChanges map[string]senderInfo
			// UnwindTxs is the unwindTxs argument value.
			UnwindTxs TxSlots
			// MinedTxs is the minedTxs argument value.
			MinedTxs TxSlots
			// ProtocolBaseFee is the protocolBaseFee argument value.
			ProtocolBaseFee uint64
			// PendingBaseFee is the pendingBaseFee argument value.
			PendingBaseFee uint64
			// BlockHeight is the blockHeight argument value.
			BlockHeight uint64
		}
	}
	lockAdd            sync.RWMutex
	lockAddNewGoodPeer sync.RWMutex
	lockGetRlp         sync.RWMutex
	lockIdHashKnown    sync.RWMutex
	lockOnNewBlock     sync.RWMutex
}

// Add calls AddFunc.
func (mock *PoolMock) Add(db kv.Tx, newTxs TxSlots) error {
	callInfo := struct {
		Db     kv.Tx
		NewTxs TxSlots
	}{
		Db:     db,
		NewTxs: newTxs,
	}
	mock.lockAdd.Lock()
	mock.calls.Add = append(mock.calls.Add, callInfo)
	mock.lockAdd.Unlock()
	if mock.AddFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.AddFunc(db, newTxs)
}

// AddCalls gets all the calls that were made to Add.
// Check the length with:
//     len(mockedPool.AddCalls())
func (mock *PoolMock) AddCalls() []struct {
	Db     kv.Tx
	NewTxs TxSlots
} {
	var calls []struct {
		Db     kv.Tx
		NewTxs TxSlots
	}
	mock.lockAdd.RLock()
	calls = mock.calls.Add
	mock.lockAdd.RUnlock()
	return calls
}

// AddNewGoodPeer calls AddNewGoodPeerFunc.
func (mock *PoolMock) AddNewGoodPeer(peerID PeerID) {
	callInfo := struct {
		PeerID PeerID
	}{
		PeerID: peerID,
	}
	mock.lockAddNewGoodPeer.Lock()
	mock.calls.AddNewGoodPeer = append(mock.calls.AddNewGoodPeer, callInfo)
	mock.lockAddNewGoodPeer.Unlock()
	if mock.AddNewGoodPeerFunc == nil {
		return
	}
	mock.AddNewGoodPeerFunc(peerID)
}

// AddNewGoodPeerCalls gets all the calls that were made to AddNewGoodPeer.
// Check the length with:
//     len(mockedPool.AddNewGoodPeerCalls())
func (mock *PoolMock) AddNewGoodPeerCalls() []struct {
	PeerID PeerID
} {
	var calls []struct {
		PeerID PeerID
	}
	mock.lockAddNewGoodPeer.RLock()
	calls = mock.calls.AddNewGoodPeer
	mock.lockAddNewGoodPeer.RUnlock()
	return calls
}

// GetRlp calls GetRlpFunc.
func (mock *PoolMock) GetRlp(hash []byte) []byte {
	callInfo := struct {
		Hash []byte
	}{
		Hash: hash,
	}
	mock.lockGetRlp.Lock()
	mock.calls.GetRlp = append(mock.calls.GetRlp, callInfo)
	mock.lockGetRlp.Unlock()
	if mock.GetRlpFunc == nil {
		var (
			bytesOut []byte
		)
		return bytesOut
	}
	return mock.GetRlpFunc(hash)
}

// GetRlpCalls gets all the calls that were made to GetRlp.
// Check the length with:
//     len(mockedPool.GetRlpCalls())
func (mock *PoolMock) GetRlpCalls() []struct {
	Hash []byte
} {
	var calls []struct {
		Hash []byte
	}
	mock.lockGetRlp.RLock()
	calls = mock.calls.GetRlp
	mock.lockGetRlp.RUnlock()
	return calls
}

// IdHashKnown calls IdHashKnownFunc.
func (mock *PoolMock) IdHashKnown(hash []byte) bool {
	callInfo := struct {
		Hash []byte
	}{
		Hash: hash,
	}
	mock.lockIdHashKnown.Lock()
	mock.calls.IdHashKnown = append(mock.calls.IdHashKnown, callInfo)
	mock.lockIdHashKnown.Unlock()
	if mock.IdHashKnownFunc == nil {
		var (
			bOut bool
		)
		return bOut
	}
	return mock.IdHashKnownFunc(hash)
}

// IdHashKnownCalls gets all the calls that were made to IdHashKnown.
// Check the length with:
//     len(mockedPool.IdHashKnownCalls())
func (mock *PoolMock) IdHashKnownCalls() []struct {
	Hash []byte
} {
	var calls []struct {
		Hash []byte
	}
	mock.lockIdHashKnown.RLock()
	calls = mock.calls.IdHashKnown
	mock.lockIdHashKnown.RUnlock()
	return calls
}

// OnNewBlock calls OnNewBlockFunc.
func (mock *PoolMock) OnNewBlock(db kv.Tx, stateChanges map[string]senderInfo, unwindTxs TxSlots, minedTxs TxSlots, protocolBaseFee uint64, pendingBaseFee uint64, blockHeight uint64) error {
	callInfo := struct {
		Db              kv.Tx
		StateChanges    map[string]senderInfo
		UnwindTxs       TxSlots
		MinedTxs        TxSlots
		ProtocolBaseFee uint64
		PendingBaseFee  uint64
		BlockHeight     uint64
	}{
		Db:              db,
		StateChanges:    stateChanges,
		UnwindTxs:       unwindTxs,
		MinedTxs:        minedTxs,
		ProtocolBaseFee: protocolBaseFee,
		PendingBaseFee:  pendingBaseFee,
		BlockHeight:     blockHeight,
	}
	mock.lockOnNewBlock.Lock()
	mock.calls.OnNewBlock = append(mock.calls.OnNewBlock, callInfo)
	mock.lockOnNewBlock.Unlock()
	if mock.OnNewBlockFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.OnNewBlockFunc(db, stateChanges, unwindTxs, minedTxs, protocolBaseFee, pendingBaseFee, blockHeight)
}

// OnNewBlockCalls gets all the calls that were made to OnNewBlock.
// Check the length with:
//     len(mockedPool.OnNewBlockCalls())
func (mock *PoolMock) OnNewBlockCalls() []struct {
	Db              kv.Tx
	StateChanges    map[string]senderInfo
	UnwindTxs       TxSlots
	MinedTxs        TxSlots
	ProtocolBaseFee uint64
	PendingBaseFee  uint64
	BlockHeight     uint64
} {
	var calls []struct {
		Db              kv.Tx
		StateChanges    map[string]senderInfo
		UnwindTxs       TxSlots
		MinedTxs        TxSlots
		ProtocolBaseFee uint64
		PendingBaseFee  uint64
		BlockHeight     uint64
	}
	mock.lockOnNewBlock.RLock()
	calls = mock.calls.OnNewBlock
	mock.lockOnNewBlock.RUnlock()
	return calls
}
