// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/cf-mysql-benchmark-app/sysbench_client/os_client"
)

type FakeOsClient struct {
	CombinedOutputStub        func(string, ...string) ([]byte, error)
	combinedOutputMutex       sync.RWMutex
	combinedOutputArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	combinedOutputReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeOsClient) CombinedOutput(arg1 string, arg2 ...string) ([]byte, error) {
	fake.combinedOutputMutex.Lock()
	fake.combinedOutputArgsForCall = append(fake.combinedOutputArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.combinedOutputMutex.Unlock()
	if fake.CombinedOutputStub != nil {
		return fake.CombinedOutputStub(arg1, arg2...)
	} else {
		return fake.combinedOutputReturns.result1, fake.combinedOutputReturns.result2
	}
}

func (fake *FakeOsClient) CombinedOutputCallCount() int {
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	return len(fake.combinedOutputArgsForCall)
}

func (fake *FakeOsClient) CombinedOutputArgsForCall(i int) (string, []string) {
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	return fake.combinedOutputArgsForCall[i].arg1, fake.combinedOutputArgsForCall[i].arg2
}

func (fake *FakeOsClient) CombinedOutputReturns(result1 []byte, result2 error) {
	fake.CombinedOutputStub = nil
	fake.combinedOutputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ os_client.OsClient = new(FakeOsClient)
