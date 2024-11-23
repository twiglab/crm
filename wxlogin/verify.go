package wxlogin

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func Verify(secret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := ExtractToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.ParseSigned(t, allSignatureAlgorithms)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims := jwt.Claims{}
		if err := token.Claims(secret, &claims); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if err := claims.Validate(jwt.Expected{}); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}
}

var allSignatureAlgorithms = []jose.SignatureAlgorithm{
	jose.HS512,
	jose.HS256,
}

func ExtractToken(req *http.Request) (string, error) {
	tokenHeader := req.Header.Get("Authorization")
	if len(tokenHeader) < 7 || !strings.EqualFold(tokenHeader[:7], "bearer ") {
		return "", errors.New("no token")
	}
	return tokenHeader[7:], nil
}

func JWTVerify(secret []byte) chi.Router {
	r := chi.NewRouter()
	r.HandleFunc("/verify", Verify(secret))
	return r
}
