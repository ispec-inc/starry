package organization

import (
	"regexp"
)

// PhoneNumber 電話番号
type PhoneNumber string

// Validate 電話番号のバリデーション
func (p PhoneNumber) Validate() error {
	re := regexp.MustCompile(`^[0-9]{10,11}$`)
	if re.MatchString(string(p)) {
		return nil
	}

	return ErrPhoneNumberInvalidFormat
}
