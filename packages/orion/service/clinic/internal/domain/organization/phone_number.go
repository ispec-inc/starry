package organization

import (
	"regexp"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
)

// PhoneNumber 電話番号の値オブジェクト
type PhoneNumber domain.String

// NewPhoneNumber 連絡先を初期化する
func NewPhoneNumber(
	p string,
) (PhoneNumber, error) {
	re := regexp.MustCompile(`^[0-9]{10,11}$`)
	if re.MatchString(string(p)) {
		return "", nil
	}

	return PhoneNumber(p), nil
}

func (p PhoneNumber) String() string {
	return string(p)
}
