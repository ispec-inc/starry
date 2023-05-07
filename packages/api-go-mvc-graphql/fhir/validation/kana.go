package validation

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func ValidateKana(fl validator.FieldLevel) bool {
	for _, r := range []rune(fl.Field().String()) {
		if !unicode.In(r, unicode.Hiragana, unicode.Katakana) {
			return false
		}
	}

	return true
}
