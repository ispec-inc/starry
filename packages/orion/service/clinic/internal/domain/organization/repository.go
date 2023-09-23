package organization

import (
	"context"

	"github.com/ispec-inc/starry/orion/app"
)

// Repository Organizationのレポジトリ
type Repository interface {
	// Get Organizationを取得する
	Get(ctx context.Context, tx *app.DB, ids []ID) (List, error)
	// Create Organizationを作成する
	Create(ctx context.Context, tx *app.DB, org Organization) error
}
