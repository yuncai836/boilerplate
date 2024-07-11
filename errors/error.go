package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	inner   error
	message string
	code    Code
}

type Code int

// Msg 创建一个带有信息的新错误
func Msg(msg string) Error {
	return Error{message: msg}
}

func (e Error) Error() string {
	return fmt.Sprintf("(%s) [cause by] (%s)", e.message, e.inner.Error())
}

func (e Error) Walk(f func(err Error)) {
	f(e)
	var v Error
	if errors.As(e.inner, &v) {
		v.Walk(f)
	}
}

// Extend 基于传入的错误扩展该错误
func (e Error) Extend(err error) Error {
	e.inner = err
	return e
}

// Code 携带错误码
func (e Error) Code(c Code) Error {
	e.code = c
	return e
}

func (e Error) GetCode() Code {
	return e.code
}
