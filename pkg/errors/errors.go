package errors

import (
	"errors"
	"fmt"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 10001
	// UnknownReason is unknown reason for error info.
	UnknownReason = "UNKNOWN_ERROR"
	// UnknownMessage is unknown reason for error info.
	UnknownMessage = "未知错误"
	// SupportPackageIsVersion1 this constant should not be referenced by any other code.
	SupportPackageIsVersion1 = true
)

type Error struct {
	Code    int32  `json:"code,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
	cause   error
}

func (e *Error) Error() string {
	return fmt.Sprintf("error: code = %d reason = %s message = %s cause = %+v", e.Code, e.Reason, e.Message, e.cause)
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Code == e.Code && se.Reason == e.Reason
	}
	return false
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (e *Error) Unwrap() error { return e.cause }

// WithCause with the underlying cause of the error.
func (e *Error) WithCause(cause error) *Error {
	err := Clone(e)
	err.cause = cause
	return err
}

// New returns an error object for the code, message.
func New(code int, reason, message string) *Error {
	return &Error{
		Code:    int32(code),
		Reason:  reason,
		Message: message,
	}
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Errorf returns an error object for the code, message and error info.
func Errorf(code int, reason, format string, a ...interface{}) error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// Code returns the http code for an error.
// It supports wrapped errors.
func Code(err error) int {
	if err == nil {
		return 200 //nolint:gomnd
	}
	return int(FromError(err).Code)
}

// Message returns the message for a particular error.
// It supports wrapped errors.
func Message(err error) string {
	if err == nil {
		return UnknownMessage
	}
	return FromError(err).Message
}

// Reason returns the reason for a particular error.
// It supports wrapped errors.
func Reason(err error) string {
	if err == nil {
		return UnknownReason
	}
	return FromError(err).Reason
}

// Clone deep clone error to a new error.
func Clone(err *Error) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		cause:   err.cause,
		Code:    err.Code,
		Reason:  err.Reason,
		Message: err.Message,
	}
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	return New(UnknownCode, err.Error(), UnknownMessage)
}
