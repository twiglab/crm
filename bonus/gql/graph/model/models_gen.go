// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type BonusDetailReq struct {
	MemberCode string `json:"memberCode"`
}

type BonusDetailResp struct {
	MemberCode string `json:"memberCode"`
	Balance    int    `json:"balance"`
}

type BonusListReq struct {
	MemberCode string `json:"memberCode"`
}

type BonusListResp struct {
	MemberCode string `json:"memberCode"`
}

type Query struct {
}
