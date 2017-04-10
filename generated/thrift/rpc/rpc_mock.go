// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3db/generated/thrift/rpc/tchan-rpc.go

package rpc

import (
	gomock "github.com/golang/mock/gomock"
	thrift "github.com/uber/tchannel-go/thrift"
)

// Mock of TChanCluster interface
type MockTChanCluster struct {
	ctrl     *gomock.Controller
	recorder *_MockTChanClusterRecorder
}

// Recorder for MockTChanCluster (not exported)
type _MockTChanClusterRecorder struct {
	mock *MockTChanCluster
}

func NewMockTChanCluster(ctrl *gomock.Controller) *MockTChanCluster {
	mock := &MockTChanCluster{ctrl: ctrl}
	mock.recorder = &_MockTChanClusterRecorder{mock}
	return mock
}

func (_m *MockTChanCluster) EXPECT() *_MockTChanClusterRecorder {
	return _m.recorder
}

func (_m *MockTChanCluster) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	ret := _m.ctrl.Call(_m, "Fetch", ctx, req)
	ret0, _ := ret[0].(*FetchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanClusterRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0, arg1)
}

func (_m *MockTChanCluster) Health(ctx thrift.Context) (*HealthResult_, error) {
	ret := _m.ctrl.Call(_m, "Health", ctx)
	ret0, _ := ret[0].(*HealthResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanClusterRecorder) Health(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Health", arg0)
}

func (_m *MockTChanCluster) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	ret := _m.ctrl.Call(_m, "Truncate", ctx, req)
	ret0, _ := ret[0].(*TruncateResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanClusterRecorder) Truncate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Truncate", arg0, arg1)
}

func (_m *MockTChanCluster) Write(ctx thrift.Context, req *WriteRequest) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanClusterRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1)
}

// Mock of TChanNode interface
type MockTChanNode struct {
	ctrl     *gomock.Controller
	recorder *_MockTChanNodeRecorder
}

// Recorder for MockTChanNode (not exported)
type _MockTChanNodeRecorder struct {
	mock *MockTChanNode
}

func NewMockTChanNode(ctrl *gomock.Controller) *MockTChanNode {
	mock := &MockTChanNode{ctrl: ctrl}
	mock.recorder = &_MockTChanNodeRecorder{mock}
	return mock
}

func (_m *MockTChanNode) EXPECT() *_MockTChanNodeRecorder {
	return _m.recorder
}

func (_m *MockTChanNode) Fetch(ctx thrift.Context, req *FetchRequest) (*FetchResult_, error) {
	ret := _m.ctrl.Call(_m, "Fetch", ctx, req)
	ret0, _ := ret[0].(*FetchResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) Fetch(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fetch", arg0, arg1)
}

func (_m *MockTChanNode) FetchBatchRaw(ctx thrift.Context, req *FetchBatchRawRequest) (*FetchBatchRawResult_, error) {
	ret := _m.ctrl.Call(_m, "FetchBatchRaw", ctx, req)
	ret0, _ := ret[0].(*FetchBatchRawResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) FetchBatchRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBatchRaw", arg0, arg1)
}

func (_m *MockTChanNode) FetchBlocksMetadataRaw(ctx thrift.Context, req *FetchBlocksMetadataRawRequest) (*FetchBlocksMetadataRawResult_, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocksMetadataRaw", ctx, req)
	ret0, _ := ret[0].(*FetchBlocksMetadataRawResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) FetchBlocksMetadataRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksMetadataRaw", arg0, arg1)
}

func (_m *MockTChanNode) FetchBlocksRaw(ctx thrift.Context, req *FetchBlocksRawRequest) (*FetchBlocksRawResult_, error) {
	ret := _m.ctrl.Call(_m, "FetchBlocksRaw", ctx, req)
	ret0, _ := ret[0].(*FetchBlocksRawResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) FetchBlocksRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FetchBlocksRaw", arg0, arg1)
}

func (_m *MockTChanNode) GetPersistRateLimit(ctx thrift.Context) (*NodePersistRateLimitResult_, error) {
	ret := _m.ctrl.Call(_m, "GetPersistRateLimit", ctx)
	ret0, _ := ret[0].(*NodePersistRateLimitResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) GetPersistRateLimit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPersistRateLimit", arg0)
}

func (_m *MockTChanNode) GetWriteNewSeriesAsync(ctx thrift.Context) (*NodeWriteNewSeriesAsyncResult_, error) {
	ret := _m.ctrl.Call(_m, "GetWriteNewSeriesAsync", ctx)
	ret0, _ := ret[0].(*NodeWriteNewSeriesAsyncResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) GetWriteNewSeriesAsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetWriteNewSeriesAsync", arg0)
}

func (_m *MockTChanNode) Health(ctx thrift.Context) (*NodeHealthResult_, error) {
	ret := _m.ctrl.Call(_m, "Health", ctx)
	ret0, _ := ret[0].(*NodeHealthResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) Health(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Health", arg0)
}

func (_m *MockTChanNode) Repair(ctx thrift.Context) error {
	ret := _m.ctrl.Call(_m, "Repair", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanNodeRecorder) Repair(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Repair", arg0)
}

func (_m *MockTChanNode) SetPersistRateLimit(ctx thrift.Context, req *NodeSetPersistRateLimitRequest) (*NodePersistRateLimitResult_, error) {
	ret := _m.ctrl.Call(_m, "SetPersistRateLimit", ctx, req)
	ret0, _ := ret[0].(*NodePersistRateLimitResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) SetPersistRateLimit(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetPersistRateLimit", arg0, arg1)
}

func (_m *MockTChanNode) SetWriteNewSeriesAsync(ctx thrift.Context, req *NodeSetWriteNewSeriesAsyncRequest) (*NodeWriteNewSeriesAsyncResult_, error) {
	ret := _m.ctrl.Call(_m, "SetWriteNewSeriesAsync", ctx, req)
	ret0, _ := ret[0].(*NodeWriteNewSeriesAsyncResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) SetWriteNewSeriesAsync(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetWriteNewSeriesAsync", arg0, arg1)
}

func (_m *MockTChanNode) Truncate(ctx thrift.Context, req *TruncateRequest) (*TruncateResult_, error) {
	ret := _m.ctrl.Call(_m, "Truncate", ctx, req)
	ret0, _ := ret[0].(*TruncateResult_)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTChanNodeRecorder) Truncate(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Truncate", arg0, arg1)
}

func (_m *MockTChanNode) Write(ctx thrift.Context, req *WriteRequest) error {
	ret := _m.ctrl.Call(_m, "Write", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanNodeRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0, arg1)
}

func (_m *MockTChanNode) WriteBatchRaw(ctx thrift.Context, req *WriteBatchRawRequest) error {
	ret := _m.ctrl.Call(_m, "WriteBatchRaw", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTChanNodeRecorder) WriteBatchRaw(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteBatchRaw", arg0, arg1)
}
