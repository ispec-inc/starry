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

type Clinic struct {
	db *gorm.DB
}

func NewClinic(db *gorm.DB) query.Clinic {
	return Clinic{db: db}
}

func (c Clinic) Get(ctx context.Context, id model.ID) (model.Clinic, error) {
	clinic := entity.Clinic{}
	if err := c.preload(ctx).First(&clinic, id).Error; err != nil {
		return model.Clinic{}, apperror.NewGormFind(err, clinic.TableName())
	}

	return adapter.ClinicFromEntity(clinic), nil
}

func (c Clinic) preload(ctx context.Context) *gorm.DB {
	return c.db.
		Session(&gorm.Session{Logger: logger.NewZerologToGormLogger(ctx)}).
		Preload("ClinicDetail")
}
