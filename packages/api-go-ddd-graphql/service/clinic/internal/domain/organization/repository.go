package organization

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
)

type Query interface {
	Get(ctx context.Context, tx *app.DB, ids []domain.ID) ([]Organization, error)
}

type Command interface {
	Create(ctx context.Context, tx *app.DB, org Organization) error
}
