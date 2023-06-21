package fhir

import (
	"github.com/go-playground/validator/v10"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/query"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/validation"
	"github.com/redis/go-redis/v9"
)

// Registry はミドルウェアとのコネクションを保持する構造体
type Registry struct {
	Q         *query.Query        // Q はクエリを実行するための構造体
	Cache     *redis.Client       // Cache はRedisのクライアント
	Validator *validator.Validate // Validator はバリデーションを実行するための構造体
}

// New はミドルウェアとのコネクションを確立させてRegistryを返す
func NewRegistry() (*Registry, error) {
	db, err := MySQL()
	if err != nil {
		return nil, err
	}
	cache := Redis()

	val := validator.New()
	val.RegisterValidation("kana", validation.ValidateKana)

	return &Registry{
		Q:         query.Use(db),
		Cache:     cache,
		Validator: val,
	}, nil
}
