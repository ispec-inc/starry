package reader

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/query"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/adapter"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/logger"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
	"gorm.io/gorm"
)

type Organization struct {
	db *gorm.DB
}

func NewOrganization(db *gorm.DB) query.Organization {
	return Organization{db: db}
}

func (o Organization) List(ctx context.Context, ids []model.ID) ([]model.Organization, error) {
	orgs := []entity.Organization{}
	if err := o.preload(ctx).Find(&orgs, ids).Error; err != nil {
		return []model.Organization{}, apperror.NewGormFind(err, entity.TableNameOrganization)
	}

	return adapter.OrganizationListFromEntityList(orgs), nil
}

func (c Organization) preload(ctx context.Context) *gorm.DB {
	return c.db.
		Session(&gorm.Session{Logger: logger.NewZerologToGormLogger(ctx)}).
		Preload("OrganizationDetail")
}
