package uc

import (
	"context"
	"errors"

	"github.com/ispec-inc/starry/orion/app"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/registry"
)

// GetOrganization Organizationを取得するユースケース
type GetOrganization struct {
	db                *app.DB
	organizationQuery organization.Query
}

// GetOrganizationInput Organizationを取得するユースケースの入力
type GetOrganizationInput struct {
	ID domain.ID
}

// GetOrganizationOutput Organizationを取得するユースケースの出力
type GetOrganizationOutput struct {
	Organization organization.Organization
}

// NewGetOrganization GetOrganizationのコンストラクタ
func NewGetOrganization(r registry.Registry) GetOrganization {
	return GetOrganization{
		db:                r.Repository().DB(),
		organizationQuery: r.Repository().NewOrganizationQuery(),
	}
}

// Do GetOrganizationのユースケースの実行
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

	if err := tx.Commit(); err != nil {
		return GetOrganizationOutput{}, errors.New("organization not found")
	}

	return GetOrganizationOutput{Organization: os[0]}, nil
}
