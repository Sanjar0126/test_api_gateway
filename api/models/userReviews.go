package models

type CreateUserReviewsModel struct {
	ID             string `json:"id"`
	ReviewMessage  string `json:"review_message"`
	Type           string `json:"type"`
	RelatedSubject string `json:"related_subject"`
	Lang           string `json:"lang"`
	ReviewId       string `json:"review_id"`
	OrderNum       string `json:"order_num"`
	DeliveryTime   int64  `json:"delivery_time"`
	CourierName    string `json:"courier_name"`
	CourierId      string `json:"courier_id"`
	OperatorName   string `json:"operator_name"`
	OperatorId     string `json:"operator_id"`
	ClientName     string `json:"client_name"`
	ClientPhone    string `json:"client_phone"`
	ClientId       string `json:"client_id"`
	CreatedAt      string `json:"created_at"`
	BranchId       string `json:"branch_id"`
	BranchName     string `json:"branch_name"`
}
