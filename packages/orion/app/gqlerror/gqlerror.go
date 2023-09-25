package gqlerror

import (
	"context"
	"errors"

	"github.com/ispec-inc/starry/orion/app"
)

// Error GraphQLのエラー
type Error struct {
	s string
	// extensions GraphQLのエラーの拡張情報。最終的にユーザーに表示される
	extensions map[string]interface{}
}

func New(ctx context.Context, err error) Error {
	code := "error"

	if errors.Is(err, app.ErrUnauthorized) {
		return Error{
			s: err.Error(),
			extensions: map[string]interface{}{
				"code":    "unauthorized",
				"message": "認証が必要です",
			},
		}
	}

	v := Error{
		s: err.Error(),
		extensions: map[string]interface{}{
			"code":    code,
			"message": "エラーが発生しました",
		},
	}

	return v
}

func NewWithCode(ctx context.Context, err error, code string, message string) Error {
	v := Error{
		s: err.Error(),
		extensions: map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}

	return v
}

// Error エラーを文字列に変換する
func (e Error) Error() string {
	return e.s
}

// Extensions エラーの拡張情報を取得する
func (e Error) Extensions() map[string]interface{} {
	return e.extensions
}
