package organization

import (
	"fmt"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
)

// Organization 組織を表現するドメインモデルの集約
type Organization struct {
	// ID 組織のID
	ID ID
	// Name 名前
	Name Name
	// Alias 別名
	Alias Alias
	// Type 種類
	Type Type
	// Contact 連絡先
	Contact PhoneNumber
}

// ID 組織のID
type ID domain.ID

// RegisterOrganization 組織を登録するドメインサービス
func RegisterOrganization(
	name Name,
	alias Alias,
	otype Type,
	contact PhoneNumber,
) (Organization, error) {
	id := domain.NewID()

	o := Organization{
		ID:      ID(id),
		Name:    name,
		Alias:   alias,
		Type:    otype,
		Contact: contact,
	}

	if err := o.Validate(); err != nil {
		return Organization{}, fmt.Errorf("RegisterOrganization: %w", err)
	}

	return o, nil
}

// Validate 組織のドメインモデルのバリデーション
func (o Organization) Validate() error {
	if err := o.Name.Validate(); err != nil {
		return err
	}

	if err := o.Alias.Validate(); err != nil {
		return err
	}

	return o.Contact.Validate()
}
