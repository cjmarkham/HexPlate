package testutils

type MockParameters struct {
	Response interface{}
	Error    error
	Times    int
}
