package model

import (
	"errors"
	"regexp"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
)

type PhoneNumber string

func (p PhoneNumber) Validate() error {
	re := regexp.MustCompile(`^[0-9]{10,11}$`)
	if re.MatchString(string(p)) {
		return nil
	}

	return app.Invalid(errors.New("phone_number: invalid format"))
}
