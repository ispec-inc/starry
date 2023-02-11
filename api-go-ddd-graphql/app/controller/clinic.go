package controller

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/resolver"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/uc"
)

type GetClinicArgs struct {
	Input struct {
		ID graphql.ID
	}
}

func (c Controller) GetClinic(ctx context.Context, args GetClinicArgs) (resolver.Clinic, error) {

	ipt := uc.GetClinicInput{
		ID: model.ID(args.Input.ID),
	}
	get := uc.NewGetClinic(c.registry)

	opt, err := get.Do(ctx, ipt)
	if err != nil {
		return resolver.Clinic{}, err
	}

	r := resolver.NewClinic(opt.Clinic)
	return r, nil
}
