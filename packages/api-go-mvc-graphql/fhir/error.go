package fhir

import "github.com/pkg/errors"

type Error struct {
	code   ErrorCode
	origin error
	msg    string
}

func WithErrorCode(err error, code ErrorCode) error {
	return errors.WithStack(&Error{
		code:   code,
		origin: err,
		msg:    err.Error(),
	})
}

func New(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeError,
		origin: err,
		msg:    err.Error(),
	})
}

func Invalid(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeInvalid,
		origin: err,
		msg:    err.Error(),
	})
}
func Unauthorized(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeUnauthorized,
		origin: err,
		msg:    err.Error(),
	})
}
func NotFound(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeNotFound,
		origin: err,
		msg:    err.Error(),
	})
}

func Wrap(code ErrorCode, msg string, err error) error {
	err = errors.Wrap(err, msg)
	return errors.WithStack(&Error{
		code:   code,
		origin: err,
		msg:    err.Error(),
	})
}

func Wrapf(err error, msg string, args ...interface{}) error {
	return errors.Wrapf(err, msg, args...)
}

func (e *Error) ErrorCode() ErrorCode {
	return e.code
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Unwrap() error {
	return e.origin
}

func Unwrap(err error) *Error {
	if err == nil {
		return nil
	}

	aerr, ok := errors.Cause(err).(*Error)
	if ok {
		aerr.msg = err.Error()
		return aerr
	}
	return nil
}

const (
	ErrorCodeInvalid      = ErrorCode("invalid")
	ErrorCodeUnauthorized = ErrorCode("unauthorized")
	ErrorCodeNotFound     = ErrorCode("not found")
	ErrorCodeError        = ErrorCode("error")
)

type ErrorCode string

func (e ErrorCode) String() string {
	return string(e)
}
