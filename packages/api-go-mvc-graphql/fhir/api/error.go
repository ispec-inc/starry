package api

import (
	"context"

	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Error struct {
	s          string
	extensions map[string]interface{}
}

func NewError(ctx context.Context, err error) Error {
	aerr := fhir.Unwrap(err)
	if aerr == nil {
		return Error{
			s: err.Error(),
			extensions: map[string]interface{}{
				"code": ErrorCodeError,
			},
		}
	}
	code, ok := codes[aerr.ErrorCode()]
	if !ok {
		return Error{}
	}

	v := Error{
		s: err.Error(),
		extensions: map[string]interface{}{
			"code": code,
		},
	}

	var tag language.Tag
	if t, ok := ctx.Value(CtxLanguageKey).(language.Tag); ok {
		tag = t
	} else {
		tag = language.English
	}

	msg := message.NewPrinter(tag).Sprintf(aerr.ErrorCode())
	v.extensions["message"] = msg

	return v
}

func (e Error) Error() string {
	return e.s
}

func (e Error) Extensions() map[string]interface{} {
	return e.extensions
}
