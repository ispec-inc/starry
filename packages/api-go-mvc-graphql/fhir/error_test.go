package fhir_test

import (
	"testing"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestError_Unwrap(t *testing.T) {
	t.Parallel()
	type (
		give struct {
			code fhir.ErrorCode
			err  error
		}
		want struct {
			msg  string
			code fhir.ErrorCode
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "invalid",
			give: give{
				code: fhir.ErrorCodeInvalid,
				err:  errors.New("invalid"),
			},
			want: want{
				msg:  "invalid",
				code: fhir.ErrorCodeInvalid,
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := fhir.WithErrorCode(tt.give.err, tt.give.code)
			aerr := fhir.Unwrap(err)
			assert.Equal(t, aerr.Error(), tt.want.msg)
			assert.Equal(t, aerr.ErrorCode(), tt.want.code)
		})
	}
}

func TestError_Is(t *testing.T) {
	t.Parallel()
	type (
		give struct {
			err error
		}
		want struct {
			is bool
		}
	)

	err := errors.New("error")

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "true_new",
			give: give{
				err: fhir.New(err),
			},
			want: want{
				is: true,
			},
		},
		{
			name: "true_errorf",
			give: give{
				err: fhir.Wrap(fhir.ErrorCodeError, "wrap", err),
			},
			want: want{
				is: true,
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.True(t, errors.Is(tt.give.err, err))
		})
	}
}
