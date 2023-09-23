package repository

import (
	"context"

	"github.com/ispec-inc/starry/orion/app"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/adapter"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/entity"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/logger"
	"gorm.io/gorm"
)

var _ organization.Repository = (*Organization)(nil)

// Organization query.Organizationの実装
type Organization struct{}

// Get 指定したIDのOrganizationを取得する。idsが空の場合は全件取得する。
// 意図せず全件取得される可能性があるため、ユースケース層でハンドリングすること。
func (o Organization) Get(ctx context.Context, tx *app.DB, ids []organization.ID) (organization.List, error) {
	ents := []entity.Organization{}
	if err := o.preload(ctx, tx.Get()).Find(&ents, ids).Error; err != nil {
		return organization.List{}, err
	}

	orgs, err := adapter.OrganizationListFromEntityList(ents)
	if err != nil {
		return organization.List{}, err
	}

	return orgs, nil
}

// Create Organizationを作成する
func (o Organization) Create(ctx context.Context, tx *app.DB, org organization.Organization) error {
	ent := adapter.OrganizationToEntity(org)

	return tx.Get().Create(&ent).Error
}

// preload 集約を組み立てるために必要な情報をプリロードする
func (o Organization) preload(ctx context.Context, tx *gorm.DB) *gorm.DB {
	return tx.
		Session(&gorm.Session{Logger: logger.NewZerologToGormLogger(ctx)}).
		Preload("OrganizationDetail")
}
