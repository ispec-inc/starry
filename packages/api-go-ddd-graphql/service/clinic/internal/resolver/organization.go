package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
)

type Organization struct {
	model organization.Organization
}

func NewOrganization(
	m organization.Organization,
) Organization {
	return Organization{model: m}
}

func (c Organization) ID() graphql.ID {
	return graphql.ID(c.model.ID)
}

func (c Organization) Name() string {
	return string(c.model.Name)
}
