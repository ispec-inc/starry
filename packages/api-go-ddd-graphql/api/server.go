package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/api/schema"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/middleware"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic"
)

type gqlController struct {
	clinic.Controller
}

const PORT = 9000

// newServer http.Serverを生成する。内部で各サービスのコントローラを呼び出し、schemaとの整合性チェックする
func newServer() (*http.Server, error) {

	s, err := schema.String()
	if err != nil {
		return nil, err
	}

	clinic, err := clinic.New()
	if err != nil {
		return nil, err
	}

	cont := &gqlController{*clinic}
	schema := graphql.MustParseSchema(s, cont)

	h := &relay.Handler{Schema: schema}
	r := chi.NewRouter()
	r = middleware.Common(r, middleware.CommonConfig{
		Timeout:      config.Router.Timeout,
		AllowOrigins: config.Router.AllowOrigins,
	})

	r.Mount("/", h)
	r.Group(func(r chi.Router) {
		r.Mount("/graphql", h)
	})

	r.Get("/schema.graphql", func(w http.ResponseWriter, r *http.Request) {
		Text(w, s)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		Response(w, map[string]string{"messsage": "ok"})
	})

	port := fmt.Sprintf(":%d", PORT)

	srv := &http.Server{Addr: port, Handler: r}
	return srv, nil

}
