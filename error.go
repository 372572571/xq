package xq

import (
	"errors"
	"fmt"
)

type CheckedError interface {
	Error() string
	Unwrap() error
	Fe7FbkySMdNx()
}

type checkedError struct {
	err error
}

// 获取错误字符串
func (self *checkedError) Error() string {
	return self.err.Error()
}

// 去除包装 返回原本的错误
func (self *checkedError) Unwrap() error {
	return self.err
}

func (self *checkedError) Fe7FbkySMdNx() {
}

// 错误工具
var Error *ErrorUtil = &ErrorUtil{}

type ErrorUtil struct{}

// 创建
func (self *ErrorUtil) New(format string, a ...Any) error {
	return fmt.Errorf(format, a...)
}
// 包装
func (self *ErrorUtil) Wrap(err error) (wrapper CheckedError) {
	if errors.As(err, &wrapper) {
		return wrapper
	}
	return &checkedError{err}
}

func (self *ErrorUtil) Catch(err *error) error {
	return self.Unwrap(err, recover())
}

// 解包
func (self *ErrorUtil) Unwrap(err *error, recovered interface{}) error {
	if err == nil {
		var e error
		err = &e
	}
	if recovered == nil {
		return *err
	}

	var checkedError CheckedError
	if !errors.As(recovered.(error), &checkedError) {
		panic(recovered)
	}

	*err = checkedError.Unwrap()
	return *err
}
