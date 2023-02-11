package uc

import (
	"context"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/query"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/registry"
)

type GetClinic struct {
	clinicQuery query.Clinic
}

type GetClinicInput struct {
	ID model.ID
}

type GetClinicOutput struct {
	Clinic model.Clinic
}

func NewGetClinic(r registry.Registry) GetClinic {
	return GetClinic{
		clinicQuery: r.Repository().NewClinicQuery(),
	}
}

func (g GetClinic) Do(ctx context.Context, ipt GetClinicInput) (GetClinicOutput, error) {
	c, err := g.clinicQuery.Get(ctx, ipt.ID)
	if err != nil {
		return GetClinicOutput{}, err
	}

	return GetClinicOutput{Clinic: c}, nil
}
