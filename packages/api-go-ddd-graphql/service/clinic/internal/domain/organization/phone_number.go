package organization

import (
	"errors"
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

	return errors.New("phone_number: invalid format")
}
