package gqlerror_test

// func TestError_Error(t *testing.T) {
// 	t.Parallel()
//
// 	type (
// 		give struct {
// 			err error
// 		}
// 		want struct {
// 			err string
// 		}
// 	)
//
// 	tests := []struct {
// 		name string
// 		give give
// 		want want
// 	}{
// 		{
// 			name: "apperror_invalid",
// 			give: give{
// 				err: app.WithErrorCode(errors.New("invalid error"), app.ErrorCodeInvalid),
// 			},
// 			want: want{
// 				err: "invalid error",
// 			},
// 		},
// 		{
// 			name: "error_unknown",
// 			give: give{
// 				err: errors.New("unknown error"),
// 			},
// 			want: want{
// 				err: "unknown error",
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			got := gqlerror.New(context.Background(), tt.give.err).Error()
// 			assert.Equal(t, tt.want.err, got)
// 		})
// 	}
// }
//
// func TestError_Extensions(t *testing.T) {
// 	t.Parallel()
//
// 	type (
// 		give struct {
// 			acceptLanguage string
// 			err            error
// 		}
// 		want struct {
// 			extensions map[string]interface{}
// 		}
// 	)
//
// 	tests := []struct {
// 		name string
// 		give give
// 		want want
// 	}{
// 		{
// 			name: "ja_apperror_invalid",
// 			give: give{
// 				acceptLanguage: "ja-JP",
// 				err:            app.WithErrorCode(errors.New("invalid error"), app.ErrorCodeInvalid),
// 			},
// 			want: want{
// 				extensions: map[string]interface{}{
// 					"code":    "INVALID",
// 					"message": "リクエストが無効です。",
// 				},
// 			},
// 		},
// 		{
// 			name: "en_apperror_invalid",
// 			give: give{
// 				acceptLanguage: "en-US",
// 				err:            app.WithErrorCode(errors.New("invalid error"), app.ErrorCodeInvalid),
// 			},
// 			want: want{
// 				extensions: map[string]interface{}{
// 					"code":    "INVALID",
// 					"message": "Invalid request",
// 				},
// 			},
// 		},
// 		{
// 			name: "error_unknown",
// 			give: give{
// 				err: errors.New("unknown error"),
// 			},
// 			want: want{
// 				extensions: nil,
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
//
// 			ctx := lang.ContextWithTag(context.Background(), tt.give.acceptLanguage)
// 			got := gqlerror.New(ctx, tt.give.err).Extensions()
// 			assert.Equal(t, tt.want.extensions, got)
// 		})
// 	}
// }

import (
	"context"
	"errors"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/gqlerror"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
)

var errTest = errors.New("error")

func TestError_Error(t *testing.T) {
	t.Parallel()

	type (
		give struct {
			lang       string
			presenters map[error]gqlerror.Presenter
			err        error
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
			name: "apperror_invalid",
			give: give{
				lang: "ja-JP",
				err:  errTest,
				presenters: map[error]gqlerror.Presenter{
					errTest: {
						Code: "INVALID",
						Lang2Msg: map[language.Tag]string{
							language.Japanese: "リクエストが無効です。",
							language.English:  "Invalid request",
						},
					},
				},
			},
			want: want{
				extensions: map[string]interface{}{
					"code":    "INVALID",
					"message": "リクエストが無効です。",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h, err := gqlerror.NewHandler(tt.give.presenters)
			assert.NoError(t, err)
			ctx := gqlerror.ContextWithTag(context.Background(), tt.give.lang)
			got := h.New(ctx, tt.give.err).Extensions()
			assert.Equal(t, tt.want.extensions, got)
		})
	}
}
