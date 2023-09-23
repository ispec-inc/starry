package organization_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
)

func TestNewPhoneNumber(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			phoneNumber string
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
				phoneNumber: "09011112222",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "valid_tel",
			give: give{
				phoneNumber: "0451112222",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "invalid_string",
			give: give{
				phoneNumber: "住所",
			},
			want: want{
				err: organization.ErrPhoneNumberInvalidFormat,
			},
		},
		{
			name: "invalid_string",
			give: give{
				phoneNumber: "012223",
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

			p, err := organization.NewPhoneNumber(tt.give.phoneNumber)

			if !errors.Is(err, tt.want.err) {
				t.Fatalf("expected %v to be %v", err, tt.want.err)
			}

			if cmp.Diff(p.String(), tt.give.phoneNumber) != "" {
				t.Fatalf("expected %v to be %v", p.String(), tt.give.phoneNumber)
			}
		})
	}

}
