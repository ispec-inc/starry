package registry

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/infra/reader"
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

func (r Repository) DB() *app.DB {
	return r.db
}

// NewOrganizationQuery query.Organizationの実装を返す
func (r Repository) NewOrganizationQuery() organization.Query {
	return reader.NewOrganization()
}
