package main

import (
	"context"
	"log"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/twiglab/crm/authz/verify"
	"github.com/twiglab/crm/authz/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}

	addr := os.Getenv("SERVER_ADDR")
	secret := os.Getenv("JWT_SECRET")

	conf := verify.AuthJWTConfig{Secret: secret}

	ctx := context.Background()

	r := chi.NewMux()
	r.Use(middleware.Logger, middleware.Recoverer)

	r.Mount("/jwt", verify.JWTRoute(conf))

	srv := web.NewHttpServer(ctx, addr, r)

	if err := web.RunServer(ctx, srv); err != nil {
		log.Fatal(err)
	}
}
