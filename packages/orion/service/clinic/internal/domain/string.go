package domain

import (
	"fmt"
	"unicode/utf8"
)

// String ドメインモデルの中で共通で使う文字列の値オブジェクト
// 文字数のバリデーションを共通化する
type String string

func (s String) ValidateLength(min, max int) error {
	length := utf8.RuneCountInString(string(s))
	if length < min || length > max {
		return fmt.Errorf("%w: got %d, expected between %d and %d", ErrStringInvalidLength, length, min, max)
	}

	return nil
}

func (s String) String() string {
	return string(s)
}
