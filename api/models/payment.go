package models

type GenerateLinkRequestModel struct {
	Amount   int64 `json:"amount"`
	OrderNum int64 `json:"order_num"`
}

type GenerateLinkResponseModel struct {
	Link string `json:"link"`
}
