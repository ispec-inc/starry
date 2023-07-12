package organization

import "errors"

var (
	ErrPhoneNumberInvalidFormat = errors.New("phone_number: invalid format") // ErrPhoneNumberInvalidFormat 電話番号のフォーマットが不正
)
