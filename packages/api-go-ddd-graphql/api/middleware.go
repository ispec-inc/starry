package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CommonConfig struct {
	Timeout      time.Duration
	AllowOrigins []string
}

func Common(r *chi.Mux, c CommonConfig) *chi.Mux {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(RequestLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(c.Timeout))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: c.AllowOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept", "Accept-Language", "Authorization",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link", "Authorization"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	return r
}

const logType = "request"

func RequestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		reqID := r.Header.Get("x-request-id")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		ctx := r.Context()
		logger := log.With().Caller().Str("request-id", reqID).Logger()
		ctx = logger.WithContext(ctx)

		r = r.WithContext(ctx)

		ww.Header().Set("request-id", reqID)

		t1 := time.Now()
		defer func() {
			logRequestInfo(ww, r, t1)
		}()

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}

func logRequestInfo(ww middleware.WrapResponseWriter, r *http.Request, timeFrom time.Time) {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	log.Ctx(r.Context()).Info().
		Str("type", logType).
		Str("scheme", scheme).
		Str("timestamp", time.Now().String()).
		Str("method", r.Method).
		Str("host", r.Host).
		Str("from", r.RemoteAddr).
		Str("status", fmt.Sprintf("%d", ww.Status())).
		Str("bytes", fmt.Sprintf("%d", ww.BytesWritten())).
		Str("elapsed", time.Since(timeFrom).String()).
		Send()
}
