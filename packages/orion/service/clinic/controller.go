package clinic

import (
	"context"
	"errors"

	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/orion/app/gqlerror"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/registry"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/resolver"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/uc"
)

// Controller Organizationのコントローラ
type Controller struct {
	registry registry.Registry
}

// New Controllerのコンストラクタ
func New() (*Controller, error) {
	rgst, err := registry.New()
	if err != nil {
		return nil, err
	}

	return &Controller{
		registry: rgst,
	}, nil
}

// Organization IDに該当するOrganizationを取得する
func (c Controller) Organization(ctx context.Context, args struct {
	ID graphql.ID
}) (resolver.Organization, error) {

	ipt := uc.GetOrganizationInput{
		ID: organization.ID(args.ID),
	}

	get := uc.NewGetOrganization(c.registry)

	opt, err := get.Do(ctx, ipt)
	if err != nil {

		if errors.Is(err, domain.ErrStringInvalidLength) {
			return resolver.Organization{}, gqlerror.NewWithCode(ctx, err, "invalid", "文字列の長さが不正です")
		}
		return resolver.Organization{}, gqlerror.New(ctx, err)
	}

	r := resolver.Organization{Model: opt.Organization}
	return r, nil
}
