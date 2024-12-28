package webx

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

var gS = &http.Server{
	Addr:        ":10008",
	IdleTimeout: 10 * time.Second,
}

func Serve(l net.Listener, handler http.Handler) error {
	gS.Handler = handler
	return gS.Serve(l)
}

func ListenAndServe(addr string, handler http.Handler) error {
	gS.Handler = handler
	gS.Addr = addr
	return gS.ListenAndServe()
}

func SetBaseContext(ctx context.Context) {
	gS.BaseContext = func(l net.Listener) context.Context { return ctx }
}

func Shutdown(ctx context.Context, timeOut time.Duration) error {

	x, cancel := context.WithTimeout(ctx, timeOut) // 4秒后退出
	defer cancel()

	if err := gS.Shutdown(x); err != nil {
		log.Fatal("Server Shutdown:", err)
		return err
	}

	return nil
}
