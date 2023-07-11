package gqlerror_test

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
