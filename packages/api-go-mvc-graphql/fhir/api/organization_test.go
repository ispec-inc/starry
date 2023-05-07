package api_test

import (
	"context"
	"testing"

	"github.com/graph-gophers/graphql-go"
	gqlclient "github.com/hasura/go-graphql-client"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/model"
	"github.com/stretchr/testify/assert"
)

func TestController_GetOrganization(t *testing.T) {
	type result struct {
		ID          gqlclient.ID
		Name        string
		Alias       string
		Contact     string
		Type        string
		Description string
	}

	type (
		query struct {
			Organization result `graphql:"organization(input: {id: $organizationId})"`
		}
		give struct {
			query query
			vars  map[string]interface{}
		}
		want struct {
			result result
		}
	)

	client, clnup, err := newClient(t, "Controller_Organization", []interface{}{
		&model.Organization{
			ID: "f170c10c-6896-46fc-b9a4-69e2b9a15154",
			OrganizationDetail: model.OrganizationDetail{
				Name:             "名前",
				Alias:            "Alias",
				OrganizationType: 1,
				Contact:          "08011112222",
				Description:      "説明",
			},
			Active: &model.ActiveOrganization{},
		},
	})

	defer clnup()
	assert.NoError(t, err)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "organization",
			give: give{
				vars: map[string]interface{}{
					"organizationId": graphql.ID("f170c10c-6896-46fc-b9a4-69e2b9a15154"),
				},
			},
			want: want{
				result{
					ID:          "f170c10c-6896-46fc-b9a4-69e2b9a15154",
					Name:        "名前",
					Alias:       "Alias",
					Contact:     "08011112222",
					Type:        "PROV",
					Description: "説明",
				},
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			q := tt.give.query
			err := client.Query(context.Background(), &q, tt.give.vars)
			assert.NoError(t, err)
			assert.Equal(t, tt.want.result, q.Organization)
		})
	}
}
