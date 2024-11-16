package wxlogin

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"
	"wxlogin/web"
	"wxlogin/wx"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/json"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/google/uuid"
)

type XClient struct {
	*wx.MemberCli
	*wx.WxCli
}

func (x *XClient) Login(ctx context.Context, wxOpenID string) error {
	x.MemberCli.QueryWxMember(ctx, wxOpenID)
	return nil
}

type AuthJWTConfig struct {
	Alg    string `json:"alg"`
	Secret string `json:"secret" yaml:"secret" mapstructure:"secret"`
	Expire int    `json:"expire" yaml:"expire" mapstructure:"expire"`
}

type NonceSource struct {
}

func (NonceSource) Nonce() (string, error) {
	return uuid.NewString(), nil
}

func NewClaims(code string) *jwt.Claims {
	return &jwt.Claims{
		Issuer:    "https://yabu.net.cn/authz",
		ID:        uuid.NewString(),
		Audience:  []string{"yabu-om", "yabu-flow"},
		Subject:   code,
		Expiry:    jwt.NewNumericDate(time.Now().Add(128 * 24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
}

func Auth(config AuthJWTConfig) http.HandlerFunc {

	secret := []byte(config.Secret)

	return func(w http.ResponseWriter, r *http.Request) {
		var (
			code string
		)

		if code = r.PostFormValue("code"); code == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		token := NewClaims(u.Code)
		tokenString, err := signed(token, secret)
		if err != nil {
			http.Error(w, "Not found user", http.StatusUnauthorized)
			return
		}

		_ = web.JsonTo(http.StatusOK,
			map[string]any{
				"code": "OK",
				"data": map[string]any{"token": tokenString, "name": u.Name, "code": u.Code}},
			w,
		)
	}
}

func signed(claims *jwt.Claims, key any) (string, error) {
	bs, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	singner, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.HS512,
		Key:       key,
	}, &jose.SignerOptions{NonceSource: NonceSource{}})

	if err != nil {
		return "", err
	}

	ss, err := singner.Sign(bs)
	if err != nil {
		return "", err
	}

	return ss.CompactSerialize()
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
