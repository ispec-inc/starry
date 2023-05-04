package adapter

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"
)

func OrganizationToEntity(m model.Organization) entity.Organization {
	return entity.Organization{
		ID: string(m.ID),
		OrganizationDetail: entity.OrganizationDetail{
			Name:    string(m.Name),
			Contact: string(m.Contact),
		},
	}
}

func OrganizationListFromEntityList(ens []entity.Organization) []model.Organization {
	ms := make([]model.Organization, len(ens))
	for i := range ens {
		ms[i] = model.Organization{
			ID:      model.ID(ens[i].ID),
			Name:    model.OrganizationName(ens[i].OrganizationDetail.Name),
			Contact: model.PhoneNumber(ens[i].OrganizationDetail.Contact),
		}
	}

	return ms
}
