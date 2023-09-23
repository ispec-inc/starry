package organization

import "errors"

var (
	ErrNotFound                 = errors.New("organization: not found")      // ErrNotFound 組織が見つからない
	ErrPhoneNumberInvalidFormat = errors.New("phone_number: invalid format") // ErrPhoneNumberInvalidFormat 電話番号のフォーマットが不正
	ErrTypeStringIsInvalid      = errors.New("type: string is invalid")      // ErrTypeStringIsInvalid 組織の種別の文字列が不正
)
