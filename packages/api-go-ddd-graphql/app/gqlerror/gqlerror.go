package gqlerror

import (
	"context"
	"errors"
)

// Error GraphQLのエラー
type Error struct {
	s          string
	extensions map[string]interface{}
}

// Handler GraphQLのエラーハンドラ
type Handler struct {
	Presenters map[error]Presenter
}

// New エラーを生成する
func (h Handler) New(ctx context.Context, err error) Error {
	for perr, pre := range h.Presenters {
		if !errors.Is(err, perr) {
			continue
		}

		v := Error{
			s: err.Error(),
			extensions: map[string]interface{}{
				"code": pre.Code,
			},
		}

		tag := TagFromContext(ctx)
		v.extensions["message"] = pre.Lang2Msg[tag]
		return v
	}

	return Error{s: err.Error()}
}

// Error エラーを文字列に変換する
func (e Error) Error() string {
	return e.s
}

// Extensions エラーの拡張情報を取得する
func (e Error) Extensions() map[string]interface{} {
	return e.extensions
}
