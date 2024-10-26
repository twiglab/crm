package main

import (
	"fmt"
	"log"
	"wxlogin/model"
)

func main() {
	x := &model.AuthInfo{JsCode: "a", WxOpenID: "b", Code: "c", PhoneNo: "d"}
	bs, err := x.MarshalMsg(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bs)
}
