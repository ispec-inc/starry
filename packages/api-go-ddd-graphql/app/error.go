package app

import "github.com/pkg/errors"

// Error ハンドリングがしやすいエラー
type Error struct {
	code   ErrorCode
	origin error
	msg    string
}

// WithErrorCode エラーコードを付与する
func WithErrorCode(err error, code ErrorCode) error {
	return errors.WithStack(&Error{
		code:   code,
		origin: err,
		msg:    err.Error(),
	})
}

// NewError エラーを作成する
func NewError(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeError,
		origin: err,
		msg:    err.Error(),
	})
}

// Invalid バリデーションエラーを作成する
func Invalid(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeInvalid,
		origin: err,
		msg:    err.Error(),
	})
}

// Unauthorized 認証エラーを作成する
func Unauthorized(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeUnauthorized,
		origin: err,
		msg:    err.Error(),
	})
}

// NotFound データみつからなかった際のエラーを作成する
func NotFound(err error) error {
	return errors.WithStack(&Error{
		code:   ErrorCodeNotFound,
		origin: err,
		msg:    err.Error(),
	})
}

// Wrap エラーをラップする
func Wrap(code ErrorCode, msg string, err error) error {
	err = errors.Wrap(err, msg)
	return errors.WithStack(&Error{
		code:   code,
		origin: err,
		msg:    err.Error(),
	})
}

// Wrapf エラーをラップする
func Wrapf(err error, msg string, args ...interface{}) error {
	return errors.Wrapf(err, msg, args...)
}

// ErrorCoder エラーコードを返す
func (e *Error) ErrorCode() ErrorCode {
	return e.code
}

// Error エラーメッセージを返す
func (e *Error) Error() string {
	return e.msg
}

// Unwrap 生のエラーを返す
func (e *Error) Unwrap() error {
	return e.origin
}

// Unwrap 生のエラーを返す
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
	// ErrorCodeInvalid バリデーションエラー
	ErrorCodeInvalid = ErrorCode("invalid")

	// ErrorCodeUnauthorized 認証エラー
	ErrorCodeUnauthorized = ErrorCode("unauthorized")

	// ErrorCodeNotFound データが見つからないエラー
	ErrorCodeNotFound = ErrorCode("not found")

	// ErrorCodeError その他のエラー
	ErrorCodeError = ErrorCode("error")
)

// ErrorCode エラーコード
type ErrorCode string

// String エラーコードを文字列で返す
func (e ErrorCode) String() string {
	return string(e)
}
