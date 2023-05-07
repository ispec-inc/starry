package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/config"
	"github.com/ispec-inc/starry/packages/api-go-mvc-graphql/fhir/schema"
	"golang.org/x/sync/errgroup"
)

const Port = 9000

// Controller はAPIのコントローラーです。
// GraphQLの操作がマウントされます。
type Controller struct {
	Registry *fhir.Registry
}

// API はAPIサーバーを起動するための構造体です。
type API struct {
	server *http.Server
}

// NewAPI はAPIサーバーを起動するための構造体を初期化します。
// ControllerにRegistryを注入し、エンドポイントとミドルウェアの設定を行います。
func New() (*API, error) {
	s, err := schema.String()
	if err != nil {
		return nil, err
	}

	registry, err := fhir.NewRegistry()
	if err != nil {
		return nil, err
	}

	cont := &Controller{
		Registry: registry,
	}
	schema := graphql.MustParseSchema(s, cont)

	h := &relay.Handler{Schema: schema}
	r := chi.NewRouter()
	r = NewMiddleware(r, MiddlewareConfig{
		Timeout:      config.Router.Timeout,
		AllowOrigins: config.Router.AllowOrigins,
	})

	r.Mount("/", h)
	r.Group(func(r chi.Router) {
		r.Mount("/graphql", h)
	})
	r.Get("/schema.graphql", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, s)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"messsage": "ok"})
	})

	port := fmt.Sprintf(":%d", Port)

	return &API{
		server: &http.Server{Addr: port, Handler: r},
	}, nil
}

// Run はAPIサーバーを起動します。
func (a API) Run(ctx context.Context) {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := a.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cs := make(chan os.Signal, 1)
	signal.Notify(cs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	select {
	case <-ctx.Done():
		break
	case <-cs:
		break
	}

	a.server.Shutdown(ctx)

	err := g.Wait()
	if err != nil {
		os.Exit(2)
	}
}
