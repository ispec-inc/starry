package gqlerror

import (
	"context"

	"golang.org/x/text/language"
)

var (
	tagKey  = struct{}{}
	matcher = language.NewMatcher([]language.Tag{
		language.Japanese,
		language.English,
	})
)

// ContextWithTag 言語タグをコンテキストに埋め込む関数
func ContextWithTag(ctx context.Context, acceptLanguage string) context.Context {
	tags, _, err := language.ParseAcceptLanguage(acceptLanguage)
	if err != nil {
		return ctx
	}
	tag, _, _ := matcher.Match(tags...)
	return context.WithValue(ctx, tagKey, tag)
}

// TagFromContext コンテキストから言語タグを取得する関数
func TagFromContext(ctx context.Context) language.Tag {
	if tag, ok := ctx.Value(tagKey).(language.Tag); ok {
		return tag
	}
	return language.Und
}
