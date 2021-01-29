package main

import "fmt"

type error interface {
	Error() string
}

type RPCError struct {
	Code    int64
	Message string
}

// 接口类型检查的时机
func main() {
	var rpcErr error = NewRPCError(400, "unknown error")
	err := AsErr(rpcErr)
	println(err)
}

// RPCError 继承了 error interface
func (e *RPCError) Error() string {
	return fmt.Sprintf("%s, code %d", e.Message, e.Code)
}

func AsErr(err error) error {
	return err
}

func NewRPCError(code int64, msg string) error {
	return &RPCError{
		Code:    code,
		Message: msg,
	}
}
