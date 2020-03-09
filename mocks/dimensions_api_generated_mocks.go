// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/ONSdigital/dp-dimension-importer/client"
	"io"
	"net/http"
	"sync"
)

var (
	lockHTTPClientMockDo sync.RWMutex
)

// Ensure, that HTTPClientMock does implement client.HTTPClient.
// If this is not the case, regenerate this file with moq.
var _ client.HTTPClient = &HTTPClientMock{}

// HTTPClientMock is a mock implementation of client.HTTPClient.
//
//     func TestSomethingThatUsesHTTPClient(t *testing.T) {
//
//         // make and configure a mocked client.HTTPClient
//         mockedHTTPClient := &HTTPClientMock{
//             DoFunc: func(req *http.Request) (*http.Response, error) {
// 	               panic("mock out the Do method")
//             },
//         }
//
//         // use mockedHTTPClient in code that requires client.HTTPClient
//         // and then make assertions.
//
//     }
type HTTPClientMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(req *http.Request) (*http.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// Req is the req argument value.
			Req *http.Request
		}
	}
}

// Do calls DoFunc.
func (mock *HTTPClientMock) Do(req *http.Request) (*http.Response, error) {
	if mock.DoFunc == nil {
		panic("HTTPClientMock.DoFunc: method is nil but HTTPClient.Do was just called")
	}
	callInfo := struct {
		Req *http.Request
	}{
		Req: req,
	}
	lockHTTPClientMockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	lockHTTPClientMockDo.Unlock()
	return mock.DoFunc(req)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedHTTPClient.DoCalls())
func (mock *HTTPClientMock) DoCalls() []struct {
	Req *http.Request
} {
	var calls []struct {
		Req *http.Request
	}
	lockHTTPClientMockDo.RLock()
	calls = mock.calls.Do
	lockHTTPClientMockDo.RUnlock()
	return calls
}

var (
	lockResponseBodyReaderMockRead sync.RWMutex
)

// Ensure, that ResponseBodyReaderMock does implement client.ResponseBodyReader.
// If this is not the case, regenerate this file with moq.
var _ client.ResponseBodyReader = &ResponseBodyReaderMock{}

// ResponseBodyReaderMock is a mock implementation of client.ResponseBodyReader.
//
//     func TestSomethingThatUsesResponseBodyReader(t *testing.T) {
//
//         // make and configure a mocked client.ResponseBodyReader
//         mockedResponseBodyReader := &ResponseBodyReaderMock{
//             ReadFunc: func(r io.Reader) ([]byte, error) {
// 	               panic("mock out the Read method")
//             },
//         }
//
//         // use mockedResponseBodyReader in code that requires client.ResponseBodyReader
//         // and then make assertions.
//
//     }
type ResponseBodyReaderMock struct {
	// ReadFunc mocks the Read method.
	ReadFunc func(r io.Reader) ([]byte, error)

	// calls tracks calls to the methods.
	calls struct {
		// Read holds details about calls to the Read method.
		Read []struct {
			// R is the r argument value.
			R io.Reader
		}
	}
}

// Read calls ReadFunc.
func (mock *ResponseBodyReaderMock) Read(r io.Reader) ([]byte, error) {
	if mock.ReadFunc == nil {
		panic("ResponseBodyReaderMock.ReadFunc: method is nil but ResponseBodyReader.Read was just called")
	}
	callInfo := struct {
		R io.Reader
	}{
		R: r,
	}
	lockResponseBodyReaderMockRead.Lock()
	mock.calls.Read = append(mock.calls.Read, callInfo)
	lockResponseBodyReaderMockRead.Unlock()
	return mock.ReadFunc(r)
}

// ReadCalls gets all the calls that were made to Read.
// Check the length with:
//     len(mockedResponseBodyReader.ReadCalls())
func (mock *ResponseBodyReaderMock) ReadCalls() []struct {
	R io.Reader
} {
	var calls []struct {
		R io.Reader
	}
	lockResponseBodyReaderMockRead.RLock()
	calls = mock.calls.Read
	lockResponseBodyReaderMockRead.RUnlock()
	return calls
}
