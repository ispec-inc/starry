package controller

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/resolver"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/uc"
)

type OrganizationArgs struct {
	ID graphql.ID
}

func (c Controller) Organization(ctx context.Context, args OrganizationArgs) (resolver.Organization, error) {

	ipt := uc.GetOrganizationInput{
		ID: model.ID(args.ID),
	}
	get := uc.NewGetOrganization(c.registry)

	opt, err := get.Do(ctx, ipt)
	if err != nil {
		return resolver.Organization{}, err
	}

	r := resolver.NewOrganization(opt.Organization)
	return r, nil
}
