package models

type GetBranchModel struct {
	ID          string   `json:"id"`
	ShipperID   string   `json:"shipper_id"`
	Name        string   `json:"name"`
	Phone       string   `json:"phone"`
	Address     string   `json:"address"`
	Destination string   `json:"destination"`
	Image       string   `json:"image"`
	WorkHours   string   `json:"work_hours" example:"9am-10pm"`
	Location    Location `json:"location"`
	IsActive    bool     `json:"is_active"`
	CreatedAt   string   `json:"created_at"`
}

type GetAllBranchesModel struct {
	Count    int              `json:"count"`
	Branches []GetBranchModel `json:"branches"`
}
