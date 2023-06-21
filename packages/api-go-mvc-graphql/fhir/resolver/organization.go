package resolver

import (
	"context"
	"encoding/json"

	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/enum"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/model"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/query"
	"github.com/redis/go-redis/v9"
)

type Organization struct {
	Registry *fhir.Registry
	Model    *model.Organization
}

func (c Organization) ID() graphql.ID {
	return graphql.ID(c.Model.ID)
}

func (c Organization) Name() string {
	return string(c.Model.OrganizationDetail.Name)
}

func (c Organization) Alias() string {
	return string(c.Model.OrganizationDetail.Alias)
}

func (c Organization) Description() string {
	return string(c.Model.OrganizationDetail.Description)
}

func (c Organization) PartOf(ctx context.Context) (*Organization, error) {
	if c.Model.PartOf == nil {
		return nil, nil
	}

	var organization *model.Organization
	if val, err := c.Registry.Cache.Get(ctx, string(c.Model.PartOf.PartOfOrganizationID)).Result(); err == nil {
		if err := json.Unmarshal([]byte(val), &organization); err != nil {
			return nil, err
		}
	} else if err == redis.Nil {
		organization, err = c.organizationDefaultQuery(ctx).First()
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(organization)
		if err != nil {
			return nil, err
		}

		if err := c.Registry.Cache.Set(ctx, string(c.Model.PartOf.PartOfOrganizationID), data, 0).Err(); err != nil {
			return nil, err
		}
	}

	return &Organization{
		Registry: c.Registry,
		Model:    organization,
	}, nil
}

func (c Organization) Type() string {
	return enum.OrganizationTypeString(c.Model.OrganizationDetail.OrganizationType)
}

func (c Organization) Contact() string {
	return string(c.Model.OrganizationDetail.Contact)
}

func (c Organization) organizationDefaultQuery(ctx context.Context) query.IOrganizationDo {
	return c.Registry.Q.Organization.WithContext(ctx).Preload(
		c.Registry.Q.Organization.OrganizationDetail,
		c.Registry.Q.Organization.Active,
		c.Registry.Q.Organization.Cancel,
	)
}
