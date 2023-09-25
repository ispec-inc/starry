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
func NewRepository() (Repository, error) {
	db, err := app.MySQL()
	if err != nil {
		return Repository{}, err
	}
	repo := Repository{
		db: db,
	}
	return repo, nil
}

// DB MySQLへのコネクションを返す
func (r Repository) DB() *app.DB {
	return r.db
}

// NewOrganizationRepository repository.Organizationの実装を返す
func (r Repository) NewOrganizationRepository() organization.Repository {
	return repository.Organization{}
}
