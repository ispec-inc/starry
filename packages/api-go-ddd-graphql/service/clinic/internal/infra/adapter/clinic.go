package adapter

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/infra/entity"
)

// OrganizationToEntity OrganizationをEntityに変換する
func OrganizationToEntity(m organization.Organization) entity.Organization {
	return entity.Organization{
		ID: string(m.ID),
		OrganizationDetail: entity.OrganizationDetail{
			Name:    string(m.Name),
			Contact: string(m.Contact),
		},
	}
}

// OrganizationListFromEntityList EntityのリストをOrganizationのリストに変換する
func OrganizationListFromEntityList(ens []entity.Organization) []organization.Organization {
	ms := make([]organization.Organization, len(ens))
	for i := range ens {
		ms[i] = organization.Organization{
			ID:      organization.ID(ens[i].ID),
			Name:    organization.Name(ens[i].OrganizationDetail.Name),
			Contact: organization.PhoneNumber(ens[i].OrganizationDetail.Contact),
		}
	}

	return ms
}
