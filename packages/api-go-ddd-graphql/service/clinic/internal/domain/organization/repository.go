package organization

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
)

// Query Organizationのクエリ
type Query interface {
	// Get Organizationを取得する
	Get(ctx context.Context, tx *app.DB, ids []domain.ID) ([]Organization, error)
}

// Command Organizationのコマンド
type Command interface {
	// Create Organizationを作成する
	Create(ctx context.Context, tx *app.DB, org Organization) error
}
