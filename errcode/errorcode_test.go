package errcode

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertError(t *testing.T, err error, code int, message string) {
	
	assert.NotNil(t, err)
	assert.Equal(t, code, Code(err))
	assert.Equal(t, fmt.Sprintf("%s", message), err.Error())
}

func TestError_Error(t *testing.T) {
	var err error = &errCode{
		code:  404,
		error: errors.New("not found"),
	}
	assert.Equal(t, "not found", err.Error())


	assertError(t, WithCode(errors.New("internal server error"), 500), 500, "internal server error")
	assertError(t, WithMessage(errors.New("bad request"), 400, "invalid input"), 400, "invalid input: bad request")
	assertError(t, WithMessagef(errors.New("bad request"), 400, "invalid input %s", "details"), 400, "invalid input details: bad request")
	assertError(t, WithStack(errors.New("internal server error"), 500), 500, "internal server error")
	assertError(t, Wrap(errors.New("bad request"), 400, "invalid input"), 400, "invalid input: bad request")
	assertError(t, Wrapf(errors.New("bad request"), 400, "invalid input %s", "details"), 400, "invalid input details: bad request")

}

func TestError_Code(t *testing.T) {
	err := &errCode{
		code: 500,
	}
	assert.Equal(t, 500, err.Code())
}

func TestError_Unwrap(t *testing.T) {
	innerErr := errors.New("inner error")
	err := &errCode{
		error: innerErr,
	}
	assert.Equal(t, innerErr, err.Unwrap())
}

func TestError_Is(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		err := &errCode{code: 400}
		assert.False(t, err.Is(nil))
	})

	t.Run("same error", func(t *testing.T) {
		err := &errCode{code: 400}
		assert.True(t, err.Is(err))
	})

	t.Run("same code", func(t *testing.T) {
		err1 := &errCode{code: 400}
		err2 := &errCode{code: 400}
		assert.True(t, err1.Is(err2))
	})

	t.Run("different code", func(t *testing.T) {
		err1 := &errCode{code: 400}
		err2 := &errCode{code: 500}
		assert.False(t, err1.Is(err2))
	})
}

func TestError_Cause(t *testing.T) {
	innerErr := errors.New("inner error")
	err := &errCode{
		error: innerErr,
	}
	assert.Equal(t, innerErr, err.Cause())
}

func TestError_As(t *testing.T) {

	t.Run("valid target", func(t *testing.T) {
		err := &errCode{error: errors.New("test")}
		var target *errCode
		assert.True(t, err.As(&target))
		assert.Equal(t, err, target)
	})
}

func TestWithStack(t *testing.T) {
	err := errors.New("test")
	wrapped := WithStack(err, 500)
	assert.NotNil(t, wrapped)
}

func TestWrap(t *testing.T) {
	err := errors.New("test")
	wrapped := Wrap(err, 500, "wrapped")
	assert.NotNil(t, wrapped)
}

func TestWrapf(t *testing.T) {
	err := errors.New("test")
	wrapped := Wrapf(err, 500, "wrapped %s", "error")
	assert.NotNil(t, wrapped)

	t.Run("format %v", func(t *testing.T) {
		err := errors.New("test")
		wrapped := Wrapf(err, 500, "wrapped %v", "error")
		assert.Equal(t, "wrapped error: test", wrapped.Error())
	})

	t.Run("format %+v", func(t *testing.T) {
		err := errors.New("test")
		wrapped := Wrapf(err, 500, "wrapped %+v", "error")
		assert.Equal(t, "wrapped error: test", wrapped.Error())
	})

	t.Run("print %v", func(t *testing.T) {
		err := errors.New("test")
		wrapped := Wrapf(err, 500, "wrapped %v", "error")
		assert.Equal(t, "[500] wrapped error: test", fmt.Sprintf("%v", wrapped))
	})

	t.Run("print %+v", func(t *testing.T) {
		err := errors.New("test")
		wrapped := Wrapf(err, 500, "wrapped %+v", "error")
		// 使用 Contains 来检查堆栈信息的关键部分
		assert.Contains(t, fmt.Sprintf("%+v", wrapped), "[500]")
		assert.Contains(t, fmt.Sprintf("%+v", wrapped), "test")
		assert.Contains(t, fmt.Sprintf("%+v", wrapped), "wrapped error")
		assert.Contains(t, fmt.Sprintf("%+v", wrapped), "github.com/zdzh/errorx/errcode.TestWrapf")
	})
}

func TestWithMessage(t *testing.T) {
	err := errors.New("test")
	wrapped := WithMessage(err, 500, "message")
	assert.NotNil(t, wrapped)
}

func TestWithMessagef(t *testing.T) {
	err := errors.New("test")
	wrapped := WithMessagef(err, 500, "message %s", "formatted")
	assert.NotNil(t, wrapped)

	t.Run("format %v", func(t *testing.T) {
		err := errors.New("test")
		wrapped := WithMessagef(err, 500, "message %v", "formatted")
		assert.Equal(t, "message formatted: test", wrapped.Error())
	})

	t.Run("format %+v", func(t *testing.T) {
		err := errors.New("test")
		wrapped := WithMessagef(err, 500, "message %+v", "formatted")
		assert.Equal(t, "message formatted: test", wrapped.Error())
	})
}

func TestCause(t *testing.T) {
	innerErr := errors.New("inner error")
	err := &errCode{
		error: innerErr,
	}
	assert.Equal(t, innerErr, Cause(err))
}

