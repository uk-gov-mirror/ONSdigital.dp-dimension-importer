// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/dataset"
	"github.com/ONSdigital/dp-dimension-importer/client"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"sync"
)

var (
	lockIClientMockChecker                      sync.RWMutex
	lockIClientMockGetInstance                  sync.RWMutex
	lockIClientMockGetInstanceDimensions        sync.RWMutex
	lockIClientMockPatchInstanceDimensionOption sync.RWMutex
)

// Ensure, that IClientMock does implement client.IClient.
// If this is not the case, regenerate this file with moq.
var _ client.IClient = &IClientMock{}

// IClientMock is a mock implementation of client.IClient.
//
//     func TestSomethingThatUsesIClient(t *testing.T) {
//
//         // make and configure a mocked client.IClient
//         mockedIClient := &IClientMock{
//             CheckerFunc: func(ctx context.Context, state *healthcheck.CheckState) error {
// 	               panic("mock out the Checker method")
//             },
//             GetInstanceFunc: func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, instanceID string) (dataset.Instance, error) {
// 	               panic("mock out the GetInstance method")
//             },
//             GetInstanceDimensionsFunc: func(ctx context.Context, serviceAuthToken string, instanceID string) (dataset.Dimensions, error) {
// 	               panic("mock out the GetInstanceDimensions method")
//             },
//             PatchInstanceDimensionOptionFunc: func(ctx context.Context, serviceAuthToken string, instanceID string, dimensionID string, optionID string, nodeID string, order *int) error {
// 	               panic("mock out the PatchInstanceDimensionOption method")
//             },
//         }
//
//         // use mockedIClient in code that requires client.IClient
//         // and then make assertions.
//
//     }
type IClientMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, state *healthcheck.CheckState) error

	// GetInstanceFunc mocks the GetInstance method.
	GetInstanceFunc func(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, instanceID string) (dataset.Instance, error)

	// GetInstanceDimensionsFunc mocks the GetInstanceDimensions method.
	GetInstanceDimensionsFunc func(ctx context.Context, serviceAuthToken string, instanceID string) (dataset.Dimensions, error)

	// PatchInstanceDimensionOptionFunc mocks the PatchInstanceDimensionOption method.
	PatchInstanceDimensionOptionFunc func(ctx context.Context, serviceAuthToken string, instanceID string, dimensionID string, optionID string, nodeID string, order *int) error

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// State is the state argument value.
			State *healthcheck.CheckState
		}
		// GetInstance holds details about calls to the GetInstance method.
		GetInstance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserAuthToken is the userAuthToken argument value.
			UserAuthToken string
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// CollectionID is the collectionID argument value.
			CollectionID string
			// InstanceID is the instanceID argument value.
			InstanceID string
		}
		// GetInstanceDimensions holds details about calls to the GetInstanceDimensions method.
		GetInstanceDimensions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// InstanceID is the instanceID argument value.
			InstanceID string
		}
		// PatchInstanceDimensionOption holds details about calls to the PatchInstanceDimensionOption method.
		PatchInstanceDimensionOption []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// InstanceID is the instanceID argument value.
			InstanceID string
			// DimensionID is the dimensionID argument value.
			DimensionID string
			// OptionID is the optionID argument value.
			OptionID string
			// NodeID is the nodeID argument value.
			NodeID string
			// Order is the order argument value.
			Order *int
		}
	}
}

// Checker calls CheckerFunc.
func (mock *IClientMock) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("IClientMock.CheckerFunc: method is nil but IClient.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		State *healthcheck.CheckState
	}{
		Ctx:   ctx,
		State: state,
	}
	lockIClientMockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	lockIClientMockChecker.Unlock()
	return mock.CheckerFunc(ctx, state)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedIClient.CheckerCalls())
func (mock *IClientMock) CheckerCalls() []struct {
	Ctx   context.Context
	State *healthcheck.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		State *healthcheck.CheckState
	}
	lockIClientMockChecker.RLock()
	calls = mock.calls.Checker
	lockIClientMockChecker.RUnlock()
	return calls
}

// GetInstance calls GetInstanceFunc.
func (mock *IClientMock) GetInstance(ctx context.Context, userAuthToken string, serviceAuthToken string, collectionID string, instanceID string) (dataset.Instance, error) {
	if mock.GetInstanceFunc == nil {
		panic("IClientMock.GetInstanceFunc: method is nil but IClient.GetInstance was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		InstanceID       string
	}{
		Ctx:              ctx,
		UserAuthToken:    userAuthToken,
		ServiceAuthToken: serviceAuthToken,
		CollectionID:     collectionID,
		InstanceID:       instanceID,
	}
	lockIClientMockGetInstance.Lock()
	mock.calls.GetInstance = append(mock.calls.GetInstance, callInfo)
	lockIClientMockGetInstance.Unlock()
	return mock.GetInstanceFunc(ctx, userAuthToken, serviceAuthToken, collectionID, instanceID)
}

// GetInstanceCalls gets all the calls that were made to GetInstance.
// Check the length with:
//     len(mockedIClient.GetInstanceCalls())
func (mock *IClientMock) GetInstanceCalls() []struct {
	Ctx              context.Context
	UserAuthToken    string
	ServiceAuthToken string
	CollectionID     string
	InstanceID       string
} {
	var calls []struct {
		Ctx              context.Context
		UserAuthToken    string
		ServiceAuthToken string
		CollectionID     string
		InstanceID       string
	}
	lockIClientMockGetInstance.RLock()
	calls = mock.calls.GetInstance
	lockIClientMockGetInstance.RUnlock()
	return calls
}

// GetInstanceDimensions calls GetInstanceDimensionsFunc.
func (mock *IClientMock) GetInstanceDimensions(ctx context.Context, serviceAuthToken string, instanceID string) (dataset.Dimensions, error) {
	if mock.GetInstanceDimensionsFunc == nil {
		panic("IClientMock.GetInstanceDimensionsFunc: method is nil but IClient.GetInstanceDimensions was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
	}{
		Ctx:              ctx,
		ServiceAuthToken: serviceAuthToken,
		InstanceID:       instanceID,
	}
	lockIClientMockGetInstanceDimensions.Lock()
	mock.calls.GetInstanceDimensions = append(mock.calls.GetInstanceDimensions, callInfo)
	lockIClientMockGetInstanceDimensions.Unlock()
	return mock.GetInstanceDimensionsFunc(ctx, serviceAuthToken, instanceID)
}

// GetInstanceDimensionsCalls gets all the calls that were made to GetInstanceDimensions.
// Check the length with:
//     len(mockedIClient.GetInstanceDimensionsCalls())
func (mock *IClientMock) GetInstanceDimensionsCalls() []struct {
	Ctx              context.Context
	ServiceAuthToken string
	InstanceID       string
} {
	var calls []struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
	}
	lockIClientMockGetInstanceDimensions.RLock()
	calls = mock.calls.GetInstanceDimensions
	lockIClientMockGetInstanceDimensions.RUnlock()
	return calls
}

// PatchInstanceDimensionOption calls PatchInstanceDimensionOptionFunc.
func (mock *IClientMock) PatchInstanceDimensionOption(ctx context.Context, serviceAuthToken string, instanceID string, dimensionID string, optionID string, nodeID string, order *int) error {
	if mock.PatchInstanceDimensionOptionFunc == nil {
		panic("IClientMock.PatchInstanceDimensionOptionFunc: method is nil but IClient.PatchInstanceDimensionOption was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		DimensionID      string
		OptionID         string
		NodeID           string
		Order            *int
	}{
		Ctx:              ctx,
		ServiceAuthToken: serviceAuthToken,
		InstanceID:       instanceID,
		DimensionID:      dimensionID,
		OptionID:         optionID,
		NodeID:           nodeID,
		Order:            order,
	}
	lockIClientMockPatchInstanceDimensionOption.Lock()
	mock.calls.PatchInstanceDimensionOption = append(mock.calls.PatchInstanceDimensionOption, callInfo)
	lockIClientMockPatchInstanceDimensionOption.Unlock()
	return mock.PatchInstanceDimensionOptionFunc(ctx, serviceAuthToken, instanceID, dimensionID, optionID, nodeID, order)
}

// PatchInstanceDimensionOptionCalls gets all the calls that were made to PatchInstanceDimensionOption.
// Check the length with:
//     len(mockedIClient.PatchInstanceDimensionOptionCalls())
func (mock *IClientMock) PatchInstanceDimensionOptionCalls() []struct {
	Ctx              context.Context
	ServiceAuthToken string
	InstanceID       string
	DimensionID      string
	OptionID         string
	NodeID           string
	Order            *int
} {
	var calls []struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		DimensionID      string
		OptionID         string
		NodeID           string
		Order            *int
	}
	lockIClientMockPatchInstanceDimensionOption.RLock()
	calls = mock.calls.PatchInstanceDimensionOption
	lockIClientMockPatchInstanceDimensionOption.RUnlock()
	return calls
}
