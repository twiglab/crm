// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ActiveCardReq struct {
	CodeBin string `json:"codeBin"`
}

type BindCardReq struct {
	Code       string `json:"code"`
	MemberCode string `json:"memberCode"`
}

type CardResp struct {
	Code          string  `json:"code"`
	CodeBin       string  `json:"codeBin"`
	Type          int     `json:"type"`
	Pic1          string  `json:"pic1"`
	Pic2          string  `json:"pic2"`
	Balance       int     `json:"balance"`
	Amount        int     `json:"amount"`
	MemberCode    *string `json:"memberCode,omitempty"`
	BindTime      *string `json:"bindTime,omitempty"`
	HitTime       int     `json:"hitTime"`
	LastCleanTime *string `json:"lastCleanTime,omitempty"`
	Status        int     `json:"status"`
}

type Mutation struct {
}

type PaginationReq struct {
	Cursor *string `json:"cursor,omitempty"`
	Limit  *int    `json:"limit,omitempty"`
	Order  *string `json:"order,omitempty"`
}

type Query struct {
}

type QueryCardByCode struct {
	Code string `json:"code"`
}

type QueryCardByMemberCode struct {
	MemberCode string `json:"memberCode"`
}
