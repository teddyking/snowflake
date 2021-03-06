// Code generated by counterfeiter. DO NOT EDIT.
package reporterfakes

import (
	"context"
	"sync"

	"github.com/teddyking/snowflake/api"
	"github.com/teddyking/snowflake/reporter"
	"google.golang.org/grpc"
)

type FakeReporterService struct {
	CreateStub        func(ctx context.Context, in *api.ReporterCreateReq, opts ...grpc.CallOption) (*api.ReporterCreateRes, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		ctx  context.Context
		in   *api.ReporterCreateReq
		opts []grpc.CallOption
	}
	createReturns struct {
		result1 *api.ReporterCreateRes
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *api.ReporterCreateRes
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReporterService) Create(ctx context.Context, in *api.ReporterCreateReq, opts ...grpc.CallOption) (*api.ReporterCreateRes, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		ctx  context.Context
		in   *api.ReporterCreateReq
		opts []grpc.CallOption
	}{ctx, in, opts})
	fake.recordInvocation("Create", []interface{}{ctx, in, opts})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(ctx, in, opts...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeReporterService) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeReporterService) CreateArgsForCall(i int) (context.Context, *api.ReporterCreateReq, []grpc.CallOption) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].ctx, fake.createArgsForCall[i].in, fake.createArgsForCall[i].opts
}

func (fake *FakeReporterService) CreateReturns(result1 *api.ReporterCreateRes, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *api.ReporterCreateRes
		result2 error
	}{result1, result2}
}

func (fake *FakeReporterService) CreateReturnsOnCall(i int, result1 *api.ReporterCreateRes, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *api.ReporterCreateRes
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *api.ReporterCreateRes
		result2 error
	}{result1, result2}
}

func (fake *FakeReporterService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReporterService) recordInvocation(key string, args []interface{}) {
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

var _ reporter.ReporterService = new(FakeReporterService)
