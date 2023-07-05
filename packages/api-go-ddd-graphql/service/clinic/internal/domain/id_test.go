package domain_test

import (
	"errors"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_NewID(t *testing.T) {
	type (
		give struct {
			name string
		}
		want struct {
			err error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "no_error",
			give: give{
				name: "user",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_empty",
			give: give{
				name: "",
			},
			want: want{
				err: errors.New("id: name is empty"),
			},
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			id, err := domain.NewID()
			if test.want.err != nil {
				assert.EqualError(t, test.want.err, err.Error())
			}
			t.Log(id)
		})
	}
}
