package wxlogin

import (
	"context"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/json"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/google/uuid"

	"github.com/twiglab/crm/wxlogin/wx"
)

type AuthJWTConfig struct {
	Alg    string `json:"alg"`
	Secret string `json:"secret" yaml:"secret" mapstructure:"secret"`
	Expire int    `json:"expire" yaml:"expire" mapstructure:"expire"`
}

type X struct {
	JwtStr string
	Code   string
	Claims *jwt.Claims
}

type AuthClient struct {
	MemberCli *wx.MemberCli
	WxCli     *wx.WxCli

	Secret []byte
}

// func (x *XClient) Login(ctx context.Context, wxOpenID string) error {
// 	x.MemberCli.QueryWxMember(ctx, wxOpenID)
// 	return nil
// }

// Login 用户登录，jscode换取用户jwt
//
//	@param ctx
//	@param jsCode
//	@return error
func (x *AuthClient) Login(ctx context.Context, jsCode string) (string, error) {
	codes, err := x.WxCli.AuthUser(ctx, jsCode)
	if err != nil {
		return "", err
	}

	member, err := x.MemberCli.LoginOrCr(ctx, codes.OpenID)
	if err != nil {
		return "", err
	}

	jwtStr, err := signed(NewClaims(member.Code), []byte("secret"))
	if err != nil {
		return "", err
	}

	return jwtStr, err
}

func (x *AuthClient) Login2(ctx context.Context, jsCode string) (*X, error) {
	codes, err := x.WxCli.AuthUser(ctx, jsCode)
	if err != nil {
		return nil, err
	}

	member, err := x.MemberCli.LoginOrCr(ctx, codes.OpenID)
	if err != nil {
		return nil, err
	}

	claims := NewClaims(member.Code)

	jwtStr, err := signed(claims, x.Secret)
	if err != nil {
		return nil, err
	}

	return &X{JwtStr: jwtStr, Code: member.Code, Claims: claims}, nil
}

type NonceSource struct {
}

func (NonceSource) Nonce() (string, error) {
	return uuid.NewString(), nil
}

func NewClaims(code string) *jwt.Claims {
	return &jwt.Claims{
		Issuer:    "https://xx/authz",
		ID:        uuid.NewString(),
		Audience:  []string{"app"},
		Subject:   code,
		Expiry:    jwt.NewNumericDate(time.Now().Add(128 * 24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
}

var nonce NonceSource = NonceSource{}

func signed(claims *jwt.Claims, key any) (string, error) {
	bs, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	singner, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.HS512,
		Key:       key,
	}, &jose.SignerOptions{NonceSource: nonce})

	if err != nil {
		return "", err
	}

	ss, err := singner.Sign(bs)
	if err != nil {
		return "", err
	}

	return ss.CompactSerialize()
}

/*
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
*/
