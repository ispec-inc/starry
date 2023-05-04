package apperror_test

import (
	"errors"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
	"github.com/stretchr/testify/assert"
)

func TestError_Unwrap(t *testing.T) {
	type (
		give struct {
			code apperror.Code
			err  error
		}
		want struct {
			msg  string
			code apperror.Code
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
				code: apperror.CodeInvalid,
				err:  errors.New("invalid"),
			},
			want: want{
				msg:  "invalid",
				code: apperror.CodeInvalid,
			},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			err := apperror.WithCode(test.give.err, test.give.code)
			aerr := apperror.Unwrap(err)
			assert.Equal(t, aerr.Error(), test.want.msg)
			assert.Equal(t, aerr.Code(), test.want.code)
		})
	}
}

func TestError_Is(t *testing.T) {
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
				err: apperror.New(err),
			},
			want: want{
				is: true,
			},
		},
		{
			name: "true_errorf",
			give: give{
				err: apperror.Wrap(apperror.CodeError, "wrap", err),
			},
			want: want{
				is: true,
			},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			assert.True(t, errors.Is(test.give.err, err))
		})
	}
}
