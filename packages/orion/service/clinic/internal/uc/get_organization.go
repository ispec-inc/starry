package uc

import (
	"context"

	"github.com/ispec-inc/starry/orion/app"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/registry"
)

// GetOrganization Organizationを取得するユースケース
type GetOrganization struct {
	db            *app.DB
	orgRepository organization.Repository
}

// GetOrganizationInput Organizationを取得するユースケースの入力
type GetOrganizationInput struct {
	ID organization.ID
}

// GetOrganizationOutput Organizationを取得するユースケースの出力
type GetOrganizationOutput struct {
	Organization organization.Organization
}

// NewGetOrganization GetOrganizationのコンストラクタ
func NewGetOrganization(r registry.Registry) GetOrganization {
	return GetOrganization{
		db:            r.Repository().DB(),
		orgRepository: r.Repository().NewOrganizationRepository(),
	}
}

// Do GetOrganizationのユースケースの実行
func (g GetOrganization) Do(ctx context.Context, ipt GetOrganizationInput) (GetOrganizationOutput, error) {

	tx := g.db.Begin()

	defer tx.Rollback()

	orgs, err := g.orgRepository.Get(ctx, tx, []organization.ID{ipt.ID})
	if err != nil {
		return GetOrganizationOutput{}, app.ErrUnauthorized
	}

	org, err := orgs.First()
	if err != nil {
		return GetOrganizationOutput{}, err
	}

	if err := tx.Commit(); err != nil {
		return GetOrganizationOutput{}, err
	}

	return GetOrganizationOutput{Organization: org}, nil
}
