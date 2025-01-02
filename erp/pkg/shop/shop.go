package shop

type QryShopReq struct {
	ShopCode string `json:"shopCode"`
}

type ShopResp struct {
	MallCode      string `json:"mallCode"`
	MallName      string `json:"mallName"`
	ContractCode  string `json:"contractCode"`
	PosCode       string `json:"posCode"`
	ShopCode      string `json:"shopCode"`
	ShopName      string `json:"shopName"`
	BizClass1     string `json:"bizClass1"`
	BizClassName1 string `json:"bizClassName1"`
	BizClass2     string `json:"bizClass2"`
	BizClassName2 string `json:"bizClassName2"`
	Status        int32  `json:"status"`
}
