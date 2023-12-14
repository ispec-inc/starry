package organization

import (
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
)

// Organization 組織を表現する集約
type Organization struct {
	// ID 組織のID
	ID ID
	// Name 名前
	Name Name
	// Type 種類
	Type Type
	// PhoneNumber 電話番号
	PhoneNumber PhoneNumber
}

// ID 組織のID
type ID domain.ID

// New 組織の集約を初期化するファクトリ
func New(
	name Name,
	otype Type,
	phoneNumber PhoneNumber,
) (Organization, error) {
	id := domain.NewID()

	o := Organization{
		ID:          ID(id),
		Name:        name,
		Type:        otype,
		PhoneNumber: phoneNumber,
	}

	return o, nil
}

// List Organizationの配列
type List []Organization

// First 配列の先頭のOrganizationを返し、 配列のサイズが0の場合はエラーを返す
func (l List) First() (Organization, error) {
	if len(l) == 0 {
		return Organization{}, ErrNotFound
	}

	return l[0], nil
}
