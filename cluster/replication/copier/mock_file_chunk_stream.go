//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by mockery v2.53.2. DO NOT EDIT.

package copier

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	protocol "github.com/weaviate/weaviate/grpc/generated/protocol/v1"
)

// MockFileChunkStream is an autogenerated mock type for the FileChunkStream type
type MockFileChunkStream struct {
	mock.Mock
}

type MockFileChunkStream_Expecter struct {
	mock *mock.Mock
}

func (_m *MockFileChunkStream) EXPECT() *MockFileChunkStream_Expecter {
	return &MockFileChunkStream_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with no fields
func (_m *MockFileChunkStream) CloseSend() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CloseSend")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileChunkStream_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockFileChunkStream_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockFileChunkStream_Expecter) CloseSend() *MockFileChunkStream_CloseSend_Call {
	return &MockFileChunkStream_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *MockFileChunkStream_CloseSend_Call) Run(run func()) *MockFileChunkStream_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileChunkStream_CloseSend_Call) Return(_a0 error) *MockFileChunkStream_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileChunkStream_CloseSend_Call) RunAndReturn(run func() error) *MockFileChunkStream_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with no fields
func (_m *MockFileChunkStream) Context() context.Context {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Context")
	}

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockFileChunkStream_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockFileChunkStream_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockFileChunkStream_Expecter) Context() *MockFileChunkStream_Context_Call {
	return &MockFileChunkStream_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockFileChunkStream_Context_Call) Run(run func()) *MockFileChunkStream_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileChunkStream_Context_Call) Return(_a0 context.Context) *MockFileChunkStream_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileChunkStream_Context_Call) RunAndReturn(run func() context.Context) *MockFileChunkStream_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with no fields
func (_m *MockFileChunkStream) Header() (metadata.MD, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Header")
	}

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileChunkStream_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockFileChunkStream_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockFileChunkStream_Expecter) Header() *MockFileChunkStream_Header_Call {
	return &MockFileChunkStream_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockFileChunkStream_Header_Call) Run(run func()) *MockFileChunkStream_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileChunkStream_Header_Call) Return(_a0 metadata.MD, _a1 error) *MockFileChunkStream_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileChunkStream_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *MockFileChunkStream_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with no fields
func (_m *MockFileChunkStream) Recv() (*protocol.FileChunk, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Recv")
	}

	var r0 *protocol.FileChunk
	var r1 error
	if rf, ok := ret.Get(0).(func() (*protocol.FileChunk, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *protocol.FileChunk); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protocol.FileChunk)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockFileChunkStream_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockFileChunkStream_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockFileChunkStream_Expecter) Recv() *MockFileChunkStream_Recv_Call {
	return &MockFileChunkStream_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockFileChunkStream_Recv_Call) Run(run func()) *MockFileChunkStream_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileChunkStream_Recv_Call) Return(_a0 *protocol.FileChunk, _a1 error) *MockFileChunkStream_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockFileChunkStream_Recv_Call) RunAndReturn(run func() (*protocol.FileChunk, error)) *MockFileChunkStream_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockFileChunkStream) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	if len(ret) == 0 {
		panic("no return value specified for RecvMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileChunkStream_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockFileChunkStream_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *MockFileChunkStream_Expecter) RecvMsg(m interface{}) *MockFileChunkStream_RecvMsg_Call {
	return &MockFileChunkStream_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockFileChunkStream_RecvMsg_Call) Run(run func(m interface{})) *MockFileChunkStream_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockFileChunkStream_RecvMsg_Call) Return(_a0 error) *MockFileChunkStream_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileChunkStream_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *MockFileChunkStream_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockFileChunkStream) Send(_a0 *protocol.GetFileRequest) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*protocol.GetFileRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileChunkStream_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockFileChunkStream_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - _a0 *protocol.GetFileRequest
func (_e *MockFileChunkStream_Expecter) Send(_a0 interface{}) *MockFileChunkStream_Send_Call {
	return &MockFileChunkStream_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockFileChunkStream_Send_Call) Run(run func(_a0 *protocol.GetFileRequest)) *MockFileChunkStream_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*protocol.GetFileRequest))
	})
	return _c
}

func (_c *MockFileChunkStream_Send_Call) Return(_a0 error) *MockFileChunkStream_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileChunkStream_Send_Call) RunAndReturn(run func(*protocol.GetFileRequest) error) *MockFileChunkStream_Send_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: _a0
func (_m *MockFileChunkStream) SendMsg(_a0 interface{}) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for SendMsg")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockFileChunkStream_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockFileChunkStream_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - _a0 interface{}
func (_e *MockFileChunkStream_Expecter) SendMsg(_a0 interface{}) *MockFileChunkStream_SendMsg_Call {
	return &MockFileChunkStream_SendMsg_Call{Call: _e.mock.On("SendMsg", _a0)}
}

func (_c *MockFileChunkStream_SendMsg_Call) Run(run func(_a0 interface{})) *MockFileChunkStream_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockFileChunkStream_SendMsg_Call) Return(_a0 error) *MockFileChunkStream_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileChunkStream_SendMsg_Call) RunAndReturn(run func(interface{}) error) *MockFileChunkStream_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with no fields
func (_m *MockFileChunkStream) Trailer() metadata.MD {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Trailer")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// MockFileChunkStream_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockFileChunkStream_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockFileChunkStream_Expecter) Trailer() *MockFileChunkStream_Trailer_Call {
	return &MockFileChunkStream_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *MockFileChunkStream_Trailer_Call) Run(run func()) *MockFileChunkStream_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockFileChunkStream_Trailer_Call) Return(_a0 metadata.MD) *MockFileChunkStream_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockFileChunkStream_Trailer_Call) RunAndReturn(run func() metadata.MD) *MockFileChunkStream_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockFileChunkStream creates a new instance of MockFileChunkStream. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFileChunkStream(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFileChunkStream {
	mock := &MockFileChunkStream{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
