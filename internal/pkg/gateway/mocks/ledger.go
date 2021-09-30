// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/ledger"
)

type Ledger struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	GetBlockByNumberStub        func(uint64) (*common.Block, error)
	getBlockByNumberMutex       sync.RWMutex
	getBlockByNumberArgsForCall []struct {
		arg1 uint64
	}
	getBlockByNumberReturns struct {
		result1 *common.Block
		result2 error
	}
	getBlockByNumberReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	GetBlockchainInfoStub        func() (*common.BlockchainInfo, error)
	getBlockchainInfoMutex       sync.RWMutex
	getBlockchainInfoArgsForCall []struct {
	}
	getBlockchainInfoReturns struct {
		result1 *common.BlockchainInfo
		result2 error
	}
	getBlockchainInfoReturnsOnCall map[int]struct {
		result1 *common.BlockchainInfo
		result2 error
	}
	GetBlocksIteratorStub        func(uint64) (ledger.ResultsIterator, error)
	getBlocksIteratorMutex       sync.RWMutex
	getBlocksIteratorArgsForCall []struct {
		arg1 uint64
	}
	getBlocksIteratorReturns struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	getBlocksIteratorReturnsOnCall map[int]struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Ledger) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *Ledger) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *Ledger) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *Ledger) GetBlockByNumber(arg1 uint64) (*common.Block, error) {
	fake.getBlockByNumberMutex.Lock()
	ret, specificReturn := fake.getBlockByNumberReturnsOnCall[len(fake.getBlockByNumberArgsForCall)]
	fake.getBlockByNumberArgsForCall = append(fake.getBlockByNumberArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.GetBlockByNumberStub
	fakeReturns := fake.getBlockByNumberReturns
	fake.recordInvocation("GetBlockByNumber", []interface{}{arg1})
	fake.getBlockByNumberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) GetBlockByNumberCallCount() int {
	fake.getBlockByNumberMutex.RLock()
	defer fake.getBlockByNumberMutex.RUnlock()
	return len(fake.getBlockByNumberArgsForCall)
}

func (fake *Ledger) GetBlockByNumberCalls(stub func(uint64) (*common.Block, error)) {
	fake.getBlockByNumberMutex.Lock()
	defer fake.getBlockByNumberMutex.Unlock()
	fake.GetBlockByNumberStub = stub
}

func (fake *Ledger) GetBlockByNumberArgsForCall(i int) uint64 {
	fake.getBlockByNumberMutex.RLock()
	defer fake.getBlockByNumberMutex.RUnlock()
	argsForCall := fake.getBlockByNumberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) GetBlockByNumberReturns(result1 *common.Block, result2 error) {
	fake.getBlockByNumberMutex.Lock()
	defer fake.getBlockByNumberMutex.Unlock()
	fake.GetBlockByNumberStub = nil
	fake.getBlockByNumberReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockByNumberReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.getBlockByNumberMutex.Lock()
	defer fake.getBlockByNumberMutex.Unlock()
	fake.GetBlockByNumberStub = nil
	if fake.getBlockByNumberReturnsOnCall == nil {
		fake.getBlockByNumberReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.getBlockByNumberReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockchainInfo() (*common.BlockchainInfo, error) {
	fake.getBlockchainInfoMutex.Lock()
	ret, specificReturn := fake.getBlockchainInfoReturnsOnCall[len(fake.getBlockchainInfoArgsForCall)]
	fake.getBlockchainInfoArgsForCall = append(fake.getBlockchainInfoArgsForCall, struct {
	}{})
	stub := fake.GetBlockchainInfoStub
	fakeReturns := fake.getBlockchainInfoReturns
	fake.recordInvocation("GetBlockchainInfo", []interface{}{})
	fake.getBlockchainInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) GetBlockchainInfoCallCount() int {
	fake.getBlockchainInfoMutex.RLock()
	defer fake.getBlockchainInfoMutex.RUnlock()
	return len(fake.getBlockchainInfoArgsForCall)
}

func (fake *Ledger) GetBlockchainInfoCalls(stub func() (*common.BlockchainInfo, error)) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = stub
}

func (fake *Ledger) GetBlockchainInfoReturns(result1 *common.BlockchainInfo, result2 error) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = nil
	fake.getBlockchainInfoReturns = struct {
		result1 *common.BlockchainInfo
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlockchainInfoReturnsOnCall(i int, result1 *common.BlockchainInfo, result2 error) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = nil
	if fake.getBlockchainInfoReturnsOnCall == nil {
		fake.getBlockchainInfoReturnsOnCall = make(map[int]struct {
			result1 *common.BlockchainInfo
			result2 error
		})
	}
	fake.getBlockchainInfoReturnsOnCall[i] = struct {
		result1 *common.BlockchainInfo
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlocksIterator(arg1 uint64) (ledger.ResultsIterator, error) {
	fake.getBlocksIteratorMutex.Lock()
	ret, specificReturn := fake.getBlocksIteratorReturnsOnCall[len(fake.getBlocksIteratorArgsForCall)]
	fake.getBlocksIteratorArgsForCall = append(fake.getBlocksIteratorArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.GetBlocksIteratorStub
	fakeReturns := fake.getBlocksIteratorReturns
	fake.recordInvocation("GetBlocksIterator", []interface{}{arg1})
	fake.getBlocksIteratorMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Ledger) GetBlocksIteratorCallCount() int {
	fake.getBlocksIteratorMutex.RLock()
	defer fake.getBlocksIteratorMutex.RUnlock()
	return len(fake.getBlocksIteratorArgsForCall)
}

func (fake *Ledger) GetBlocksIteratorCalls(stub func(uint64) (ledger.ResultsIterator, error)) {
	fake.getBlocksIteratorMutex.Lock()
	defer fake.getBlocksIteratorMutex.Unlock()
	fake.GetBlocksIteratorStub = stub
}

func (fake *Ledger) GetBlocksIteratorArgsForCall(i int) uint64 {
	fake.getBlocksIteratorMutex.RLock()
	defer fake.getBlocksIteratorMutex.RUnlock()
	argsForCall := fake.getBlocksIteratorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Ledger) GetBlocksIteratorReturns(result1 ledger.ResultsIterator, result2 error) {
	fake.getBlocksIteratorMutex.Lock()
	defer fake.getBlocksIteratorMutex.Unlock()
	fake.GetBlocksIteratorStub = nil
	fake.getBlocksIteratorReturns = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *Ledger) GetBlocksIteratorReturnsOnCall(i int, result1 ledger.ResultsIterator, result2 error) {
	fake.getBlocksIteratorMutex.Lock()
	defer fake.getBlocksIteratorMutex.Unlock()
	fake.GetBlocksIteratorStub = nil
	if fake.getBlocksIteratorReturnsOnCall == nil {
		fake.getBlocksIteratorReturnsOnCall = make(map[int]struct {
			result1 ledger.ResultsIterator
			result2 error
		})
	}
	fake.getBlocksIteratorReturnsOnCall[i] = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *Ledger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.getBlockByNumberMutex.RLock()
	defer fake.getBlockByNumberMutex.RUnlock()
	fake.getBlockchainInfoMutex.RLock()
	defer fake.getBlockchainInfoMutex.RUnlock()
	fake.getBlocksIteratorMutex.RLock()
	defer fake.getBlocksIteratorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Ledger) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
