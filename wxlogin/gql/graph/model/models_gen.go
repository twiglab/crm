// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type JsCodeReq struct {
	JsCode string `json:"jsCode"`
}

type LoginUserResp struct {
	Code   string `json:"code"`
	Jwt    string `json:"jwt"`
	Level  int32  `json:"level"`
	Status int32  `json:"status"`
}

type Query struct {
}
