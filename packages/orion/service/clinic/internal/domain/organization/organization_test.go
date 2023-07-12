package organization_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain"
	"github.com/ispec-inc/starry/orion/service/clinic/internal/domain/organization"
)

func TestName_Validate(t *testing.T) {
	t.Parallel()

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
				name: "鈴木歯科医院",
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
				err: domain.ErrStringInvalidLength,
			},
		},
		{
			name: "name_is_just_50",
			give: give{
				name: strings.Repeat("あ", 50),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_over_50",
			give: give{
				name: strings.Repeat("あ", 51),
			},
			want: want{
				err: domain.ErrStringInvalidLength,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			name := organization.Name(tt.give.name)
			err := name.Validate()
			if !errors.Is(err, tt.want.err) {
				t.Fatalf("expected %v to be %v", err, tt.want.err)
			}
		})
	}
}
