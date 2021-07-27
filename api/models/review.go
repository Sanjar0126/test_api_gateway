package models

type GetReviewModel struct {
	ID             string   `json:"id"`
	RelatedSubject string   `json:"related_subject"`
	Message        Language `json:"message"`
	Type           string   `json:"type"`
	Active         bool     `json:"active"`
	Created_at     string   `json:"created_at"`
	Updated_at     string   `json:"updated_at"`
	Shipper_id     string   `json:"shipper_id"`
}
type GetAllReviewsModel struct {
	Reviews []GetReviewModel `json:"reviews"`
	Count   int64            `json:"count"`
}
