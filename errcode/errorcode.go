package errcode

import (
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// ErrorCode 错误码类型
type ErrorCode = int

type errCode struct {
	code ErrorCode
	error
}

func (e *errCode) Error() string {
	return fmt.Sprintf("[%d] %s", e.code, e.error.Error())
}

// Code 返回错误码
func (e *errCode) Code() ErrorCode {
	return e.code
}

// Unwrap 返回原始错误，用于错误链的解包
func (e *errCode) Unwrap() error {
	return e.error
}

// Is 判断当前错误是否与目标错误相等或具有相同的错误码
func (e *errCode) Is(err error) bool {
	if err == nil {
		return false
	}
	if e == err {
		return true
	}

	if errWithCode, ok := err.(interface{ Code() ErrorCode }); ok {
		return e.code == errWithCode.Code()
	}
	return false
}

func (w *errCode) Cause() error { return w.error }

// As 将当前错误转换为目标类型
func (e *errCode) As(target interface{}) bool {
	if target == nil {
		return false
	}
	return errors.As(e, target)
}

func (w *errCode) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "[%d] %+v",w.Code(), w.Cause())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}



// WithCode 内部函数，用于创建带有错误码的Error实例
func WithCode(err error, code ErrorCode) *errCode {
	return &errCode{code: code, error: err}
}

// WithStack 创建带有堆栈信息和错误码的错误
func WithStack(err error, code ErrorCode) error {
	return errors.WithStack(WithCode(err, code))
}

// Wrap 包装错误并添加错误码和消息
func Wrap(err error, code ErrorCode, message string) error {
	return WithCode(errors.Wrap(err, message), code)
}

// Wrapf 包装错误并添加错误码和格式化消息
func Wrapf(err error, code ErrorCode, format string, args ...interface{}) error {
	return WithCode(errors.Wrapf(err, format, args...), code)
}

// WithMessage 添加错误码和消息，但不包装错误
func WithMessage(err error, code ErrorCode, message string) error {
	return WithCode(errors.WithMessage(err, message), code)
}

// WithMessagef 添加错误码和格式化消息，但不包装错误
func WithMessagef(err error, code ErrorCode, format string, args ...interface{}) error {
	return WithCode(errors.WithMessagef(err, format, args...), code)
}

// Cause 返回错误的根本原因
func Cause(err error) error {
	return errors.Cause(err)
}

// 返回错误码
var defaultErrorCode ErrorCode = -1

// SetDefaultCode 允许设置全局默认错误码
func SetDefaultCode(code ErrorCode) {
	defaultErrorCode = code
}

// Code 返回错误的错误码，如果没有则返回全局默认错误码
func Code(err error) ErrorCode {
	if err == nil {
		return defaultErrorCode
	}
	//
	for err != nil {
		if errWithCode, ok := err.(interface{ Code() ErrorCode }); ok {
			return errWithCode.Code()
		}
		// 尝试解包
		type unwrapper interface{ Unwrap() error }
		if uw, ok := err.(unwrapper); ok {
			err = uw.Unwrap()
		} else {
			break
		}
	}

	return defaultErrorCode
}
