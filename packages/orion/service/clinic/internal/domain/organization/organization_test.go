package organization_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
)

func TestNew(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			name             organization.Name
			organizationType organization.Type
			phoneNumber      organization.PhoneNumber
		}
		want struct {
			org organization.Organization
			err error
		}
	)

	name, err := organization.NewName("鈴木歯科医院", "suzuki")
	if err != nil {
		t.Fatal(err)
	}

	typ, err := organization.NewType(1)
	if err != nil {
		t.Fatal(err)
	}

	phoneNumber, err := organization.NewPhoneNumber("09012345678")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "[OK] return organization",
			give: give{
				name:             name,
				organizationType: typ,
				phoneNumber:      phoneNumber,
			},
			want: want{
				org: organization.Organization{
					Name:        name,
					Type:        typ,
					PhoneNumber: phoneNumber,
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			org, err := organization.New(tt.give.name, tt.give.organizationType, tt.give.phoneNumber)
			if !errors.Is(err, tt.want.err) {
				t.Errorf("error: %v, want: %v", err, tt.want.err)
			}

			if diff := cmp.Diff(tt.want.org, org, cmpopts.IgnoreFields(organization.Organization{}, "ID")); diff != "" {
				t.Errorf("diff: %v", diff)
			}
		})
	}
}
