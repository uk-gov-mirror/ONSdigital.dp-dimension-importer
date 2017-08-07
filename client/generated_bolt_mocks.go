// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package client

import (
	"database/sql/driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"sync"
	"time"
)

var (
	lockNeoConnMockBegin           sync.RWMutex
	lockNeoConnMockClose           sync.RWMutex
	lockNeoConnMockExecNeo         sync.RWMutex
	lockNeoConnMockExecPipeline    sync.RWMutex
	lockNeoConnMockPrepareNeo      sync.RWMutex
	lockNeoConnMockPreparePipeline sync.RWMutex
	lockNeoConnMockQueryNeo        sync.RWMutex
	lockNeoConnMockQueryNeoAll     sync.RWMutex
	lockNeoConnMockQueryPipeline   sync.RWMutex
	lockNeoConnMockSetChunkSize    sync.RWMutex
	lockNeoConnMockSetTimeout      sync.RWMutex
)

// NeoConnMock is a mock implementation of NeoConn.
//
//     func TestSomethingThatUsesNeoConn(t *testing.T) {
//
//         // make and configure a mocked NeoConn
//         mockedNeoConn := &NeoConnMock{
//             BeginFunc: func() (driver.Tx, error) {
// 	               panic("TODO: mock out the Begin method")
//             },
//             CloseFunc: func() error {
// 	               panic("TODO: mock out the Close method")
//             },
//             ExecNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Result, error) {
// 	               panic("TODO: mock out the ExecNeo method")
//             },
//             ExecPipelineFunc: func(query []string, params ...map[string]interface{}) ([]golangNeo4jBoltDriver.Result, error) {
// 	               panic("TODO: mock out the ExecPipeline method")
//             },
//             PrepareNeoFunc: func(query string) (golangNeo4jBoltDriver.Stmt, error) {
// 	               panic("TODO: mock out the PrepareNeo method")
//             },
//             PreparePipelineFunc: func(query ...string) (golangNeo4jBoltDriver.PipelineStmt, error) {
// 	               panic("TODO: mock out the PreparePipeline method")
//             },
//             QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
// 	               panic("TODO: mock out the QueryNeo method")
//             },
//             QueryNeoAllFunc: func(query string, params map[string]interface{}) ([][]interface{}, map[string]interface{}, map[string]interface{}, error) {
// 	               panic("TODO: mock out the QueryNeoAll method")
//             },
//             QueryPipelineFunc: func(query []string, params ...map[string]interface{}) (golangNeo4jBoltDriver.PipelineRows, error) {
// 	               panic("TODO: mock out the QueryPipeline method")
//             },
//             SetChunkSizeFunc: func(in1 uint16)  {
// 	               panic("TODO: mock out the SetChunkSize method")
//             },
//             SetTimeoutFunc: func(in1 time.Duration)  {
// 	               panic("TODO: mock out the SetTimeout method")
//             },
//         }
//
//         // TODO: use mockedNeoConn in code that requires NeoConn
//         //       and then make assertions.
//
//     }
type NeoConnMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func() (driver.Tx, error)

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ExecNeoFunc mocks the ExecNeo method.
	ExecNeoFunc func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Result, error)

	// ExecPipelineFunc mocks the ExecPipeline method.
	ExecPipelineFunc func(query []string, params ...map[string]interface{}) ([]golangNeo4jBoltDriver.Result, error)

	// PrepareNeoFunc mocks the PrepareNeo method.
	PrepareNeoFunc func(query string) (golangNeo4jBoltDriver.Stmt, error)

	// PreparePipelineFunc mocks the PreparePipeline method.
	PreparePipelineFunc func(query ...string) (golangNeo4jBoltDriver.PipelineStmt, error)

	// QueryNeoFunc mocks the QueryNeo method.
	QueryNeoFunc func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error)

	// QueryNeoAllFunc mocks the QueryNeoAll method.
	QueryNeoAllFunc func(query string, params map[string]interface{}) ([][]interface{}, map[string]interface{}, map[string]interface{}, error)

	// QueryPipelineFunc mocks the QueryPipeline method.
	QueryPipelineFunc func(query []string, params ...map[string]interface{}) (golangNeo4jBoltDriver.PipelineRows, error)

	// SetChunkSizeFunc mocks the SetChunkSize method.
	SetChunkSizeFunc func(in1 uint16)

	// SetTimeoutFunc mocks the SetTimeout method.
	SetTimeoutFunc func(in1 time.Duration)

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// ExecNeo holds details about calls to the ExecNeo method.
		ExecNeo []struct {
			// Query is the query argument value.
			Query string
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// ExecPipeline holds details about calls to the ExecPipeline method.
		ExecPipeline []struct {
			// Query is the query argument value.
			Query []string
			// Params is the params argument value.
			Params []map[string]interface{}
		}
		// PrepareNeo holds details about calls to the PrepareNeo method.
		PrepareNeo []struct {
			// Query is the query argument value.
			Query string
		}
		// PreparePipeline holds details about calls to the PreparePipeline method.
		PreparePipeline []struct {
			// Query is the query argument value.
			Query []string
		}
		// QueryNeo holds details about calls to the QueryNeo method.
		QueryNeo []struct {
			// Query is the query argument value.
			Query string
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// QueryNeoAll holds details about calls to the QueryNeoAll method.
		QueryNeoAll []struct {
			// Query is the query argument value.
			Query string
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// QueryPipeline holds details about calls to the QueryPipeline method.
		QueryPipeline []struct {
			// Query is the query argument value.
			Query []string
			// Params is the params argument value.
			Params []map[string]interface{}
		}
		// SetChunkSize holds details about calls to the SetChunkSize method.
		SetChunkSize []struct {
			// In1 is the in1 argument value.
			In1 uint16
		}
		// SetTimeout holds details about calls to the SetTimeout method.
		SetTimeout []struct {
			// In1 is the in1 argument value.
			In1 time.Duration
		}
	}
}

// Begin calls BeginFunc.
func (mock *NeoConnMock) Begin() (driver.Tx, error) {
	if mock.BeginFunc == nil {
		panic("moq: NeoConnMock.BeginFunc is nil but NeoConn.Begin was just called")
	}
	callInfo := struct {
	}{}
	lockNeoConnMockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	lockNeoConnMockBegin.Unlock()
	return mock.BeginFunc()
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedNeoConn.BeginCalls())
func (mock *NeoConnMock) BeginCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoConnMockBegin.RLock()
	calls = mock.calls.Begin
	lockNeoConnMockBegin.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *NeoConnMock) Close() error {
	if mock.CloseFunc == nil {
		panic("moq: NeoConnMock.CloseFunc is nil but NeoConn.Close was just called")
	}
	callInfo := struct {
	}{}
	lockNeoConnMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockNeoConnMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedNeoConn.CloseCalls())
func (mock *NeoConnMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoConnMockClose.RLock()
	calls = mock.calls.Close
	lockNeoConnMockClose.RUnlock()
	return calls
}

// ExecNeo calls ExecNeoFunc.
func (mock *NeoConnMock) ExecNeo(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Result, error) {
	if mock.ExecNeoFunc == nil {
		panic("moq: NeoConnMock.ExecNeoFunc is nil but NeoConn.ExecNeo was just called")
	}
	callInfo := struct {
		Query  string
		Params map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockNeoConnMockExecNeo.Lock()
	mock.calls.ExecNeo = append(mock.calls.ExecNeo, callInfo)
	lockNeoConnMockExecNeo.Unlock()
	return mock.ExecNeoFunc(query, params)
}

// ExecNeoCalls gets all the calls that were made to ExecNeo.
// Check the length with:
//     len(mockedNeoConn.ExecNeoCalls())
func (mock *NeoConnMock) ExecNeoCalls() []struct {
	Query  string
	Params map[string]interface{}
} {
	var calls []struct {
		Query  string
		Params map[string]interface{}
	}
	lockNeoConnMockExecNeo.RLock()
	calls = mock.calls.ExecNeo
	lockNeoConnMockExecNeo.RUnlock()
	return calls
}

// ExecPipeline calls ExecPipelineFunc.
func (mock *NeoConnMock) ExecPipeline(query []string, params ...map[string]interface{}) ([]golangNeo4jBoltDriver.Result, error) {
	if mock.ExecPipelineFunc == nil {
		panic("moq: NeoConnMock.ExecPipelineFunc is nil but NeoConn.ExecPipeline was just called")
	}
	callInfo := struct {
		Query  []string
		Params []map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockNeoConnMockExecPipeline.Lock()
	mock.calls.ExecPipeline = append(mock.calls.ExecPipeline, callInfo)
	lockNeoConnMockExecPipeline.Unlock()
	return mock.ExecPipelineFunc(query, params...)
}

// ExecPipelineCalls gets all the calls that were made to ExecPipeline.
// Check the length with:
//     len(mockedNeoConn.ExecPipelineCalls())
func (mock *NeoConnMock) ExecPipelineCalls() []struct {
	Query  []string
	Params []map[string]interface{}
} {
	var calls []struct {
		Query  []string
		Params []map[string]interface{}
	}
	lockNeoConnMockExecPipeline.RLock()
	calls = mock.calls.ExecPipeline
	lockNeoConnMockExecPipeline.RUnlock()
	return calls
}

// PrepareNeo calls PrepareNeoFunc.
func (mock *NeoConnMock) PrepareNeo(query string) (golangNeo4jBoltDriver.Stmt, error) {
	if mock.PrepareNeoFunc == nil {
		panic("moq: NeoConnMock.PrepareNeoFunc is nil but NeoConn.PrepareNeo was just called")
	}
	callInfo := struct {
		Query string
	}{
		Query: query,
	}
	lockNeoConnMockPrepareNeo.Lock()
	mock.calls.PrepareNeo = append(mock.calls.PrepareNeo, callInfo)
	lockNeoConnMockPrepareNeo.Unlock()
	return mock.PrepareNeoFunc(query)
}

// PrepareNeoCalls gets all the calls that were made to PrepareNeo.
// Check the length with:
//     len(mockedNeoConn.PrepareNeoCalls())
func (mock *NeoConnMock) PrepareNeoCalls() []struct {
	Query string
} {
	var calls []struct {
		Query string
	}
	lockNeoConnMockPrepareNeo.RLock()
	calls = mock.calls.PrepareNeo
	lockNeoConnMockPrepareNeo.RUnlock()
	return calls
}

// PreparePipeline calls PreparePipelineFunc.
func (mock *NeoConnMock) PreparePipeline(query ...string) (golangNeo4jBoltDriver.PipelineStmt, error) {
	if mock.PreparePipelineFunc == nil {
		panic("moq: NeoConnMock.PreparePipelineFunc is nil but NeoConn.PreparePipeline was just called")
	}
	callInfo := struct {
		Query []string
	}{
		Query: query,
	}
	lockNeoConnMockPreparePipeline.Lock()
	mock.calls.PreparePipeline = append(mock.calls.PreparePipeline, callInfo)
	lockNeoConnMockPreparePipeline.Unlock()
	return mock.PreparePipelineFunc(query...)
}

// PreparePipelineCalls gets all the calls that were made to PreparePipeline.
// Check the length with:
//     len(mockedNeoConn.PreparePipelineCalls())
func (mock *NeoConnMock) PreparePipelineCalls() []struct {
	Query []string
} {
	var calls []struct {
		Query []string
	}
	lockNeoConnMockPreparePipeline.RLock()
	calls = mock.calls.PreparePipeline
	lockNeoConnMockPreparePipeline.RUnlock()
	return calls
}

// QueryNeo calls QueryNeoFunc.
func (mock *NeoConnMock) QueryNeo(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
	if mock.QueryNeoFunc == nil {
		panic("moq: NeoConnMock.QueryNeoFunc is nil but NeoConn.QueryNeo was just called")
	}
	callInfo := struct {
		Query  string
		Params map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockNeoConnMockQueryNeo.Lock()
	mock.calls.QueryNeo = append(mock.calls.QueryNeo, callInfo)
	lockNeoConnMockQueryNeo.Unlock()
	return mock.QueryNeoFunc(query, params)
}

// QueryNeoCalls gets all the calls that were made to QueryNeo.
// Check the length with:
//     len(mockedNeoConn.QueryNeoCalls())
func (mock *NeoConnMock) QueryNeoCalls() []struct {
	Query  string
	Params map[string]interface{}
} {
	var calls []struct {
		Query  string
		Params map[string]interface{}
	}
	lockNeoConnMockQueryNeo.RLock()
	calls = mock.calls.QueryNeo
	lockNeoConnMockQueryNeo.RUnlock()
	return calls
}

// QueryNeoAll calls QueryNeoAllFunc.
func (mock *NeoConnMock) QueryNeoAll(query string, params map[string]interface{}) ([][]interface{}, map[string]interface{}, map[string]interface{}, error) {
	if mock.QueryNeoAllFunc == nil {
		panic("moq: NeoConnMock.QueryNeoAllFunc is nil but NeoConn.QueryNeoAll was just called")
	}
	callInfo := struct {
		Query  string
		Params map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockNeoConnMockQueryNeoAll.Lock()
	mock.calls.QueryNeoAll = append(mock.calls.QueryNeoAll, callInfo)
	lockNeoConnMockQueryNeoAll.Unlock()
	return mock.QueryNeoAllFunc(query, params)
}

// QueryNeoAllCalls gets all the calls that were made to QueryNeoAll.
// Check the length with:
//     len(mockedNeoConn.QueryNeoAllCalls())
func (mock *NeoConnMock) QueryNeoAllCalls() []struct {
	Query  string
	Params map[string]interface{}
} {
	var calls []struct {
		Query  string
		Params map[string]interface{}
	}
	lockNeoConnMockQueryNeoAll.RLock()
	calls = mock.calls.QueryNeoAll
	lockNeoConnMockQueryNeoAll.RUnlock()
	return calls
}

// QueryPipeline calls QueryPipelineFunc.
func (mock *NeoConnMock) QueryPipeline(query []string, params ...map[string]interface{}) (golangNeo4jBoltDriver.PipelineRows, error) {
	if mock.QueryPipelineFunc == nil {
		panic("moq: NeoConnMock.QueryPipelineFunc is nil but NeoConn.QueryPipeline was just called")
	}
	callInfo := struct {
		Query  []string
		Params []map[string]interface{}
	}{
		Query:  query,
		Params: params,
	}
	lockNeoConnMockQueryPipeline.Lock()
	mock.calls.QueryPipeline = append(mock.calls.QueryPipeline, callInfo)
	lockNeoConnMockQueryPipeline.Unlock()
	return mock.QueryPipelineFunc(query, params...)
}

// QueryPipelineCalls gets all the calls that were made to QueryPipeline.
// Check the length with:
//     len(mockedNeoConn.QueryPipelineCalls())
func (mock *NeoConnMock) QueryPipelineCalls() []struct {
	Query  []string
	Params []map[string]interface{}
} {
	var calls []struct {
		Query  []string
		Params []map[string]interface{}
	}
	lockNeoConnMockQueryPipeline.RLock()
	calls = mock.calls.QueryPipeline
	lockNeoConnMockQueryPipeline.RUnlock()
	return calls
}

// SetChunkSize calls SetChunkSizeFunc.
func (mock *NeoConnMock) SetChunkSize(in1 uint16) {
	if mock.SetChunkSizeFunc == nil {
		panic("moq: NeoConnMock.SetChunkSizeFunc is nil but NeoConn.SetChunkSize was just called")
	}
	callInfo := struct {
		In1 uint16
	}{
		In1: in1,
	}
	lockNeoConnMockSetChunkSize.Lock()
	mock.calls.SetChunkSize = append(mock.calls.SetChunkSize, callInfo)
	lockNeoConnMockSetChunkSize.Unlock()
	mock.SetChunkSizeFunc(in1)
}

// SetChunkSizeCalls gets all the calls that were made to SetChunkSize.
// Check the length with:
//     len(mockedNeoConn.SetChunkSizeCalls())
func (mock *NeoConnMock) SetChunkSizeCalls() []struct {
	In1 uint16
} {
	var calls []struct {
		In1 uint16
	}
	lockNeoConnMockSetChunkSize.RLock()
	calls = mock.calls.SetChunkSize
	lockNeoConnMockSetChunkSize.RUnlock()
	return calls
}

// SetTimeout calls SetTimeoutFunc.
func (mock *NeoConnMock) SetTimeout(in1 time.Duration) {
	if mock.SetTimeoutFunc == nil {
		panic("moq: NeoConnMock.SetTimeoutFunc is nil but NeoConn.SetTimeout was just called")
	}
	callInfo := struct {
		In1 time.Duration
	}{
		In1: in1,
	}
	lockNeoConnMockSetTimeout.Lock()
	mock.calls.SetTimeout = append(mock.calls.SetTimeout, callInfo)
	lockNeoConnMockSetTimeout.Unlock()
	mock.SetTimeoutFunc(in1)
}

// SetTimeoutCalls gets all the calls that were made to SetTimeout.
// Check the length with:
//     len(mockedNeoConn.SetTimeoutCalls())
func (mock *NeoConnMock) SetTimeoutCalls() []struct {
	In1 time.Duration
} {
	var calls []struct {
		In1 time.Duration
	}
	lockNeoConnMockSetTimeout.RLock()
	calls = mock.calls.SetTimeout
	lockNeoConnMockSetTimeout.RUnlock()
	return calls
}

var (
	lockNeoDriverPoolMockOpenPool sync.RWMutex
)

// NeoDriverPoolMock is a mock implementation of NeoDriverPool.
//
//     func TestSomethingThatUsesNeoDriverPool(t *testing.T) {
//
//         // make and configure a mocked NeoDriverPool
//         mockedNeoDriverPool := &NeoDriverPoolMock{
//             OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
// 	               panic("TODO: mock out the OpenPool method")
//             },
//         }
//
//         // TODO: use mockedNeoDriverPool in code that requires NeoDriverPool
//         //       and then make assertions.
//
//     }
type NeoDriverPoolMock struct {
	// OpenPoolFunc mocks the OpenPool method.
	OpenPoolFunc func() (golangNeo4jBoltDriver.Conn, error)

	// calls tracks calls to the methods.
	calls struct {
		// OpenPool holds details about calls to the OpenPool method.
		OpenPool []struct {
		}
	}
}

// OpenPool calls OpenPoolFunc.
func (mock *NeoDriverPoolMock) OpenPool() (golangNeo4jBoltDriver.Conn, error) {
	if mock.OpenPoolFunc == nil {
		panic("moq: NeoDriverPoolMock.OpenPoolFunc is nil but NeoDriverPool.OpenPool was just called")
	}
	callInfo := struct {
	}{}
	lockNeoDriverPoolMockOpenPool.Lock()
	mock.calls.OpenPool = append(mock.calls.OpenPool, callInfo)
	lockNeoDriverPoolMockOpenPool.Unlock()
	return mock.OpenPoolFunc()
}

// OpenPoolCalls gets all the calls that were made to OpenPool.
// Check the length with:
//     len(mockedNeoDriverPool.OpenPoolCalls())
func (mock *NeoDriverPoolMock) OpenPoolCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoDriverPoolMockOpenPool.RLock()
	calls = mock.calls.OpenPool
	lockNeoDriverPoolMockOpenPool.RUnlock()
	return calls
}

var (
	lockNeoStmtMockClose    sync.RWMutex
	lockNeoStmtMockExecNeo  sync.RWMutex
	lockNeoStmtMockQueryNeo sync.RWMutex
)

// NeoStmtMock is a mock implementation of NeoStmt.
//
//     func TestSomethingThatUsesNeoStmt(t *testing.T) {
//
//         // make and configure a mocked NeoStmt
//         mockedNeoStmt := &NeoStmtMock{
//             CloseFunc: func() error {
// 	               panic("TODO: mock out the Close method")
//             },
//             ExecNeoFunc: func(params map[string]interface{}) (golangNeo4jBoltDriver.Result, error) {
// 	               panic("TODO: mock out the ExecNeo method")
//             },
//             QueryNeoFunc: func(params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
// 	               panic("TODO: mock out the QueryNeo method")
//             },
//         }
//
//         // TODO: use mockedNeoStmt in code that requires NeoStmt
//         //       and then make assertions.
//
//     }
type NeoStmtMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ExecNeoFunc mocks the ExecNeo method.
	ExecNeoFunc func(params map[string]interface{}) (golangNeo4jBoltDriver.Result, error)

	// QueryNeoFunc mocks the QueryNeo method.
	QueryNeoFunc func(params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// ExecNeo holds details about calls to the ExecNeo method.
		ExecNeo []struct {
			// Params is the params argument value.
			Params map[string]interface{}
		}
		// QueryNeo holds details about calls to the QueryNeo method.
		QueryNeo []struct {
			// Params is the params argument value.
			Params map[string]interface{}
		}
	}
}

// Close calls CloseFunc.
func (mock *NeoStmtMock) Close() error {
	if mock.CloseFunc == nil {
		panic("moq: NeoStmtMock.CloseFunc is nil but NeoStmt.Close was just called")
	}
	callInfo := struct {
	}{}
	lockNeoStmtMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockNeoStmtMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedNeoStmt.CloseCalls())
func (mock *NeoStmtMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoStmtMockClose.RLock()
	calls = mock.calls.Close
	lockNeoStmtMockClose.RUnlock()
	return calls
}

// ExecNeo calls ExecNeoFunc.
func (mock *NeoStmtMock) ExecNeo(params map[string]interface{}) (golangNeo4jBoltDriver.Result, error) {
	if mock.ExecNeoFunc == nil {
		panic("moq: NeoStmtMock.ExecNeoFunc is nil but NeoStmt.ExecNeo was just called")
	}
	callInfo := struct {
		Params map[string]interface{}
	}{
		Params: params,
	}
	lockNeoStmtMockExecNeo.Lock()
	mock.calls.ExecNeo = append(mock.calls.ExecNeo, callInfo)
	lockNeoStmtMockExecNeo.Unlock()
	return mock.ExecNeoFunc(params)
}

// ExecNeoCalls gets all the calls that were made to ExecNeo.
// Check the length with:
//     len(mockedNeoStmt.ExecNeoCalls())
func (mock *NeoStmtMock) ExecNeoCalls() []struct {
	Params map[string]interface{}
} {
	var calls []struct {
		Params map[string]interface{}
	}
	lockNeoStmtMockExecNeo.RLock()
	calls = mock.calls.ExecNeo
	lockNeoStmtMockExecNeo.RUnlock()
	return calls
}

// QueryNeo calls QueryNeoFunc.
func (mock *NeoStmtMock) QueryNeo(params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
	if mock.QueryNeoFunc == nil {
		panic("moq: NeoStmtMock.QueryNeoFunc is nil but NeoStmt.QueryNeo was just called")
	}
	callInfo := struct {
		Params map[string]interface{}
	}{
		Params: params,
	}
	lockNeoStmtMockQueryNeo.Lock()
	mock.calls.QueryNeo = append(mock.calls.QueryNeo, callInfo)
	lockNeoStmtMockQueryNeo.Unlock()
	return mock.QueryNeoFunc(params)
}

// QueryNeoCalls gets all the calls that were made to QueryNeo.
// Check the length with:
//     len(mockedNeoStmt.QueryNeoCalls())
func (mock *NeoStmtMock) QueryNeoCalls() []struct {
	Params map[string]interface{}
} {
	var calls []struct {
		Params map[string]interface{}
	}
	lockNeoStmtMockQueryNeo.RLock()
	calls = mock.calls.QueryNeo
	lockNeoStmtMockQueryNeo.RUnlock()
	return calls
}

var (
	lockNeoQueryRowsMockAll      sync.RWMutex
	lockNeoQueryRowsMockClose    sync.RWMutex
	lockNeoQueryRowsMockColumns  sync.RWMutex
	lockNeoQueryRowsMockMetadata sync.RWMutex
	lockNeoQueryRowsMockNextNeo  sync.RWMutex
)

// NeoQueryRowsMock is a mock implementation of NeoQueryRows.
//
//     func TestSomethingThatUsesNeoQueryRows(t *testing.T) {
//
//         // make and configure a mocked NeoQueryRows
//         mockedNeoQueryRows := &NeoQueryRowsMock{
//             AllFunc: func() ([][]interface{}, map[string]interface{}, error) {
// 	               panic("TODO: mock out the All method")
//             },
//             CloseFunc: func() error {
// 	               panic("TODO: mock out the Close method")
//             },
//             ColumnsFunc: func() []string {
// 	               panic("TODO: mock out the Columns method")
//             },
//             MetadataFunc: func() map[string]interface{} {
// 	               panic("TODO: mock out the Metadata method")
//             },
//             NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
// 	               panic("TODO: mock out the NextNeo method")
//             },
//         }
//
//         // TODO: use mockedNeoQueryRows in code that requires NeoQueryRows
//         //       and then make assertions.
//
//     }
type NeoQueryRowsMock struct {
	// AllFunc mocks the All method.
	AllFunc func() ([][]interface{}, map[string]interface{}, error)

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// ColumnsFunc mocks the Columns method.
	ColumnsFunc func() []string

	// MetadataFunc mocks the Metadata method.
	MetadataFunc func() map[string]interface{}

	// NextNeoFunc mocks the NextNeo method.
	NextNeoFunc func() ([]interface{}, map[string]interface{}, error)

	// calls tracks calls to the methods.
	calls struct {
		// All holds details about calls to the All method.
		All []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Columns holds details about calls to the Columns method.
		Columns []struct {
		}
		// Metadata holds details about calls to the Metadata method.
		Metadata []struct {
		}
		// NextNeo holds details about calls to the NextNeo method.
		NextNeo []struct {
		}
	}
}

// All calls AllFunc.
func (mock *NeoQueryRowsMock) All() ([][]interface{}, map[string]interface{}, error) {
	if mock.AllFunc == nil {
		panic("moq: NeoQueryRowsMock.AllFunc is nil but NeoQueryRows.All was just called")
	}
	callInfo := struct {
	}{}
	lockNeoQueryRowsMockAll.Lock()
	mock.calls.All = append(mock.calls.All, callInfo)
	lockNeoQueryRowsMockAll.Unlock()
	return mock.AllFunc()
}

// AllCalls gets all the calls that were made to All.
// Check the length with:
//     len(mockedNeoQueryRows.AllCalls())
func (mock *NeoQueryRowsMock) AllCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoQueryRowsMockAll.RLock()
	calls = mock.calls.All
	lockNeoQueryRowsMockAll.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *NeoQueryRowsMock) Close() error {
	if mock.CloseFunc == nil {
		panic("moq: NeoQueryRowsMock.CloseFunc is nil but NeoQueryRows.Close was just called")
	}
	callInfo := struct {
	}{}
	lockNeoQueryRowsMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockNeoQueryRowsMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedNeoQueryRows.CloseCalls())
func (mock *NeoQueryRowsMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoQueryRowsMockClose.RLock()
	calls = mock.calls.Close
	lockNeoQueryRowsMockClose.RUnlock()
	return calls
}

// Columns calls ColumnsFunc.
func (mock *NeoQueryRowsMock) Columns() []string {
	if mock.ColumnsFunc == nil {
		panic("moq: NeoQueryRowsMock.ColumnsFunc is nil but NeoQueryRows.Columns was just called")
	}
	callInfo := struct {
	}{}
	lockNeoQueryRowsMockColumns.Lock()
	mock.calls.Columns = append(mock.calls.Columns, callInfo)
	lockNeoQueryRowsMockColumns.Unlock()
	return mock.ColumnsFunc()
}

// ColumnsCalls gets all the calls that were made to Columns.
// Check the length with:
//     len(mockedNeoQueryRows.ColumnsCalls())
func (mock *NeoQueryRowsMock) ColumnsCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoQueryRowsMockColumns.RLock()
	calls = mock.calls.Columns
	lockNeoQueryRowsMockColumns.RUnlock()
	return calls
}

// Metadata calls MetadataFunc.
func (mock *NeoQueryRowsMock) Metadata() map[string]interface{} {
	if mock.MetadataFunc == nil {
		panic("moq: NeoQueryRowsMock.MetadataFunc is nil but NeoQueryRows.Metadata was just called")
	}
	callInfo := struct {
	}{}
	lockNeoQueryRowsMockMetadata.Lock()
	mock.calls.Metadata = append(mock.calls.Metadata, callInfo)
	lockNeoQueryRowsMockMetadata.Unlock()
	return mock.MetadataFunc()
}

// MetadataCalls gets all the calls that were made to Metadata.
// Check the length with:
//     len(mockedNeoQueryRows.MetadataCalls())
func (mock *NeoQueryRowsMock) MetadataCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoQueryRowsMockMetadata.RLock()
	calls = mock.calls.Metadata
	lockNeoQueryRowsMockMetadata.RUnlock()
	return calls
}

// NextNeo calls NextNeoFunc.
func (mock *NeoQueryRowsMock) NextNeo() ([]interface{}, map[string]interface{}, error) {
	if mock.NextNeoFunc == nil {
		panic("moq: NeoQueryRowsMock.NextNeoFunc is nil but NeoQueryRows.NextNeo was just called")
	}
	callInfo := struct {
	}{}
	lockNeoQueryRowsMockNextNeo.Lock()
	mock.calls.NextNeo = append(mock.calls.NextNeo, callInfo)
	lockNeoQueryRowsMockNextNeo.Unlock()
	return mock.NextNeoFunc()
}

// NextNeoCalls gets all the calls that were made to NextNeo.
// Check the length with:
//     len(mockedNeoQueryRows.NextNeoCalls())
func (mock *NeoQueryRowsMock) NextNeoCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoQueryRowsMockNextNeo.RLock()
	calls = mock.calls.NextNeo
	lockNeoQueryRowsMockNextNeo.RUnlock()
	return calls
}

var (
	lockNeoResultMockLastInsertId sync.RWMutex
	lockNeoResultMockMetadata     sync.RWMutex
	lockNeoResultMockRowsAffected sync.RWMutex
)

// NeoResultMock is a mock implementation of NeoResult.
//
//     func TestSomethingThatUsesNeoResult(t *testing.T) {
//
//         // make and configure a mocked NeoResult
//         mockedNeoResult := &NeoResultMock{
//             LastInsertIdFunc: func() (int64, error) {
// 	               panic("TODO: mock out the LastInsertId method")
//             },
//             MetadataFunc: func() map[string]interface{} {
// 	               panic("TODO: mock out the Metadata method")
//             },
//             RowsAffectedFunc: func() (int64, error) {
// 	               panic("TODO: mock out the RowsAffected method")
//             },
//         }
//
//         // TODO: use mockedNeoResult in code that requires NeoResult
//         //       and then make assertions.
//
//     }
type NeoResultMock struct {
	// LastInsertIdFunc mocks the LastInsertId method.
	LastInsertIdFunc func() (int64, error)

	// MetadataFunc mocks the Metadata method.
	MetadataFunc func() map[string]interface{}

	// RowsAffectedFunc mocks the RowsAffected method.
	RowsAffectedFunc func() (int64, error)

	// calls tracks calls to the methods.
	calls struct {
		// LastInsertId holds details about calls to the LastInsertId method.
		LastInsertId []struct {
		}
		// Metadata holds details about calls to the Metadata method.
		Metadata []struct {
		}
		// RowsAffected holds details about calls to the RowsAffected method.
		RowsAffected []struct {
		}
	}
}

// LastInsertId calls LastInsertIdFunc.
func (mock *NeoResultMock) LastInsertId() (int64, error) {
	if mock.LastInsertIdFunc == nil {
		panic("moq: NeoResultMock.LastInsertIdFunc is nil but NeoResult.LastInsertId was just called")
	}
	callInfo := struct {
	}{}
	lockNeoResultMockLastInsertId.Lock()
	mock.calls.LastInsertId = append(mock.calls.LastInsertId, callInfo)
	lockNeoResultMockLastInsertId.Unlock()
	return mock.LastInsertIdFunc()
}

// LastInsertIdCalls gets all the calls that were made to LastInsertId.
// Check the length with:
//     len(mockedNeoResult.LastInsertIdCalls())
func (mock *NeoResultMock) LastInsertIdCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoResultMockLastInsertId.RLock()
	calls = mock.calls.LastInsertId
	lockNeoResultMockLastInsertId.RUnlock()
	return calls
}

// Metadata calls MetadataFunc.
func (mock *NeoResultMock) Metadata() map[string]interface{} {
	if mock.MetadataFunc == nil {
		panic("moq: NeoResultMock.MetadataFunc is nil but NeoResult.Metadata was just called")
	}
	callInfo := struct {
	}{}
	lockNeoResultMockMetadata.Lock()
	mock.calls.Metadata = append(mock.calls.Metadata, callInfo)
	lockNeoResultMockMetadata.Unlock()
	return mock.MetadataFunc()
}

// MetadataCalls gets all the calls that were made to Metadata.
// Check the length with:
//     len(mockedNeoResult.MetadataCalls())
func (mock *NeoResultMock) MetadataCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoResultMockMetadata.RLock()
	calls = mock.calls.Metadata
	lockNeoResultMockMetadata.RUnlock()
	return calls
}

// RowsAffected calls RowsAffectedFunc.
func (mock *NeoResultMock) RowsAffected() (int64, error) {
	if mock.RowsAffectedFunc == nil {
		panic("moq: NeoResultMock.RowsAffectedFunc is nil but NeoResult.RowsAffected was just called")
	}
	callInfo := struct {
	}{}
	lockNeoResultMockRowsAffected.Lock()
	mock.calls.RowsAffected = append(mock.calls.RowsAffected, callInfo)
	lockNeoResultMockRowsAffected.Unlock()
	return mock.RowsAffectedFunc()
}

// RowsAffectedCalls gets all the calls that were made to RowsAffected.
// Check the length with:
//     len(mockedNeoResult.RowsAffectedCalls())
func (mock *NeoResultMock) RowsAffectedCalls() []struct {
} {
	var calls []struct {
	}
	lockNeoResultMockRowsAffected.RLock()
	calls = mock.calls.RowsAffected
	lockNeoResultMockRowsAffected.RUnlock()
	return calls
}
