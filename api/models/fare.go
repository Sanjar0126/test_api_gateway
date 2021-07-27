package models

type DeliveryPriceModel struct {
	Price float32 `json:"price"`
}

type ComputeDeliveryPriceRequestModel struct {
	BranchID string  `json:"branch_id"`
	Long     float32 `json:"long"`
	Lat      float32 `json:"lat"`
}

type ComputeDeliveryPriceResponseModel struct {
	Price    float32 `json:"price"`
	Distance float32 `json:"distance"`
}
