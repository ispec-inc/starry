package registry

import (
	"github.com/ispec-inc/starry/orion/app"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/repository"
)

// Repository リポジトリのレジストリ
type Repository struct {
	db *app.DB
}

// NewRepository リポジトリのレジストリを返す
// 内部でMySQLへのコネクションを確立する
func NewRepository(db *app.DB) Repository {
	return Repository{
		db: db,
	}
}

// DB MySQLへのコネクションを返す
func (r Repository) DB() *app.DB {
	return r.db
}

// NewOrganizationQuery query.Organizationの実装を返す
func (r Repository) NewOrganizationQuery() organization.Repository {
	return repository.Organization{}
}
