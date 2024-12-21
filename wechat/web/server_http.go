package web

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"time"
)

type ServerOptFunc func(*http.Server)

func ErrorLog(h slog.Handler, l slog.Level) ServerOptFunc {
	return func(s *http.Server) {
		s.ErrorLog = slog.NewLogLogger(h, l)
	}
}

func IdlerTimeout(d time.Duration) ServerOptFunc {
	return func(s *http.Server) {
		s.IdleTimeout = d
	}
}

func Handler(h http.Handler) ServerOptFunc {
	return func(s *http.Server) {
		s.Handler = h
	}
}

func Addr(addr string) ServerOptFunc {
	return func(s *http.Server) {
		s.Addr = addr
	}
}

func NewHttpServer(ctx context.Context, addr string, handler http.Handler, opts ...ServerOptFunc) *http.Server {
	s := &http.Server{
		Addr:        addr,
		Handler:     handler,
		IdleTimeout: 10 * time.Second,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	for _, f := range opts {
		f(s)
	}
	return s
}

/*
func NewWebServer(ctx context.Context, opts ...ServerOptFunc) *http.Server {
	s := &http.Server{
		IdleTimeout: 10 * time.Second,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	for _, f := range opts {
		f(s)
	}
	return s
}
*/
