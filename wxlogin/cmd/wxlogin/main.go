package main

import (
	"context"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/Khan/genqlient/graphql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/twiglab/crm/wxlogin"
	"github.com/twiglab/crm/wxlogin/gql"
	"github.com/twiglab/crm/wxlogin/web"
	"github.com/twiglab/crm/wxlogin/wx"
)

const (
	SERVER_ADDR = "SERVER_ADDR"
	JWT_SECRET  = "JWT_SECRET"
	MEMBER_URL  = "MEMBER_URL"
	WECHAT_URL  = "WECHAT_URL"
)

func Getenv(s string) string {
	x := os.Getenv(s)
	log.Printf("env %s = %s\n", s, x)
	return x
}

func main() {

	addr := Getenv(SERVER_ADDR)
	jwt := Getenv(JWT_SECRET)
	mbURL := Getenv(MEMBER_URL)
	wxURL := Getenv(WECHAT_URL)

	secret := []byte(jwt)

	wxCli := graphql.NewClient(wxURL, http.DefaultClient)
	mbCli := graphql.NewClient(mbURL, http.DefaultClient)

	mb := &wx.MemberCli{Client: mbCli}
	wx := &wx.WxCli{Client: wxCli}

	cli := &wxlogin.AuthClient{MemberCli: mb, WxCli: wx}

	root := chi.NewMux()

	root.Use(middleware.Logger, middleware.Recoverer)

	root.Mount("/gql", gql.GQLRouter(cli))
	root.Mount("/jwt", wxlogin.JWTVerify(secret))

	svr := web.NewHttpServer(context.Background(), addr, root)
	if err := web.RunServer(context.Background(), svr); err != nil {
		log.Fatal(err)
	}
}
