package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/wechat/pkg/data"
	"github.com/twiglab/crm/wechat/pkg/low"
)

func main() {
	client := graphql.NewClient("http://it9i.com:10009/rpc/gql", http.DefaultClient)

	resp, err := low.AuthUser(context.Background(), client, data.JsCodeReq{JsCode: "xx"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
