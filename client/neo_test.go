package client

import (
	"testing"

	"github.com/ONSdigital/dp-dimension-importer/common"
	"github.com/ONSdigital/dp-dimension-importer/mocks"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	rowsData  [][]interface{} = [][]interface{}{{1}}
	mockError                 = errors.New("mock error")
)

func TestNeo4j_Query_ConnNil(t *testing.T) {
	neo4j := Neo4j{}

	Convey("Given a conn is nil", t, func() {
		Convey("When neo4j.Query is invoked", func() {
			rows, err := neo4j.Query(nil, "", nil)

			Convey("Then the expected error is returned and rows is nil", func() {
				So(err.Error(), ShouldEqual, connNilErr.Error())
				So(rows, ShouldBeNil)
			})
		})
	})
}

func TestNeo4j_Query_EmptyQuery(t *testing.T) {
	neo4j := Neo4j{}
	Convey("Given a conn and an empty query", t, func() {
		connMock := &mocks.NeoConnMock{}

		Convey("When neo4j.Query is invoked", func() {
			rows, err := neo4j.Query(connMock, "", nil)

			Convey("Then the expected error is returned and rows is nil", func() {
				So(err.Error(), ShouldEqual, errors.New("query required but was empty").Error())
				So(rows, ShouldBeNil)
			})

			Convey("And conn.PrepareNeo is never invoked", func() {
				So(len(connMock.PrepareNeoCalls()), ShouldEqual, 0)
			})
		})
	})
}

func TestNeo4j_Query_ConnPrepareNeoErr(t *testing.T) {
	neo4j := Neo4j{}

	Convey("Given a valid conn and query", t, func() {

		Convey("When conn.PrepareNeo returns an error", func() {
			connMock := &mocks.NeoConnMock{
				PrepareNeoFunc: func(query string) (bolt.Stmt, error) {
					return nil, mockError
				},
			}
			rows, err := neo4j.Query(connMock, "MATCH (N) RETURN N LIMIT 5", nil)

			Convey("Then the expected error is returned and rows is nil", func() {
				So(err.Error(), ShouldEqual, errors.Wrap(mockError, "error while attempting to create statement").Error())
				So(rows, ShouldBeNil)
			})

			Convey("And conn.PrepareNeo is invoked 1 time", func() {
				So(len(connMock.PrepareNeoCalls()), ShouldEqual, 1)
			})
		})
	})
}

func TestNeo4j_Query_StmtQueryNeoErr(t *testing.T) {
	neo4j := Neo4j{}
	Convey("Given a valid conn and query", t, func() {

		Convey("When neoStmt.QueryNeo returns an error", func() {
			stmtMock := &mocks.NeoStmtMock{
				QueryNeoFunc: func(params map[string]interface{}) (bolt.Rows, error) {
					return nil, mockError
				},
				CloseFunc: closeNoErr,
			}

			connMock := &mocks.NeoConnMock{
				PrepareNeoFunc: func(query string) (bolt.Stmt, error) {
					return stmtMock, nil
				},
			}

			rows, err := neo4j.Query(connMock, "MATCH (N) RETURN N LIMIT 5", nil)

			Convey("Then the expected error is returned and rows is nil", func() {
				So(err.Error(), ShouldEqual, errors.Wrap(mockError, "error while attempting to execute query").Error())
				So(rows, ShouldBeNil)
			})

			Convey("And conn.PrepareNeo is invoked 1 time", func() {
				So(len(connMock.PrepareNeoCalls()), ShouldEqual, 1)
			})

			Convey("And neoStmt.QueryNeo is invoked 1 time", func() {
				So(len(stmtMock.QueryNeoCalls()), ShouldEqual, 1)
			})

			Convey("And neoStmt.Close is invoked 1 time", func() {
				So(len(stmtMock.CloseCalls()), ShouldEqual, 1)
			})
		})
	})
}

func TestNeo4j_Query_Success(t *testing.T) {
	neo4j := Neo4j{}
	Convey("Given a valid conn and query", t, func() {

		Convey("When neo4j.Query is invoked", func() {

			rowsMock := &mocks.NeoQueryRowsMock{
				AllFunc: func() ([][]interface{}, map[string]interface{}, error) {
					return rowsData, nil, nil
				},
				CloseFunc: closeNoErr,
			}

			stmtMock := &mocks.NeoStmtMock{
				QueryNeoFunc: func(params map[string]interface{}) (bolt.Rows, error) {
					return rowsMock, nil
				},
				CloseFunc: closeNoErr,
			}

			connMock := &mocks.NeoConnMock{
				PrepareNeoFunc: func(query string) (bolt.Stmt, error) {
					return stmtMock, nil
				},
			}

			rows, err := neo4j.Query(connMock, "MATCH (N) RETURN N LIMIT 5", nil)

			Convey("Then the expected value is returned and error is nil", func() {
				expected := &common.NeoRows{Data: rowsData}
				So(rows, ShouldResemble, expected)
				So(err, ShouldBeNil)
			})

			Convey("And conn.PrepareNeo is invoked 1 time", func() {
				So(len(connMock.PrepareNeoCalls()), ShouldEqual, 1)
			})

			Convey("And neoStmt.QueryNeo is invoked 1 time", func() {
				So(len(stmtMock.QueryNeoCalls()), ShouldEqual, 1)
			})

			Convey("And close is called 1 times on stmt & rows", func() {
				So(len(stmtMock.CloseCalls()), ShouldEqual, 1)
				So(len(rowsMock.CloseCalls()), ShouldEqual, 1)
			})
		})
	})
}

func TestNeo4j_ExecStmt(t *testing.T) {
	Convey("Given valid parameters", t, func() {
		stmt := "Valar morghulis"
		params := map[string]interface{}{"param1": "Valar dohaeris"}

		mockResult := &mocks.NeoResultMock{}
		mockStmt := &mocks.NeoStmtMock{
			ExecNeoFunc: func(params map[string]interface{}) (bolt.Result, error) {
				return mockResult, nil
			},
			CloseFunc: closeNoErr,
		}
		mockConn := &mocks.NeoConnMock{
			PrepareNeoFunc: func(query string) (bolt.Stmt, error) {
				return mockStmt, nil
			},
			CloseFunc: closeNoErr,
		}

		neo := Neo4j{}

		Convey("When ExecStmt is called with valid parameters", func() {

			results, err := neo.ExecStmt(mockConn, stmt, params)

			Convey("Then the expected result should be returned and no errors", func() {
				So(err, ShouldEqual, nil)
				So(results, ShouldResemble, mockResult)
			})

			Convey("And conn.PrepareNeo is called 1 time with the expected parameters", func() {
				So(len(mockConn.PrepareNeoCalls()), ShouldEqual, 1)
				So(mockConn.PrepareNeoCalls()[0].Query, ShouldEqual, stmt)
			})

			Convey("And neoStmt.ExecNeo is called 1 time with the expected parameters", func() {
				calls := mockStmt.ExecNeoCalls()
				So(len(calls), ShouldEqual, 1)
				So(calls[0].Params, ShouldEqual, params)
			})

			Convey("And stmt.Close is called 1 time", func() {
				So(len(mockStmt.CloseCalls()), ShouldEqual, 1)
			})
		})

		Convey("When ExecStmt is called with an empty stmt parameter", func() {
			stmt = ""
			results, err := neo.ExecStmt(mockConn, stmt, params)

			Convey("Then no results and the expected error are returned", func() {
				So(err.Error(), ShouldEqual, errors.New("statement required but was empty").Error())
				So(results, ShouldEqual, nil)
			})
		})

		Convey("When neoStmt.PrepareNeo returns an error", func() {

			mockConn.PrepareNeoFunc = func(query string) (bolt.Stmt, error) {
				return nil, mockError
			}

			results, err := neo.ExecStmt(mockConn, stmt, params)

			Convey("And conn.PrepareNeo is called 1 time with the expected params", func() {
				calls := mockConn.PrepareNeoCalls()
				So(len(calls), ShouldEqual, 1)
				So(calls[0].Query, ShouldEqual, stmt)
			})

			Convey("And the expected error response is returned", func() {
				So(err.Error(), ShouldEqual, errors.Wrap(mockError, "error while attempting to create statement").Error())
				So(results, ShouldEqual, nil)
			})
		})

		Convey("When neoStmt.ExecNeo returns an error", func() {

			mockStmt.ExecNeoFunc = func(params map[string]interface{}) (bolt.Result, error) {
				return nil, mockError
			}

			results, err := neo.ExecStmt(mockConn, stmt, params)

			Convey("And conn.PrepareNeo should be called 1 time with the expected parameters", func() {
				So(len(mockConn.PrepareNeoCalls()), ShouldEqual, 1)
				So(mockConn.PrepareNeoCalls()[0].Query, ShouldEqual, stmt)
			})

			Convey("And neoStmt.ExecNeo is called 1 time with the expected parameters", func() {
				calls := mockStmt.ExecNeoCalls()
				So(len(calls), ShouldEqual, 1)
				So(calls[0].Params, ShouldEqual, params)
			})

			Convey("And stmt.Close is invoked 1 time", func() {
				So(len(mockStmt.CloseCalls()), ShouldEqual, 1)
			})

			Convey("And the expected error response is returned", func() {
				So(err.Error(), ShouldEqual, errors.Wrap(mockError, "error while attempting to execute neo statement").Error())
				So(results, ShouldEqual, nil)
			})
		})
	})
}

func closeNoErr() error {
	return nil
}
