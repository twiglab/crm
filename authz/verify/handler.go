package  verify

import (
	"github.com/go-chi/chi/v5"
)

func JWTRoute(conf AuthJWTConfig) chi.Router {
	r := chi.NewRouter()
	r.HandleFunc("/verify", Verify(conf))
	return r
}
