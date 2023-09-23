package repository_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ispec-inc/starry/orion/app/test"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/adapter"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/repository"
	"github.com/rs/zerolog/log"
)

func organizationFactory(count int) (organization.List, error) {
	orgs := make(organization.List, 0, count)
	for i := 0; i < count; i++ {
		name, err := organization.NewName("鈴木歯科医院", "suzuki")
		if err != nil {
			return organization.List{}, err
		}

		typ, err := organization.NewType(1)
		if err != nil {
			return organization.List{}, err
		}

		phone, err := organization.NewPhoneNumber("09012345678")
		if err != nil {
			return organization.List{}, err
		}

		org, err := organization.New(name, typ, phone)
		if err != nil {
			return organization.List{}, err
		}

		orgs = append(orgs, org)

	}

	return orgs, nil
}

func TestOrganization_Get(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			ids []organization.ID
		}
	)
	orgs, err := organizationFactory(3)
	if err != nil {
		t.Fatal(err)
	}

	ents := adapter.OrganizationListToEntityList(orgs)
	seeds := []interface{}{}
	for i := range ents {
		seeds = append(seeds, &ents[i])
	}

	db, cleanup := test.PrepareDB(t, "Organization_List", seeds)
	t.Cleanup(cleanup)

	tests := []struct {
		name string
		give give
		want organization.List
	}{
		{
			name: "[OK] idが存在する場合、Organizationを取得できる",
			give: give{
				ids: []organization.ID{orgs[0].ID},
			},
			want: organization.List{orgs[0]},
		},
		{
			name: "[OK] idを指定しない場合、Organizationを全件取得できる",
			give: give{
				ids: []organization.ID{},
			},
			want: orgs,
		},
		{
			name: "[OK] idが存在しない場合、空のOrganizationを取得できる",
			give: give{
				ids: []organization.ID{"not_organization_list_uuid"},
			},
			want: []organization.Organization{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			logger := log.With().Caller().Str(tt.name, t.Name()).Logger()
			ctx = logger.WithContext(ctx)

			o := repository.Organization{}
			orgs, err := o.Get(ctx, db, tt.give.ids)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tt.want, orgs); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}

func TestOrganization_Create(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			org organization.Organization
		}
	)

	db, cleanup := test.PrepareDB(t, "Organization_Create", []interface{}{})
	t.Cleanup(cleanup)

	org, err := organizationFactory(1)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		give give
		err  error
	}{
		{
			name: "[OK] idが存在する場合、Organizationを取得できる",
			give: give{
				org: org[0],
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// t.Parallel()

			ctx := context.Background()
			logger := log.With().Caller().Str(tt.name, t.Name()).Logger()
			ctx = logger.WithContext(ctx)

			o := repository.Organization{}
			err := o.Create(ctx, db, tt.give.org)
			if err != nil {
				t.Error(err)
			}

		})
	}
}
