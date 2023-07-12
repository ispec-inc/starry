package organization_test

import (
	"errors"
	"testing"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
)

func TestPhoneNumber_Validate(t *testing.T) {
	t.Parallel()

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
				err: organization.ErrPhoneNumberInvalidFormat,
			},
		},
		{
			name: "invalid_string",
			give: give{
				phone: "012223",
			},
			want: want{
				err: organization.ErrPhoneNumberInvalidFormat,
			},
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := tt.give.phone.Validate()
			if !errors.Is(err, tt.want.err) {
				t.Fatalf("expected %v to be %v", err, tt.want.err)
			}

		})
	}

}
