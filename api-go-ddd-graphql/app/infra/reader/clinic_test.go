package reader_test

import (
	"context"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/domain/model"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/entity"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/infra/reader"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/rdb"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/testool"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestClinic_Get(t *testing.T) {
	type (
		give struct {
			id model.ID
		}
		want struct {
			name model.ClinicName
		}
	)

	db, cleanup := testool.PrepareDB(t, rdb.DBMSMySQL, "Clinic_Get", []interface{}{
		&entity.Clinic{ID: "uuid"},
		&entity.ClinicDetail{ID: "uuid", ClinicID: "uuid", Name: "name"},
	})
	defer cleanup()

	test := struct {
		name string
		give give
		want want
	}{
		give: give{
			id: "uuid",
		},
		want: want{
			name: "name",
		},
	}

	t.Run(test.name, func(t *testing.T) {
		ctx := context.Background()
		logger := log.With().Caller().Str("test case", t.Name()).Logger()
		ctx = logger.WithContext(ctx)

		c := reader.NewClinic(db)
		clinic, err := c.Get(ctx, test.give.id)
		assert.NoError(t, err)
		assert.Equal(t, test.want.name, clinic.Name)
	})
}
