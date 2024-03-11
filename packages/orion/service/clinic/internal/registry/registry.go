package registry

import (
	"context"

	"github.com/ispec-inc/starry/orion/app"
)

// Registry インターフェースに対して実装を提供するレジストリ
type Registry struct {
	repo Repository
}

// New Registryのコンストラクタ
func New() (Registry, error) {
	db, err := app.MySQL()
	if err != nil {
		return Registry{}, err
	}
	repo := NewRepository(db)

	return Registry{
		repo: repo,
	}, nil
}

// NewTest テスト用Registryのコンストラクタ
func NewTest(ctx context.Context, db *app.DB) Registry {
	repo := NewRepository(db)

	return Registry{
		repo: repo,
	}
}

// Repository リポジトリのレジストリを返す
func (r Registry) Repository() Repository {
	return r.repo
}
