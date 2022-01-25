// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package handlers

import (
	"context"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"sync"
)

// Ensure, that RendererClientMock does implement RendererClient.
// If this is not the case, regenerate this file with moq.
var _ RendererClient = &RendererClientMock{}

// RendererClientMock is a mock implementation of RendererClient.
//
// 	func TestSomethingThatUsesRendererClient(t *testing.T) {
//
// 		// make and configure a mocked RendererClient
// 		mockedRendererClient := &RendererClientMock{
// 			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
// 				panic("mock out the Checker method")
// 			},
// 			DoFunc: func(s string, bytes []byte) ([]byte, error) {
// 				panic("mock out the Do method")
// 			},
// 		}
//
// 		// use mockedRendererClient in code that requires RendererClient
// 		// and then make assertions.
//
// 	}
type RendererClientMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, check *health.CheckState) error

	// DoFunc mocks the Do method.
	DoFunc func(s string, bytes []byte) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Check is the check argument value.
			Check *health.CheckState
		}
		// Do holds details about calls to the Do method.
		Do []struct {
			// S is the s argument value.
			S string
			// Bytes is the bytes argument value.
			Bytes []byte
		}
	}
	lockChecker sync.RWMutex
	lockDo      sync.RWMutex
}

// Checker calls CheckerFunc.
func (mock *RendererClientMock) Checker(ctx context.Context, check *health.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("RendererClientMock.CheckerFunc: method is nil but RendererClient.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Check *health.CheckState
	}{
		Ctx:   ctx,
		Check: check,
	}
	mock.lockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	mock.lockChecker.Unlock()
	return mock.CheckerFunc(ctx, check)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedRendererClient.CheckerCalls())
func (mock *RendererClientMock) CheckerCalls() []struct {
	Ctx   context.Context
	Check *health.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		Check *health.CheckState
	}
	mock.lockChecker.RLock()
	calls = mock.calls.Checker
	mock.lockChecker.RUnlock()
	return calls
}

// Do calls DoFunc.
func (mock *RendererClientMock) Do(s string, bytes []byte) ([]byte, error) {
	if mock.DoFunc == nil {
		panic("RendererClientMock.DoFunc: method is nil but RendererClient.Do was just called")
	}
	callInfo := struct {
		S     string
		Bytes []byte
	}{
		S:     s,
		Bytes: bytes,
	}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	return mock.DoFunc(s, bytes)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedRendererClient.DoCalls())
func (mock *RendererClientMock) DoCalls() []struct {
	S     string
	Bytes []byte
} {
	var calls []struct {
		S     string
		Bytes []byte
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}
