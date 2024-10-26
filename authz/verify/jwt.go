package  verify

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

type Param struct {
	Code   string `json:"code,omitempty"`
	Passwd string `json:"passwd,omitempty"`
}

type AuthJWTConfig struct {
	Alg    string `json:"alg"`
	Secret string `json:"secret" yaml:"secret" mapstructure:"secret"`
	Expire int    `json:"expire" yaml:"expire" mapstructure:"expire"`
}

func Verify(config AuthJWTConfig) http.HandlerFunc {

	secret := []byte(config.Secret)

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
	jose.HS256,
}

func ExtractToken(req *http.Request) (string, error) {
	tokenHeader := req.Header.Get("Authorization")
	if len(tokenHeader) < 7 || !strings.EqualFold(tokenHeader[:7], "bearer ") {
		return "", errors.New("no token")
	}
	return tokenHeader[7:], nil
}
