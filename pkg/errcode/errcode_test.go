package errcode

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/require"
)

func TestCode_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		code     Code
		expected string
	}{
		{name: "unknown", code: CodeUnknown, expected: "Unknown"},
		{name: "invalid argument", code: CodeInvalidArgument, expected: "Invalid argument"},
		{name: "not found", code: CodeNotFound, expected: "Not found"},
		{name: "already exists", code: CodeAlreadyExists, expected: "Already exists"},
		{name: "aborted", code: CodeAborted, expected: "Aborted"},
		{name: "precondition", code: CodePrecondition, expected: "Precondition"},
		{name: "internal", code: CodeInternal, expected: "Internal"},
		{name: "unauthorized", code: CodeUnauthenticated, expected: "Unauthorized"},
		{name: "unImplemented", code: CodeUnimplemented, expected: "UnImplemented"},
		{name: "default", code: 100, expected: "Unknown: 100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.code.String())
		})
	}
}

func TestCode_HTTPStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		code     Code
		expected int
	}{
		{name: "bad request", code: CodeInvalidArgument, expected: http.StatusBadRequest},
		{name: "not found", code: CodeNotFound, expected: http.StatusNotFound},
		{name: "precondition", code: CodePrecondition, expected: http.StatusPreconditionFailed},
		{name: "unauthorized", code: CodeUnauthenticated, expected: http.StatusUnauthorized},
		{name: "conflict", code: CodeAlreadyExists, expected: http.StatusConflict},
		{name: "default", code: CodeUnknown, expected: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.code.HTTPStatus())
		})
	}
}

func TestHTTPStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{name: "err is nil", err: nil, expected: http.StatusOK},
		{name: "success", err: &Error{Code: CodeNotFound}, expected: http.StatusNotFound},
		{name: "success", err: errors.New("error"), expected: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, HTTPStatus(tt.err))
		})
	}
}

func TestHTTPMessage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{name: "err is nil", err: nil, expected: ""},
		{name: "success", err: &Error{origin: errors.New("error")}, expected: "error"},
		{name: "success", err: errors.New("error"), expected: "error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, Message(tt.err))
		})
	}
}

func TestIsNotFound(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{name: "return true", err: &Error{Code: CodeNotFound}, expected: true},
		{name: "return false", err: errors.New("error"), expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, IsNotFound(tt.err))
		})
	}
}

func TestIsAlreadyExists(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{name: "return true", err: &Error{Code: CodeAlreadyExists}, expected: true},
		{name: "return false", err: errors.New("error"), expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, IsAlreadyExists(tt.err))
		})
	}
}

func TestIsInvalidArgument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{name: "return true", err: &Error{Code: CodeInvalidArgument}, expected: true},
		{name: "return false", err: errors.New("error"), expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, IsInvalidArgument(tt.err))
		})
	}
}

func TestError_Error(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      Error
		expected string
	}{
		{
			name:     "origin is nil",
			err:      Error{Code: CodeInternal, stack: "stack"},
			expected: "Internal: StackTrace:\nstack",
		},
		{
			name:     "origin is not nil",
			err:      Error{Code: CodeInvalidArgument, origin: errors.New("error"), stack: "stack"},
			expected: "Invalid argument: error\nStackTrace:\nstack",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

func TestError_Retryable(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      Error
		expected bool
	}{
		{name: "return true", err: Error{Code: CodeInternal}, expected: true},
		{name: "return false", err: Error{Code: CodeInvalidArgument}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.err.Retryable())
		})
	}
}

func TestError_Unwrap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      Error
		expected error
	}{
		{
			name:     "success",
			err:      Error{origin: errors.New("error")},
			expected: errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.err.Unwrap())
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected error
	}{
		{
			name:     "err is nil",
			err:      nil,
			expected: nil,
		},
		{
			name:     "already err type",
			err:      &Error{Code: CodeInternal},
			expected: &Error{Code: CodeInternal},
		},
		{
			name:     "validation error",
			err:      validator.ValidationErrors{},
			expected: &Error{Code: CodeInvalidArgument, origin: validator.ValidationErrors{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := New(tt.err)
			if actual == nil {
				require.Equal(t, tt.expected, actual)
				return
			}
			expected := tt.expected.(*Error)
			require.Equal(t, expected.Code, actual.(*Error).Code)
			require.Equal(t, expected.origin, actual.(*Error).origin)
		})
	}
}

func TestNewCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		code     Code
		expected Code
	}{
		{name: "success", code: CodeNotFound, expected: CodeNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewCode(tt.code).(*Error)
			require.Equal(t, tt.expected, actual.Code)
		})
	}
}

func TestGetCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected Code
	}{
		{
			name:     "err is nil",
			expected: CodeUnknown,
		},
		{
			name:     "success",
			err:      NewCode(CodeNotFound),
			expected: CodeNotFound,
		},
		{
			name:     "unknown",
			err:      errors.New("error"),
			expected: CodeUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, GetCode(tt.err))
		})
	}
}

func TestNewNotFound(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		format   string
		args     interface{}
		expected *Error
	}{
		{
			name:     "success",
			format:   "format: %s",
			args:     "args",
			expected: &Error{Code: CodeNotFound, origin: errors.New("format: args")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewNotFound(tt.format, tt.args).(*Error)
			require.Equal(t, tt.expected.Code, actual.Code)
			require.Equal(t, tt.expected.origin, actual.origin)
		})
	}
}

func TestNewInvalidArgument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		format   string
		args     interface{}
		expected *Error
	}{
		{
			name:     "success",
			format:   "format: %s",
			args:     "args",
			expected: &Error{Code: CodeInvalidArgument, origin: errors.New("format: args")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewInvalidArgument(tt.format, tt.args).(*Error)
			require.Equal(t, tt.expected.Code, actual.Code)
			require.Equal(t, tt.expected.origin, actual.origin)
		})
	}
}

func TestWrapInvalidArgument(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		err      error
		expected *Error
	}{
		{
			name:     "success",
			err:      errors.New("error"),
			expected: &Error{Code: CodeInvalidArgument, origin: errors.New("error")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := WrapInvalidArgument(tt.err).(*Error)
			require.Equal(t, tt.expected.Code, actual.Code)
			require.Equal(t, tt.expected.origin, actual.origin)
		})
	}
}

func TestNewInternal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		format   string
		args     interface{}
		expected *Error
	}{
		{
			name:     "success",
			format:   "format: %s",
			args:     "args",
			expected: &Error{Code: CodeInternal, origin: errors.New("format: args")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewInternal(tt.format, tt.args).(*Error)
			require.Equal(t, tt.expected.Code, actual.Code)
			require.Equal(t, tt.expected.origin, actual.origin)
		})
	}
}

func TestNewPrecondition(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		format   string
		args     interface{}
		expected *Error
	}{
		{
			name:     "success",
			format:   "format: %s",
			args:     "args",
			expected: &Error{Code: CodePrecondition, origin: errors.New("format: args")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewPrecondition(tt.format, tt.args).(*Error)
			require.Equal(t, tt.expected.Code, actual.Code)
			require.Equal(t, tt.expected.origin, actual.origin)
		})
	}
}

func TestNewAlreadyExists(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		format   string
		args     interface{}
		expected *Error
	}{
		{
			name:     "success",
			format:   "format: %s",
			args:     "args",
			expected: &Error{Code: CodeAlreadyExists, origin: errors.New("format: args")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := NewAlreadyExists(tt.format, tt.args).(*Error)
			require.Equal(t, tt.expected.Code, actual.Code)
			require.Equal(t, tt.expected.origin, actual.origin)
		})
	}
}
