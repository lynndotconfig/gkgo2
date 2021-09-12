package gkgo2err

import (
	"fmt"
	"runtime"
	"strings"
)

type (
	Error struct {
		codeMsg    *Code
		cause      error
		InvokeLink *invokeNode
		Location   string
	}
	invokeNode struct {
		Name   string
		Param  interface{}
		Line   int
		ErrMsg string
		Next   *invokeNode
	}
)

func (p *Error) Error() string {
	return fmt.Sprintf("celestial celestialErr, location: %s, celestialErr: %s", p.Location, p.cause.Error())
}

func (p *Error) CodeMsg() *Code {
	return p.codeMsg
}

// 完全新建一个错误，用于偏底层操作错误
func New(cm *Code, param interface{}, format string, a ...interface{}) *Error {
	_, filename, line, _ := runtime.Caller(1)
	err := fmt.Errorf(format, a...)
	return &Error{
		codeMsg:    cm,
		cause:      err,
		InvokeLink: NewInvokeNode(param, err.Error()),
		Location:   fmt.Sprintf("file %s, line %v", filename, line),
	}
}

// 以当前错误为基础生成一个错误，用于偏底层操作错误
func With(cm *Code, param interface{}, err error) *Error {
	if err == nil {
		return nil
	}
	_, filename, line, _ := runtime.Caller(1)
	return &Error{
		codeMsg:    cm,
		cause:      err,
		InvokeLink: NewInvokeNode(param, err.Error()),
		Location:   fmt.Sprintf("file %s, line %v",filename, line),
	}
}

// 包括一个错误，可以构建调用链，用于偏上层操作错误
func Wrap(err *Error, param interface{}, cm *Code) *Error {
	_, filename, line, _ := runtime.Caller(1)
	error := Error{
		codeMsg:    cm,
		cause:      err.cause,
		InvokeLink: NewInvokeNode(param, err.cause.Error()),
		Location:   fmt.Sprintf("file %s, line %v", filename, line),
	}
	error.InvokeLink.Next = err.InvokeLink
	return &error
}

func NewInvokeNode(param interface{}, errMsg string) *invokeNode {
	pc, _, line, ok := runtime.Caller(2)
	if !ok {
		return &invokeNode{
			Name:   "XXX",
			Param:  param,
			Line:   0,
			ErrMsg: "",
			Next:   nil,
		}
	}
	funcName := runtime.FuncForPC(pc).Name()
	pos := strings.LastIndex(funcName, ".")
	funcName = funcName[pos+1:]
	return &invokeNode{
		Name:   funcName,
		Param:  param,
		Line:   line,
		ErrMsg: errMsg,
		Next:   nil,
	}
}

