// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package low

import (
	"context"

	"github.com/Khan/genqlient/graphql"
	"github.com/twiglab/crm/member/pkg/data"
)

// CreateWxMemberResponse is returned by CreateWxMember on success.
type CreateWxMemberResponse struct {
	CreateWxMember data.MemberResp `json:"CreateWxMember"`
}

// GetCreateWxMember returns CreateWxMemberResponse.CreateWxMember, and is useful for accessing the field via an interface.
func (v *CreateWxMemberResponse) GetCreateWxMember() data.MemberResp { return v.CreateWxMember }

// QueryWxMemberResponse is returned by QueryWxMember on success.
type QueryWxMemberResponse struct {
	QueryWxMember data.MemberResp `json:"QueryWxMember"`
}

// GetQueryWxMember returns QueryWxMemberResponse.QueryWxMember, and is useful for accessing the field via an interface.
func (v *QueryWxMemberResponse) GetQueryWxMember() data.MemberResp { return v.QueryWxMember }

// __CreateWxMemberInput is used internally by genqlient
type __CreateWxMemberInput struct {
	Input data.CreateWxMemberReq `json:"input"`
}

// GetInput returns __CreateWxMemberInput.Input, and is useful for accessing the field via an interface.
func (v *__CreateWxMemberInput) GetInput() data.CreateWxMemberReq { return v.Input }

// __QueryWxMemberInput is used internally by genqlient
type __QueryWxMemberInput struct {
	Input data.OpenIDReq `json:"input"`
}

// GetInput returns __QueryWxMemberInput.Input, and is useful for accessing the field via an interface.
func (v *__QueryWxMemberInput) GetInput() data.OpenIDReq { return v.Input }

// The query or mutation executed by CreateWxMember.
const CreateWxMember_Operation = `
mutation CreateWxMember ($input: CreateWxMemberReq!) {
	CreateWxMember(input: $input) {
		code
		wxOpenID
	}
}
`

func CreateWxMember(
	ctx_ context.Context,
	client_ graphql.Client,
	input data.CreateWxMemberReq,
) (*CreateWxMemberResponse, error) {
	req_ := &graphql.Request{
		OpName: "CreateWxMember",
		Query:  CreateWxMember_Operation,
		Variables: &__CreateWxMemberInput{
			Input: input,
		},
	}
	var err_ error

	var data_ CreateWxMemberResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by QueryWxMember.
const QueryWxMember_Operation = `
query QueryWxMember ($input: OpenIDReq!) {
	QueryWxMember(input: $input) {
		code
		wxOpenID
	}
}
`

func QueryWxMember(
	ctx_ context.Context,
	client_ graphql.Client,
	input data.OpenIDReq,
) (*QueryWxMemberResponse, error) {
	req_ := &graphql.Request{
		OpName: "QueryWxMember",
		Query:  QueryWxMember_Operation,
		Variables: &__QueryWxMemberInput{
			Input: input,
		},
	}
	var err_ error

	var data_ QueryWxMemberResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}