package domain

import (
	"github.com/oklog/ulid/v2"
)

// ID IDの値オブジェクト
type ID string

// NewID IDの値オブジェクトを生成するコンストラクタ
// ソート可能にするために、ULIDを採用している
func NewID() ID {
	id := ulid.Make()
	return ID(id.String())
}
