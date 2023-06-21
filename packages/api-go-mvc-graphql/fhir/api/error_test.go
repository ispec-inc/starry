package api_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/api"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

func TestError_Error(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			err error
		}
		want struct {
			err string
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "fhir_invalid",
			give: give{
				err: fhir.WithErrorCode(errors.New("invalid error"), fhir.ErrorCodeInvalid),
			},
			want: want{
				err: "invalid error",
			},
		},
		{
			name: "error_unknown",
			give: give{
				err: errors.New("unknown error"),
			},
			want: want{
				err: "unknown error",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := api.NewError(context.Background(), tt.give.err).Error()
			assert.Equal(t, tt.want.err, got)
		})
	}
}

func TestError_Extensions(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			tag language.Tag
			err error
		}
		want struct {
			extensions map[string]interface{}
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "ja_fhir_invalid",
			give: give{
				tag: language.Japanese,
				err: fhir.WithErrorCode(errors.New("invalid error"), fhir.ErrorCodeInvalid),
			},
			want: want{
				extensions: map[string]interface{}{
					"code":    "INVALID",
					"message": "リクエストが無効です。",
				},
			},
		},
		{
			name: "en_fhir_invalid",
			give: give{
				tag: language.English,
				err: fhir.WithErrorCode(errors.New("invalid error"), fhir.ErrorCodeInvalid),
			},
			want: want{
				extensions: map[string]interface{}{
					"code":    "INVALID",
					"message": "Invalid request",
				},
			},
		},
		{
			name: "error_unknown",
			give: give{
				err: errors.New("unknown error"),
			},
			want: want{
				extensions: nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.WithValue(context.Background(), api.CtxLanguageKey, tt.give.tag)
			got := api.NewError(ctx, tt.give.err).Extensions()
			assert.Equal(t, tt.want.extensions, got)
		})
	}
}
