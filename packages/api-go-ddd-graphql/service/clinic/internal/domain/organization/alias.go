package organization

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// Alias 組織の略称
type Alias string

// Validate 組織の略称のバリデーション
func (a Alias) Validate() error {
	if a == "" {
		return errors.New("organization: alias is empty")
	}

	if utf8.RuneCountInString(string(a)) > maxOrganizationAliasChars {
		err := fmt.Errorf("organization: alias exceeds %d characters", maxOrganizationAliasChars)
		return err
	}

	return nil
}
