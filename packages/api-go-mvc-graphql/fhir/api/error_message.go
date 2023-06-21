package api

import (
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	ErrorCodeInvalid      = "INVALID"
	ErrorCodeUnauthorized = "UNAUTHORIZED"
	ErrorCodeNotFound     = "NOT_FOUND"
	ErrorCodeError        = "ERROR"
)

var messages = map[fhir.ErrorCode]map[language.Tag]string{
	fhir.ErrorCodeInvalid:      {language.Japanese: "リクエストが無効です。", language.English: "Invalid request"},
	fhir.ErrorCodeUnauthorized: {language.Japanese: "ログインしてください。", language.English: "Unauthorized"},
	fhir.ErrorCodeNotFound:     {language.Japanese: "データが見つかりません。", language.English: "Not found"},
	fhir.ErrorCodeError:        {language.Japanese: "エラーが発生しました。", language.English: "Error occurred"},
}

var codes = map[fhir.ErrorCode]string{
	fhir.ErrorCodeInvalid:      ErrorCodeInvalid,
	fhir.ErrorCodeUnauthorized: ErrorCodeUnauthorized,
	fhir.ErrorCodeNotFound:     ErrorCodeNotFound,
	fhir.ErrorCodeError:        ErrorCodeError,
}

func init() {
	for key, tagMsgs := range messages {
		for tag, msg := range tagMsgs {
			if err := message.SetString(tag, key.String(), msg); err != nil {
				panic(err)
			}
		}
	}
}
