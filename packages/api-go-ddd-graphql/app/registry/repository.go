package registry

import (
	"gorm.io/gorm"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/query"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/reader"
)

// Repository リポジトリのレジストリ
type Repository struct {
	db *gorm.DB
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

// NewOrganizationQuery query.Organizationの実装を返す
func (r Repository) NewOrganizationQuery() query.Organization {
	return reader.NewOrganization(r.db)
}
