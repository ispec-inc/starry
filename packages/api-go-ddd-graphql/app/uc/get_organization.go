package uc

import (
	"context"
	"errors"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/query"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/registry"
)

type GetOrganization struct {
	organizationQuery query.Organization
}

type GetOrganizationInput struct {
	ID model.ID
}

type GetOrganizationOutput struct {
	Organization model.Organization
}

func NewGetOrganization(r registry.Registry) GetOrganization {
	return GetOrganization{
		organizationQuery: r.Repository().NewOrganizationQuery(),
	}
}

func (g GetOrganization) Do(ctx context.Context, ipt GetOrganizationInput) (GetOrganizationOutput, error) {
	os, err := g.organizationQuery.Get(ctx, []model.ID{ipt.ID})
	if err != nil {
		return GetOrganizationOutput{}, err
	}

	if len(os) == 0 {
		return GetOrganizationOutput{}, app.NotFound(errors.New("organization not found"))
	}

	return GetOrganizationOutput{Organization: os[0]}, nil
}
