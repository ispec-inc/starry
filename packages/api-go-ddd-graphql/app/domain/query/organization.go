package query

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
)

type Organization interface {
	Get(ctx context.Context, ids []model.ID) ([]model.Organization, error)
}
