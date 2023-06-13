package reader_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/reader"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/test"
	"github.com/rs/zerolog/log"
)

func TestOrganization_List(t *testing.T) {
	type (
		give struct {
			ids []model.ID
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
		want []model.Organization
	}{
		{
			name: "[OK] idが存在する場合、Organizationを取得できる",
			give: give{
				ids: []model.ID{"organization_list_uuid_1"},
			},
			want: []model.Organization{
				{
					ID:   "organization_list_uuid_1",
					Name: "医療法人ispec",
				},
			},
		},
		{
			name: "[OK] idが空の場合、Organizationを全件取得できる",
			give: give{
				ids: []model.ID{},
			},
			want: []model.Organization{
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
				ids: []model.ID{"not_organization_list_uuid"},
			},
			want: []model.Organization{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			logger := log.With().Caller().Str(tt.name, t.Name()).Logger()
			ctx = logger.WithContext(ctx)

			c := reader.NewOrganization(db)
			orgs, err := c.List(ctx, tt.give.ids)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, orgs); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
