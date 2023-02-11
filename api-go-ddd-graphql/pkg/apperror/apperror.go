package apperror

import (
	"github.com/pkg/errors"
)

type Error struct {
	code   Code
	origin error
	msg    string
}

func WithCode(err error, code Code) error {
	return errors.WithStack(&Error{
		code:   code,
		origin: err,
		msg:    err.Error(),
	})
}

func New(err error) error {
	return errors.WithStack(&Error{
		code:   CodeError,
		origin: err,
		msg:    err.Error(),
	})
}

func Invalid(err error) error {
	return errors.WithStack(&Error{
		code:   CodeInvalid,
		origin: err,
		msg:    err.Error(),
	})
}
func Unauthorized(err error) error {
	return errors.WithStack(&Error{
		code:   CodeUnauthorized,
		origin: err,
		msg:    err.Error(),
	})
}
func NotFound(err error) error {
	return errors.WithStack(&Error{
		code:   CodeNotFound,
		origin: err,
		msg:    err.Error(),
	})
}

func Wrap(code Code, msg string, err error) error {
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

func (e *Error) Code() Code {
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
