package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
)

type Organization struct {
	model model.Organization
}

func NewOrganization(
	m model.Organization,
) Organization {
	return Organization{model: m}
}

func (c Organization) ID() graphql.ID {
	return graphql.ID(c.model.ID)
}

func (c Organization) Name() string {
	return string(c.model.Name)
}
