package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type API struct {
	server *http.Server
}

func NewAPI() (API, func() error, error) {
	s, clnup, err := newServer()
	if err != nil {
		return API{}, nil, err
	}

	return API{
		server: s,
	}, clnup, nil
}

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
