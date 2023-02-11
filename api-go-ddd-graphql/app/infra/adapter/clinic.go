package adapter

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"
)

func ClinicToEntity(m model.Clinic) entity.Clinic {
	return entity.Clinic{
		ID: string(m.ID),
		ClinicDetail: entity.ClinicDetail{
			Name:     string(m.Name),
			NameKana: string(m.NameKana),
			Contact:  string(m.Contact),
		},
	}
}

func ClinicFromEntity(en entity.Clinic) model.Clinic {
	return model.Clinic{
		ID:       model.ID(en.ID),
		Name:     model.ClinicName(en.ClinicDetail.Name),
		NameKana: model.ClinicNameKana(en.ClinicDetail.NameKana),
		Contact:  model.PhoneNumber(en.ClinicDetail.Contact),
	}
}
