package reader

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/query"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/adapter"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/logger"
	"gorm.io/gorm"
)

var _ query.Organization = (*Organization)(nil)

// Organization query.Organizationの実装
type Organization struct {
	// MySQL DB
	db *gorm.DB
}

// NewOrganization Organizationのコンストラクタ
func NewOrganization(db *gorm.DB) Organization {
	return Organization{db: db}
}

// Get 指定したIDのOrganizationを取得する
// idsが空の場合は全件取得する
func (o Organization) Get(ctx context.Context, ids []model.ID) ([]model.Organization, error) {
	orgs := []entity.Organization{}
	if err := o.preload(ctx).Find(&orgs, ids).Error; err != nil {
		return []model.Organization{}, app.NotFound(err)
	}

	return adapter.OrganizationListFromEntityList(orgs), nil
}

func (c Organization) preload(ctx context.Context) *gorm.DB {
	return c.db.
		Session(&gorm.Session{Logger: logger.NewZerologToGormLogger(ctx)}).
		Preload("OrganizationDetail")
}
