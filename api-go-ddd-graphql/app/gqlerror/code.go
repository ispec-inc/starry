package gqlerror

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
)

// FIXME: クライアントが扱いやすい文字列に変更する。GraphQLのスキーマにTypeとして定義するのが理想。
var apperrorCodes = map[apperror.Code]string{
	apperror.CodeInvalid:      "INVALID",
	apperror.CodeUnauthorized: "UNAUTHORIZED",
	apperror.CodeNotFound:     "NOT_FOUND",
	apperror.CodeError:        "ERROR",
}

var domainErrorCodes = map[error]string{}
