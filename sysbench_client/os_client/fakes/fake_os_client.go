// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/cf-mysql-benchmark-app/sysbench_client/os_client"
)

type FakeOsClient struct {
	ExecStub        func(...string) error
	execMutex       sync.RWMutex
	execArgsForCall []struct {
		arg1 []string
	}
	execReturns struct {
		result1 error
	}
}

func (fake *FakeOsClient) Exec(arg1 ...string) error {
	fake.execMutex.Lock()
	fake.execArgsForCall = append(fake.execArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.execMutex.Unlock()
	if fake.ExecStub != nil {
		return fake.ExecStub(arg1...)
	} else {
		return fake.execReturns.result1
	}
}

func (fake *FakeOsClient) ExecCallCount() int {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return len(fake.execArgsForCall)
}

func (fake *FakeOsClient) ExecArgsForCall(i int) []string {
	fake.execMutex.RLock()
	defer fake.execMutex.RUnlock()
	return fake.execArgsForCall[i].arg1
}

func (fake *FakeOsClient) ExecReturns(result1 error) {
	fake.ExecStub = nil
	fake.execReturns = struct {
		result1 error
	}{result1}
}

var _ os_client.OsClient = new(FakeOsClient)