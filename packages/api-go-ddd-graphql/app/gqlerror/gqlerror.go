package gqlerror

import (
	"context"
	"errors"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/api/lang"
)

type Error struct {
	s          string
	extensions map[string]interface{}
}

type Handler struct {
	Presenters map[error]Presenter
}

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

		tag := lang.TagFromContext(ctx)
		v.extensions["message"] = pre.Lang2Msg[tag]
		return v
	}

	return Error{s: err.Error()}
}

func (e Error) Error() string {
	return e.s
}

func (e Error) Extensions() map[string]interface{} {
	return e.extensions
}
