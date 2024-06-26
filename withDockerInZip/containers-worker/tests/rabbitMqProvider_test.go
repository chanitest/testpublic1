// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package tests

import (
	"context"
	"github.com/Checkmarx-Containers/containers-worker/internal/ContainersEngine"
	"github.com/Checkmarx-Containers/containers-worker/internal/messages"
	"sync"
)

// Ensure, that RabbitMqIntMock does implement RabbitMqInt.
// If this is not the case, regenerate this file with moq.
var _ messages.RabbitMqInt = &RabbitMqIntMock{}

// RabbitMqIntMock is a mock implementation of RabbitMqInt.
//
//	func TestSomethingThatUsesRabbitMqInt(t *testing.T) {
//
//		// make and configure a mocked RabbitMqInt
//		mockedRabbitMqInt := &RabbitMqIntMock{
//			SendNewScanRequestFunc: func(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error {
//				panic("mock out the SendNewScanRequest method")
//			},
//		}
//
//		// use mockedRabbitMqInt in code that requires RabbitMqInt
//		// and then make assertions.
//
//	}
type RabbitMqIntMock struct {
	// SendNewScanRequestFunc mocks the SendNewScanRequest method.
	SendNewScanRequestFunc func(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error

	// calls tracks calls to the methods.
	calls struct {
		// SendNewScanRequest holds details about calls to the SendNewScanRequest method.
		SendNewScanRequest []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// NewScanRequest is the newScanRequest argument value.
			NewScanRequest *ContainersEngine.NewScanRequest
		}
	}
	lockSendNewScanRequest sync.RWMutex
}

// SendNewScanRequest calls SendNewScanRequestFunc.
func (mock *RabbitMqIntMock) SendNewScanRequest(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error {
	if mock.SendNewScanRequestFunc == nil {
		panic("RabbitMqIntMock.SendNewScanRequestFunc: method is nil but RabbitMqInt.SendNewScanRequest was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		NewScanRequest *ContainersEngine.NewScanRequest
	}{
		Ctx:            ctx,
		NewScanRequest: newScanRequest,
	}
	mock.lockSendNewScanRequest.Lock()
	mock.calls.SendNewScanRequest = append(mock.calls.SendNewScanRequest, callInfo)
	mock.lockSendNewScanRequest.Unlock()
	return mock.SendNewScanRequestFunc(ctx, newScanRequest)
}

// SendNewScanRequestCalls gets all the calls that were made to SendNewScanRequest.
// Check the length with:
//
//	len(mockedRabbitMqInt.SendNewScanRequestCalls())
func (mock *RabbitMqIntMock) SendNewScanRequestCalls() []struct {
	Ctx            context.Context
	NewScanRequest *ContainersEngine.NewScanRequest
} {
	var calls []struct {
		Ctx            context.Context
		NewScanRequest *ContainersEngine.NewScanRequest
	}
	mock.lockSendNewScanRequest.RLock()
	calls = mock.calls.SendNewScanRequest
	mock.lockSendNewScanRequest.RUnlock()
	return calls
}
