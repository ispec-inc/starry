package app_test

import (
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestError_Unwrap(t *testing.T) {
	t.Parallel()
	type (
		give struct {
			code app.ErrorCode
			err  error
		}
		want struct {
			msg  string
			code app.ErrorCode
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
				code: app.ErrorCodeInvalid,
				err:  errors.New("invalid"),
			},
			want: want{
				msg:  "invalid",
				code: app.ErrorCodeInvalid,
			},
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := app.WithErrorCode(tt.give.err, tt.give.code)
			aerr := app.Unwrap(err)
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
				err: app.NewError(err),
			},
			want: want{
				is: true,
			},
		},
		{
			name: "true_errorf",
			give: give{
				err: app.Wrap(app.ErrorCodeError, "wrap", err),
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
