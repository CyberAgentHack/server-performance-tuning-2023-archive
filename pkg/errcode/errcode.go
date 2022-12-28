package errcode

import (
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/go-playground/validator"
)

type Code int

const (
	CodeUnknown Code = iota
	CodeInvalidArgument
	CodeNotFound
	CodeAlreadyExists
	CodeAborted
	CodePrecondition
	CodeInternal
	CodeUnauthenticated
	CodeUnimplemented
)

func (c Code) String() string {
	switch c {
	case CodeUnknown:
		return "Unknown"
	case CodeInvalidArgument:
		return "Invalid argument"
	case CodeNotFound:
		return "Not found"
	case CodeAlreadyExists:
		return "Already exists"
	case CodeAborted:
		return "Aborted"
	case CodePrecondition:
		return "Precondition"
	case CodeInternal:
		return "Internal"
	case CodeUnauthenticated:
		return "Unauthorized"
	case CodeUnimplemented:
		return "UnImplemented"
	}
	return fmt.Sprintf("Unknown: %d", c)
}

func (c Code) HTTPStatus() int {
	switch c {
	case CodeInvalidArgument:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	case CodePrecondition:
		return http.StatusPreconditionFailed
	case CodeUnauthenticated:
		return http.StatusUnauthorized
	case CodeAlreadyExists:
		return http.StatusConflict
	}
	return http.StatusInternalServerError
}

func HTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	var cerr *Error
	if errors.As(err, &cerr) {
		return cerr.Code.HTTPStatus()
	}
	return http.StatusInternalServerError
}

func Message(err error) string {
	if err == nil {
		return ""
	}
	var cerr *Error
	if errors.As(err, &cerr) {
		return cerr.origin.Error()
	}
	return err.Error()
}

func IsNotFound(err error) bool {
	var cerr *Error
	if errors.As(err, &cerr) {
		return cerr.Code == CodeNotFound
	}
	return false
}

func IsAlreadyExists(err error) bool {
	var cerr *Error
	if errors.As(err, &cerr) {
		return cerr.Code == CodeAlreadyExists
	}
	return false
}

func IsInvalidArgument(err error) bool {
	var cerr *Error
	if errors.As(err, &cerr) {
		return cerr.Code == CodeInvalidArgument
	}
	return false
}

type Error struct {
	Code   Code
	origin error
	stack  string
}

func (e *Error) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("%s: StackTrace:\n%s", e.Code.String(), e.stack)
	}
	return fmt.Sprintf("%s: %s\nStackTrace:\n%s", e.Code.String(), e.origin.Error(), e.stack)
}

func (e *Error) Retryable() bool {
	switch e.Code {
	case CodeInternal, CodeUnknown, CodeAborted:
		return true
	}
	return false
}

func (e *Error) Unwrap() error {
	return e.origin
}

func New(err error) error {
	if err == nil {
		return nil
	}
	// if err is already Error type, nothing.
	var e *Error
	if errors.As(err, &e) {
		return err
	}

	newerr := &Error{
		Code:   CodeInternal,
		origin: err,
		stack:  string(debug.Stack()),
	}
	// check validation error
	var verr validator.ValidationErrors
	if errors.As(err, &verr) {
		newerr.Code = CodeInvalidArgument
		return newerr
	}

	return newerr
}

func NewCode(code Code) error {
	stack := debug.Stack()
	return &Error{
		Code:  code,
		stack: string(stack),
	}
}

func GetCode(err error) Code {
	if err == nil {
		return CodeUnknown
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return CodeUnknown
}

func NewNotFound(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeNotFound,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewInvalidArgument(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeInvalidArgument,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func WrapInvalidArgument(err error) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeInvalidArgument,
		origin: err,
		stack:  string(stack),
	}
}

func NewInternal(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeInternal,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewPrecondition(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodePrecondition,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewAlreadyExists(format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   CodeAlreadyExists,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}
