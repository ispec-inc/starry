package adapter

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/infra/entity"
)

func OrganizationToEntity(m organization.Organization) entity.Organization {
	return entity.Organization{
		ID: string(m.ID),
		OrganizationDetail: entity.OrganizationDetail{
			Name:    string(m.Name),
			Contact: string(m.Contact),
		},
	}
}

func OrganizationListFromEntityList(ens []entity.Organization) []organization.Organization {
	ms := make([]organization.Organization, len(ens))
	for i := range ens {
		ms[i] = organization.Organization{
			ID:      domain.ID(ens[i].ID),
			Name:    organization.OrganizationName(ens[i].OrganizationDetail.Name),
			Contact: organization.PhoneNumber(ens[i].OrganizationDetail.Contact),
		}
	}

	return ms
}
