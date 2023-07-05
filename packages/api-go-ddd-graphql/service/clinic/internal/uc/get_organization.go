package uc

import (
	"context"
	"errors"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pubsub"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/registry"
)

type GetOrganization struct {
	db                *app.DB
	organizationQuery organization.Query
}

type GetOrganizationInput struct {
	ID domain.ID
}

type GetOrganizationOutput struct {
	Organization organization.Organization
}

func NewGetOrganization(r registry.Registry) GetOrganization {
	return GetOrganization{
		db:                r.Repository().DB(),
		organizationQuery: r.Repository().NewOrganizationQuery(),
	}
}

func (g GetOrganization) Do(ctx context.Context, ipt GetOrganizationInput) (GetOrganizationOutput, error) {

	tx := g.db.Begin()

	defer tx.Rollback()

	os, err := g.organizationQuery.Get(ctx, tx, []domain.ID{ipt.ID})
	if err != nil {
		return GetOrganizationOutput{}, err
	}

	if len(os) == 0 {
		return GetOrganizationOutput{}, errors.New("organization not found")
	}

	if err := pubsub.Publish(ctx, pubsub.DoneKey, pubsub.Done{
		ID: 1,
	}); err != nil {
		return GetOrganizationOutput{}, errors.New("organization not found")
	}

	if err := tx.Commit(); err != nil {
		return GetOrganizationOutput{}, errors.New("organization not found")
	}

	return GetOrganizationOutput{Organization: os[0]}, nil
}
