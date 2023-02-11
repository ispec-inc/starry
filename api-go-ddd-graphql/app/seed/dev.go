package seed

import "github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"

func Dev() []interface{} {
	return []interface{}{
		&entity.Clinic{
			ID: "f170c10c-6896-46fc-b9a4-69e2b9a15154",
			ClinicDetail: entity.ClinicDetail{
				ID:       "f170c10c-6896-46fc-b9a4-69e2b9a15154",
				Name:     "名前",
				NameKana: "名前",
			},
		},
	}
}
