package gqlerror

import (
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/apperror"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	keyErrorInvalid                = "error_invalid"
	keyErrorUnauthorized           = "error_unauthorized"
	keyErrorNotFound               = "error_not_found"
	keyErrorGeneral                = "error_general"
	keyErrorTokenIsExpired         = "error_token_is_expired"
	keyErrorNoViewArticlePrivilege = "error_no_view_article_privilege"
)

// FIXME: 適切なメッセージに変更する。
var messages = map[string]map[language.Tag]string{
	keyErrorInvalid:        {language.Japanese: "リクエストが無効です。", language.English: "Invalid request"},
	keyErrorUnauthorized:   {language.Japanese: "ログインしてください。", language.English: "Unauthorized"},
	keyErrorNotFound:       {language.Japanese: "データが見つかりません。", language.English: "Not found"},
	keyErrorGeneral:        {language.Japanese: "エラーが発生しました。", language.English: "Error occurred"},
	keyErrorTokenIsExpired: {language.Japanese: "IDトークンの期限が切れています。", language.English: "ID Token is expired."},
}

var apperrorKeys = map[apperror.Code]string{
	apperror.CodeInvalid:      keyErrorInvalid,
	apperror.CodeUnauthorized: keyErrorUnauthorized,
	apperror.CodeNotFound:     keyErrorNotFound,
	apperror.CodeError:        keyErrorGeneral,
}

var domainErrorKeys = map[error]string{}

func init() {
	for key, tagMsgs := range messages {
		for tag, msg := range tagMsgs {
			if err := message.SetString(tag, key, msg); err != nil {
				panic(err)
			}
		}
	}
}
