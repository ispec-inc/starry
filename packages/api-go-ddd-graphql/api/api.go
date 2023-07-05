package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// API APIサーバーの構造体
type API struct {
	server *http.Server
}

// NewAPI APIサーバーのコンストラクタ
func NewAPI() (API, error) {
	s, err := newServer()
	if err != nil {
		return API{}, err
	}

	return API{
		server: s,
	}, nil
}

// Run APIサーバーを起動する関数
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

	if err := a.server.Shutdown(ctx); err != nil {
		log.Printf("error: %v\n", err)
	}

	if err := g.Wait(); err != nil {
		os.Exit(2)
	}
}
