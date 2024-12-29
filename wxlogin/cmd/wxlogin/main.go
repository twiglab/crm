package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Khan/genqlient/graphql"

	"github.com/twiglab/crm/psdk/conf"
	"github.com/twiglab/crm/psdk/webx"

	"github.com/twiglab/crm/wxlogin"
	"github.com/twiglab/crm/wxlogin/gql"
	"github.com/twiglab/crm/wxlogin/wx"

	"github.com/twiglab/crm/wxlogin/cmd/wxlogin/config"
)

func main() {

	cfgCtx := conf.WithContext(context.Background())
	_ = cfgCtx.ReadInConfig()

	appConf := config.App{}
	_ = cfgCtx.Unmarshal(&appConf)

	secret := []byte(appConf.Jwt.Secret)

	wxCli := graphql.NewClient(appConf.Wechat.RpcUrl, http.DefaultClient)
	mbCli := graphql.NewClient(appConf.Member.RpcUrl, http.DefaultClient)

	mb := &wx.MemberCli{Client: mbCli}
	wx := &wx.WxCli{Client: wxCli}

	cli := &wxlogin.AuthClient{MemberCli: mb, WxCli: wx, Secret: secret}

	root := chi.NewMux()

	root.Use(middleware.Logger, middleware.Recoverer)

	root.Mount("/gql", gql.GQLRouter(cli))
	root.Mount("/jwt", wxlogin.JWTVerify(secret))

	webx.SetBaseContext(cfgCtx)
	if err := webx.ListenAndServe(appConf.Web.Addr, root); err != nil {
		log.Fatal(err)
	}

}
