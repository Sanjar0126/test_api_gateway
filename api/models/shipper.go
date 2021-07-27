package models

type GetShipperModel struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	UserName      string   `json:"username"`
	Password      string   `json:"password"`
	Phone         []string `json:"phone"`
	Logo          string   `json:"logo"`
	Description   string   `json:"description"`
	IsActive      bool     `json:"is_active"`
	CreatedAt     string   `json:"created_at"`
	CallCenterTg  string   `json:"call_center_tg" example:"@admin"`
	WorkHourStart string   `json:"work_hour_start" example:"05:00"`
	WorkHourEnd   string   `json:"work_hour_end" example:"00:00"`
	MenuImage     string   `json:"manu_image"`
}
