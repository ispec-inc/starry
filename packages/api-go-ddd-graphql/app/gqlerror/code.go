package gqlerror

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app"
)

// FIXME: クライアントが扱いやすい文字列に変更する。GraphQLのスキーマにTypeとして定義するのが理想。
var apperrorCodes = map[app.ErrorCode]string{
	app.ErrorCodeInvalid:      "INVALID",
	app.ErrorCodeUnauthorized: "UNAUTHORIZED",
	app.ErrorCodeNotFound:     "NOT_FOUND",
	app.ErrorCodeError:        "ERROR",
}

var domainErrorCodes = map[error]string{}
