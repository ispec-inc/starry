package organization

import (
	"errors"
	"regexp"
)

type PhoneNumber string

func (p PhoneNumber) Validate() error {
	re := regexp.MustCompile(`^[0-9]{10,11}$`)
	if re.MatchString(string(p)) {
		return nil
	}

	return errors.New("phone_number: invalid format")
}
