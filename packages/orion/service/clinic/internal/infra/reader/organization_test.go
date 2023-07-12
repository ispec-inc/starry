package reader_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ispec-inc/starry/orion/app/test"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/entity"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/infra/reader"
	"github.com/rs/zerolog/log"
)

func TestOrganization_Get(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			ids []domain.ID
		}
	)

	db, cleanup := test.PrepareDB(t, "Organization_List", []interface{}{
		&entity.Organization{
			ID: "organization_list_uuid_1",
			OrganizationDetail: entity.OrganizationDetail{
				Name: "医療法人ispec",
			},
		},
		&entity.Organization{
			ID: "organization_list_uuid_2",
			OrganizationDetail: entity.OrganizationDetail{
				Name: "株式会社ispec",
			},
		},
	})
	defer cleanup()

	tests := []struct {
		name string
		give give
		want []organization.Organization
	}{
		{
			name: "[OK] idが存在する場合、Organizationを取得できる",
			give: give{
				ids: []domain.ID{"organization_list_uuid_1"},
			},
			want: []organization.Organization{
				{
					ID:   "organization_list_uuid_1",
					Name: "医療法人ispec",
				},
			},
		},
		{
			name: "[OK] idが空の場合、Organizationを全件取得できる",
			give: give{
				ids: []domain.ID{},
			},
			want: []organization.Organization{
				{
					ID:   "organization_list_uuid_1",
					Name: "医療法人ispec",
				},
				{
					ID:   "organization_list_uuid_2",
					Name: "株式会社ispec",
				},
			},
		},
		{
			name: "[OK] idが存在しない場合、空のOrganizationを取得できる",
			give: give{
				ids: []domain.ID{"not_organization_list_uuid"},
			},
			want: []organization.Organization{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()
			logger := log.With().Caller().Str(tt.name, t.Name()).Logger()
			ctx = logger.WithContext(ctx)

			o := reader.Organization{}
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
