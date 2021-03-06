// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/platform"
)

type FakeWindowsDiskLinker struct {
	IsLinkedStub        func(location string) (target string, err error)
	isLinkedMutex       sync.RWMutex
	isLinkedArgsForCall []struct {
		location string
	}
	isLinkedReturns struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWindowsDiskLinker) IsLinked(location string) (target string, err error) {
	fake.isLinkedMutex.Lock()
	fake.isLinkedArgsForCall = append(fake.isLinkedArgsForCall, struct {
		location string
	}{location})
	fake.recordInvocation("IsLinked", []interface{}{location})
	fake.isLinkedMutex.Unlock()
	if fake.IsLinkedStub != nil {
		return fake.IsLinkedStub(location)
	} else {
		return fake.isLinkedReturns.result1, fake.isLinkedReturns.result2
	}
}

func (fake *FakeWindowsDiskLinker) IsLinkedCallCount() int {
	fake.isLinkedMutex.RLock()
	defer fake.isLinkedMutex.RUnlock()
	return len(fake.isLinkedArgsForCall)
}

func (fake *FakeWindowsDiskLinker) IsLinkedArgsForCall(i int) string {
	fake.isLinkedMutex.RLock()
	defer fake.isLinkedMutex.RUnlock()
	return fake.isLinkedArgsForCall[i].location
}

func (fake *FakeWindowsDiskLinker) IsLinkedReturns(result1 string, result2 error) {
	fake.IsLinkedStub = nil
	fake.isLinkedReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeWindowsDiskLinker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isLinkedMutex.RLock()
	defer fake.isLinkedMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWindowsDiskLinker) recordInvocation(key string, args []interface{}) {
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

var _ platform.WindowsDiskLinker = new(FakeWindowsDiskLinker)
