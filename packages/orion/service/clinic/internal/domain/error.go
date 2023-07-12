package domain

import "errors"

var (
	ErrUnauthorized        = errors.New("unauthorized") // ErrUnauthorized 認証に失敗したエラー
	ErrStringInvalidLength = errors.New("string: invalid length")
)
