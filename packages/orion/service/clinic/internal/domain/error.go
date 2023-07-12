package domain

import "errors"

var (
	ErrUnauthorized        = errors.New("unauthorized")          // ErrUnauthorized 認証に失敗したエラー
	ErrStringInvalidLength = errors.New("invalid string length") // ErrStringInvalidLength 文字列の長さが不正なエラー
)
