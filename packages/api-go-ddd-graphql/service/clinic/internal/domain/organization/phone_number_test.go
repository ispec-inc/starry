package organization_test

import (
	"errors"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/stretchr/testify/assert"
)

func TestPhoneNumber_Validate(t *testing.T) {
	type (
		give struct {
			phone organization.PhoneNumber
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
			name: "valid_cell",
			give: give{
				phone: "09011112222",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "valid_tel",
			give: give{
				phone: "0451112222",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "invalid_string",
			give: give{
				phone: "住所",
			},
			want: want{
				err: errors.New("phone_number: invalid format"),
			},
		},
		{
			name: "invalid_string",
			give: give{
				phone: "012223",
			},
			want: want{
				err: errors.New("phone_number: invalid format"),
			},
		},
	}
	for i := range tests {
		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			err := test.give.phone.Validate()
			if test.want.err != nil {
				assert.EqualError(t, test.want.err, err.Error())
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
