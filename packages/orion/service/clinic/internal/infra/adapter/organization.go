package adapter

import (
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/entity"
)

// OrganizationToEntity OrganizationをEntityに変換する
func OrganizationToEntity(m organization.Organization) entity.Organization {
	return entity.Organization{
		ID: string(m.ID),
		OrganizationDetail: entity.OrganizationDetail{
			Name:             m.Name.Name.String(),
			Alias:            m.Name.Alias.String(),
			Contact:          m.PhoneNumber.String(),
			OrganizationType: int32(m.Type),
		},
	}
}

// OrganizationListToEntityList OrganizationListをEntityListに変換する
func OrganizationListToEntityList(ms organization.List) []entity.Organization {
	ens := make([]entity.Organization, len(ms))
	for i := range ms {
		ens[i] = OrganizationToEntity(ms[i])
	}
	return ens
}

// OrganizationListFromEntityList EntityのリストをOrganizationのリストに変換する
// もし値オブジェクトの初期化に失敗した場合はエラーを返す
func OrganizationListFromEntityList(ens []entity.Organization) (organization.List, error) {
	ms := make(organization.List, len(ens))
	for i := range ens {

		name, err := organization.NewName(ens[i].OrganizationDetail.Name, ens[i].OrganizationDetail.Alias)
		if err != nil {
			return organization.List{}, err
		}

		typ, err := organization.NewType(uint(ens[i].OrganizationDetail.OrganizationType))
		if err != nil {
			return organization.List{}, err
		}

		phone, err := organization.NewPhoneNumber(ens[i].OrganizationDetail.Contact)
		if err != nil {
			return organization.List{}, err
		}

		ms[i] = organization.Organization{
			ID:          organization.ID(ens[i].ID),
			Name:        name,
			Type:        typ,
			PhoneNumber: phone,
		}
	}

	return ms, nil
}
