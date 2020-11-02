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
//             PutInstanceDimensionOptionNodeIDFunc: func(ctx context.Context, serviceAuthToken string, instanceID string, dimensionID string, optionID string, nodeID string) error {
// 	               panic("mock out the PutInstanceDimensionOptionNodeID method")
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

	// PutInstanceDimensionOptionNodeIDFunc mocks the PutInstanceDimensionOptionNodeID method.
	PutInstanceDimensionOptionNodeIDFunc func(ctx context.Context, serviceAuthToken string, instanceID string, dimensionID string, optionID string, nodeID string) error

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
		// PutInstanceDimensionOptionNodeID holds details about calls to the PutInstanceDimensionOptionNodeID method.
		PutInstanceDimensionOptionNodeID []struct {
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
		}
	}
	lockChecker                          sync.RWMutex
	lockGetInstance                      sync.RWMutex
	lockGetInstanceDimensions            sync.RWMutex
	lockPutInstanceDimensionOptionNodeID sync.RWMutex
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
	mock.lockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	mock.lockChecker.Unlock()
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
	mock.lockChecker.RLock()
	calls = mock.calls.Checker
	mock.lockChecker.RUnlock()
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
	mock.lockGetInstance.Lock()
	mock.calls.GetInstance = append(mock.calls.GetInstance, callInfo)
	mock.lockGetInstance.Unlock()
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
	mock.lockGetInstance.RLock()
	calls = mock.calls.GetInstance
	mock.lockGetInstance.RUnlock()
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
	mock.lockGetInstanceDimensions.Lock()
	mock.calls.GetInstanceDimensions = append(mock.calls.GetInstanceDimensions, callInfo)
	mock.lockGetInstanceDimensions.Unlock()
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
	mock.lockGetInstanceDimensions.RLock()
	calls = mock.calls.GetInstanceDimensions
	mock.lockGetInstanceDimensions.RUnlock()
	return calls
}

// PutInstanceDimensionOptionNodeID calls PutInstanceDimensionOptionNodeIDFunc.
func (mock *IClientMock) PutInstanceDimensionOptionNodeID(ctx context.Context, serviceAuthToken string, instanceID string, dimensionID string, optionID string, nodeID string) error {
	if mock.PutInstanceDimensionOptionNodeIDFunc == nil {
		panic("IClientMock.PutInstanceDimensionOptionNodeIDFunc: method is nil but IClient.PutInstanceDimensionOptionNodeID was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		DimensionID      string
		OptionID         string
		NodeID           string
	}{
		Ctx:              ctx,
		ServiceAuthToken: serviceAuthToken,
		InstanceID:       instanceID,
		DimensionID:      dimensionID,
		OptionID:         optionID,
		NodeID:           nodeID,
	}
	mock.lockPutInstanceDimensionOptionNodeID.Lock()
	mock.calls.PutInstanceDimensionOptionNodeID = append(mock.calls.PutInstanceDimensionOptionNodeID, callInfo)
	mock.lockPutInstanceDimensionOptionNodeID.Unlock()
	return mock.PutInstanceDimensionOptionNodeIDFunc(ctx, serviceAuthToken, instanceID, dimensionID, optionID, nodeID)
}

// PutInstanceDimensionOptionNodeIDCalls gets all the calls that were made to PutInstanceDimensionOptionNodeID.
// Check the length with:
//     len(mockedIClient.PutInstanceDimensionOptionNodeIDCalls())
func (mock *IClientMock) PutInstanceDimensionOptionNodeIDCalls() []struct {
	Ctx              context.Context
	ServiceAuthToken string
	InstanceID       string
	DimensionID      string
	OptionID         string
	NodeID           string
} {
	var calls []struct {
		Ctx              context.Context
		ServiceAuthToken string
		InstanceID       string
		DimensionID      string
		OptionID         string
		NodeID           string
	}
	mock.lockPutInstanceDimensionOptionNodeID.RLock()
	calls = mock.calls.PutInstanceDimensionOptionNodeID
	mock.lockPutInstanceDimensionOptionNodeID.RUnlock()
	return calls
}
