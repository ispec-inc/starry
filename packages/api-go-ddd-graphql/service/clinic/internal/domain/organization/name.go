package organization

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Name 組織名
type Name string

// Validate 組織名のバリデーション
func (n Name) Validate() error {
	if n == "" {
		return errors.New("organization: name is empty")
	}

	if utf8.RuneCountInString(string(n)) > maxOrganizationNameChars {
		err := fmt.Errorf("organization: name exceeds %d characters", maxOrganizationNameChars)
		return err
	}

	return nil
}
