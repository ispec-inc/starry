package reader

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/infra/adapter"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/infra/entity"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/infra/logger"
	"gorm.io/gorm"
)

var _ organization.Query = (*Organization)(nil)

// Organization query.Organizationの実装
type Organization struct {
	// MySQL DB
}

// NewOrganization Organizationのコンストラクタ
func NewOrganization() Organization {
	return Organization{}
}

// Get 指定したIDのOrganizationを取得する
// idsが空の場合は全件取得する
func (o Organization) Get(ctx context.Context, tx *app.DB, ids []domain.ID) ([]organization.Organization, error) {
	orgs := []entity.Organization{}
	if err := o.preload(ctx, tx.Get()).Find(&orgs, ids).Error; err != nil {
		return []organization.Organization{}, err
	}

	return adapter.OrganizationListFromEntityList(orgs), nil
}

func (o Organization) preload(ctx context.Context, tx *gorm.DB) *gorm.DB {
	return tx.
		Session(&gorm.Session{Logger: logger.NewZerologToGormLogger(ctx)}).
		Preload("OrganizationDetail")
}
