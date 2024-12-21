package data

type BindCardReq struct {
	Code       string `json:"code"`
	MemberCode string `json:"memberCode"`
}

type CardResp struct {
	Code       string  `json:"code"`
	CardCode   string  `json:"cardCode"`
	MemberCode *string `json:"memberCode,omitempty"`
	Type       int     `json:"type"`
	Balance    int     `json:"balance"`
	Amount     int     `json:"amount"`
}

type CreateCardReq struct {
	Type     int     `json:"type"`
	Balance  int     `json:"balance"`
	CardCode *string `json:"cardCode,omitempty"`
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
	MemberCode string         `json:"memberCode"`
	PageInfo   *PaginationReq `json:"pageInfo,omitempty"`
}
