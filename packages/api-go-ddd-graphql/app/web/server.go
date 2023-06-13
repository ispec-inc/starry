package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/config"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/controller"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/registry"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/app/schema"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/middleware"
	"github.com/ispec-inc/starry/api-go-ddd-graphql/pkg/presenter"
)

const PORT = 9000

func newServer() (*http.Server, error) {
	s, err := schema.String()
	if err != nil {
		return nil, err
	}

	rgst, err := registry.New()
	if err != nil {
		return nil, err
	}

	schema := graphql.MustParseSchema(s, controller.New(rgst))

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
		presenter.Text(w, s)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Response(w, map[string]string{"messsage": "ok"})
	})

	port := fmt.Sprintf(":%d", PORT)

	srv := &http.Server{Addr: port, Handler: r}
	return srv, nil

}
