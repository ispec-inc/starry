package model

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
)

const (
	maxOrganizationNameChars        int = 50
	maxOrganizationAliasChars       int = 50
	maxOrganizationDescriptionChars int = 50
)

type Organization struct {
	ID          ID
	Name        OrganizationName
	Alias       OrganizationAlias
	Type        OrganizationType
	Contact     PhoneNumber
	Description OrganizationDescription
}

func RegisterOrganization(
	name OrganizationName,
	alias OrganizationAlias,
	otype OrganizationType,
	contact PhoneNumber,
	description OrganizationDescription,
) (Organization, error) {
	id, err := NewID("organization")
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

func (c Organization) Validate() error {
	if err := c.Name.Validate(); err != nil {
		return err
	}

	if err := c.Alias.Validate(); err != nil {
		return err
	}

	if err := c.Contact.Validate(); err != nil {
		return err
	}

	if err := c.Description.Validate(); err != nil {
		return err
	}

	return nil
}

type OrganizationType int

var (
	OrganizationTypeProv int32 = 1
	OrganizationTypeOrg  int32 = 2
)

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

func NewOrganizationType(v string) int32 {
	switch v {
	case "PROV":
		return OrganizationTypeProv
	case "ORG":
		return OrganizationTypeOrg
	default:
		return 0
	}
}

type OrganizationName string

func (c OrganizationName) Validate() error {
	if c == "" {
		return apperror.New(errors.New("organization: name is empty"))
	}

	if utf8.RuneCountInString(string(c)) > maxOrganizationNameChars {
		err := fmt.Errorf("organization: name exceeds %d characters", maxOrganizationNameChars)
		return apperror.New(err)
	}

	return nil
}

type OrganizationAlias string

func (c OrganizationAlias) Validate() error {
	if c == "" {
		return apperror.New(errors.New("organization: alias is empty"))
	}

	if utf8.RuneCountInString(string(c)) > maxOrganizationAliasChars {
		err := fmt.Errorf("organization: alias exceeds %d characters", maxOrganizationAliasChars)
		return apperror.New(err)
	}

	return nil
}

type OrganizationDescription string

func (c OrganizationDescription) Validate() error {
	if c == "" {
		return apperror.New(errors.New("organization: description is empty"))
	}

	if utf8.RuneCountInString(string(c)) > maxOrganizationDescriptionChars {
		err := fmt.Errorf("organization: description exceeds %d characters", maxOrganizationDescriptionChars)
		return apperror.New(err)
	}

	return nil
}
