package seed

import "github.com/ispec-inc/starry/orion/service/clinic/internal/infra/entity"

// Dev 開発環境用のシードデータ
func Dev() []interface{} {
	return []interface{}{
		&entity.Organization{
			ID: "f170c10c-6896-46fc-b9a4-69e2b9a15154",
			OrganizationDetail: entity.OrganizationDetail{
				Name: "名前",
			},
		},
	}
}
