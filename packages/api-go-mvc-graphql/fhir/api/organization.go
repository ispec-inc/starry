package api

import (
	"context"
	"crypto/rand"
	"encoding/json"

	"github.com/graph-gophers/graphql-go"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/enum"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/model"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/query"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/resolver"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/service"
	"github.com/oklog/ulid"
	"github.com/redis/go-redis/v9"
)

type OrganizationArgs struct {
	Input struct {
		ID graphql.ID
	}
}

func (c Controller) Organization(ctx context.Context, args OrganizationArgs) (resolver.Organization, error) {

	// NOTE これくらい肥大化したらサービスに切り出してもよいかも
	var organization *model.Organization
	val, err := c.Registry.Cache.Get(ctx, string(args.Input.ID)).Result()

	if err == nil {
		if err := json.Unmarshal([]byte(val), &organization); err != nil {
			return resolver.Organization{}, NewError(ctx, err)
		}
	} else if err == redis.Nil {
		// キャッシュヒットしなかった場合はDBから取得
		organization, err = c.organizationDefaultQuery(ctx).First()
		if err != nil {
			return resolver.Organization{}, NewError(ctx, err)
		}

		data, err := json.Marshal(organization)
		if err != nil {
			return resolver.Organization{}, NewError(ctx, err)
		}

		if err := c.Registry.Cache.Set(ctx, string(args.Input.ID), data, 0).Err(); err != nil {
			return resolver.Organization{}, NewError(ctx, err)
		}
	} else {
		return resolver.Organization{}, NewError(ctx, err)
	}

	return resolver.Organization{
		Model:    organization,
		Registry: c.Registry,
	}, nil
}

type RegisterOrganizationArgs struct {
	Input struct {
		Name                 string
		Alias                string
		Contact              string
		Description          string
		Type                 string
		PartOfOrganizationID *graphql.ID
	}
}

func (c Controller) RegisterOrganization(ctx context.Context, args RegisterOrganizationArgs) (resolver.Organization, error) {
	organization := &model.Organization{
		ID: ulid.MustNew(ulid.Now(), rand.Reader).String(),
		OrganizationDetail: model.OrganizationDetail{
			Name:             args.Input.Name,
			Alias:            args.Input.Alias,
			Contact:          args.Input.Contact,
			Description:      args.Input.Description,
			OrganizationType: enum.NewOrganizationType(args.Input.Type),
		},
		Active: &model.ActiveOrganization{},
	}

	if args.Input.PartOfOrganizationID != nil {
		organization.PartOf = &model.PartOfOrganization{
			PartOfOrganizationID: string(fhir.GraphQLIDValue(args.Input.PartOfOrganizationID)),
		}
	}

	if err := c.Registry.Validator.Struct(organization); err != nil {
		return resolver.Organization{}, NewError(ctx, err)
	}

	if err := c.organizationDefaultQuery(ctx).Create(organization); err != nil {
		return resolver.Organization{}, NewError(ctx, err)
	}

	return resolver.Organization{
		Model:    organization,
		Registry: c.Registry,
	}, nil
}

type CancelOrganizationArgs struct {
	Input struct {
		ID graphql.ID
	}
}

// CalcelOrganization はOrganizationを削除する
// Cancelは論理削除なので、実際には削除されない
func (c Controller) CancelOrganization(ctx context.Context, args CancelOrganizationArgs) (resolver.Organization, error) {
	cancel := service.CancelOrganization{
		Q: c.Registry.Q,
	}

	if err := cancel.Do(ctx, string(args.Input.ID)); err != nil {
		return resolver.Organization{}, NewError(ctx, err)
	}

	organization, err := c.organizationDefaultQuery(ctx).First()
	if err != nil {
		return resolver.Organization{}, NewError(ctx, err)
	}
	return resolver.Organization{
		Model:    organization,
		Registry: c.Registry,
	}, nil
}

func (c Controller) organizationDefaultQuery(ctx context.Context) query.IOrganizationDo {
	return c.Registry.Q.Organization.WithContext(ctx).Preload(
		c.Registry.Q.Organization.OrganizationDetail,
		c.Registry.Q.Organization.Active,
		c.Registry.Q.Organization.Cancel,
		c.Registry.Q.Organization.PartOf,
	)
}
