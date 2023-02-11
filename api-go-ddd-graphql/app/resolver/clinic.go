package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
)

type Clinic struct {
	model model.Clinic
}

func NewClinic(
	m model.Clinic,
) Clinic {
	return Clinic{model: m}
}

func (c Clinic) ID() graphql.ID {
	return graphql.ID(c.model.ID)
}

func (c Clinic) Name() string {
	return string(c.model.Name)
}

func (c Clinic) NameKana() string {
	return string(c.model.NameKana)
}
