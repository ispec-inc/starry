package gqlerror

import (
	"golang.org/x/text/language"
)

type Presenter struct {
	Code     string
	Lang2Msg map[language.Tag]string
}
