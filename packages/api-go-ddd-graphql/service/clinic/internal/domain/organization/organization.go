package organization

import (
	"errors"

	"fmt"
	"unicode/utf8"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
)

const (
	maxOrganizationNameChars        int = 50
	maxOrganizationAliasChars       int = 50
	maxOrganizationDescriptionChars int = 50
)

// Organization 組織を表現するドメインモデルの集約
type Organization struct {
	ID          domain.ID
	Name        Name
	Alias       Alias
	Type        OrganizationType
	Contact     PhoneNumber
	Description Description
}

// RegisterOrganization 組織を登録するドメインサービス
func RegisterOrganization(
	name Name,
	alias Alias,
	otype OrganizationType,
	contact PhoneNumber,
	description Description,
) (Organization, error) {
	id, err := domain.NewID()
	if err != nil {
		return Organization{}, err

	}

	o := Organization{
		ID:          id,
		Name:        name,
		Alias:       alias,
		Type:        otype,
		Contact:     contact,
		Description: description,
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

	if err := o.Contact.Validate(); err != nil {
		return err
	}

	if err := o.Description.Validate(); err != nil {
		return err
	}

	return nil
}

// OrganizationType 組織の種別
type OrganizationType int

var (
	OrganizationTypeProv int32 = 1
	OrganizationTypeOrg  int32 = 2
)

// OrganizationTypeString 組織の種別を文字列に変換する
func OrganizationTypeString(v int32) string {
	switch v {
	case OrganizationTypeProv:
		return "PROV"
	case OrganizationTypeOrg:
		return "ORG"
	default:
		return ""
	}
}

// NewType 組織の種別を文字列から生成する
func NewType(v string) int32 {
	switch v {
	case "PROV":
		return OrganizationTypeProv
	case "ORG":
		return OrganizationTypeOrg
	default:
		return 0
	}
}

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
