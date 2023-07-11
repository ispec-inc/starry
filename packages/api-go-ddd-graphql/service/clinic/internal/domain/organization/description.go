package organization

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Description 組織の説明
type Description string

// Validate 組織の説明のバリデーション
func (d Description) Validate() error {
	if d == "" {
		return errors.New("organization: description is empty")
	}

	if utf8.RuneCountInString(string(d)) > maxOrganizationDescriptionChars {
		err := fmt.Errorf("organization: description exceeds %d characters", maxOrganizationDescriptionChars)
		return err
	}

	return nil
}
