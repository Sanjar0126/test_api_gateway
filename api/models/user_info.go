package models

type UserInfo struct {
	ID         string `json:"id"`
	UserType   string `json:"role"`
	UserID     string `json:"user_id"`
	ShipperID  string `json:"shipper_id"`
	UserTypeID string `json:"user_type_id"`
}
