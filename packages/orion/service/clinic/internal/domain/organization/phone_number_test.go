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
			phoneNumber organization.PhoneNumber
			err         error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "[OK] valid mobile phone number",
			give: give{
				phoneNumber: "09011112222",
			},
			want: want{
				phoneNumber: organization.PhoneNumber("09011112222"),
				err:         nil,
			},
		},
		{
			name: "[OK] valid telephone number",
			give: give{
				phoneNumber: "0451112222",
			},
			want: want{
				phoneNumber: organization.PhoneNumber("0451112222"),
				err:         nil,
			},
		},
		{
			name: "[NG] invalid string",
			give: give{
				phoneNumber: "住所",
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

			got, err := organization.NewPhoneNumber(tt.give.phoneNumber)

			if !errors.Is(err, tt.want.err) {
				t.Fatalf("expected %v to be %v", err, tt.want.err)
			}

			if cmp.Diff(got, tt.want.phoneNumber) != "" {
				t.Fatalf("expected %v to be %v", got.String(), tt.give.phoneNumber)
			}
		})
	}

}
