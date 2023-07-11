package gqlerror

import (
	"golang.org/x/text/language"
)

// Presenter エラーを表示するための情報を埋め込む構造体
// 以下のようなメッセージが生成される
//
//	{
//	    "code": "UNAUTHORIZED", <- Presenter.Code
//	    "message": "リクエストが無効です。" <- Presenter.Lang2Msgを言語毎にローカライズ
//	}
type Presenter struct {
	Code     string
	Lang2Msg map[language.Tag]string
}
