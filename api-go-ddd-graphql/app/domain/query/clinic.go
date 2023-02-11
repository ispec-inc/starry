package query

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
)

type Clinic interface {
	Get(ctx context.Context, id model.ID) (model.Clinic, error)
}
