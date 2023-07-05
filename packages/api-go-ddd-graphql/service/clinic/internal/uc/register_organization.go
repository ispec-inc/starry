package uc

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/registry"
)

type RegisterOrganization struct {
	organizationQuery organization.Query
}

type RegisterOrganizationInput struct {
	ID domain.ID
}

type RegisterOrganizationOutput struct {
	Organization organization.Organization
}

func NewRegisterOrganization(r registry.Registry) RegisterOrganization {
	return RegisterOrganization{
		organizationQuery: r.Repository().NewOrganizationQuery(),
	}
}

func (g RegisterOrganization) Do(ctx context.Context, ipt RegisterOrganizationInput) (RegisterOrganizationOutput, error) {
	// TODO 実装
	return RegisterOrganizationOutput{}, nil
}
