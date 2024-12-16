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
	Jwt    string
	Code   string
	OpenID string
	Claims *jwt.Claims
}

type AuthClient struct {
	MemberCli *wx.MemberCli
	WxCli     *wx.WxCli

	Secret []byte

	// AuthJWTConfig AuthJWTConfig
}

func (x *AuthClient) Login3(ctx context.Context, jsCode string) (*X, error) {
	codes, err := x.WxCli.AuthUser(ctx, jsCode)
	if err != nil {
		return nil, err
	}

	Code := uuid.NewString()

	claims := NewClaims(Code)

	jwt, err := signed(claims, x.Secret)
	if err != nil {
		return nil, err
	}

	return &X{Jwt: jwt, Code: Code, OpenID: codes.OpenID, Claims: claims}, nil
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

	jwt, err := signed(claims, x.Secret)
	if err != nil {
		return nil, err
	}

	return &X{Jwt: jwt, Code: member.Code, OpenID: codes.OpenID, Claims: claims}, nil
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

type NonceSource struct {
}

func (NonceSource) Nonce() (string, error) {
	return uuid.NewString(), nil
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
