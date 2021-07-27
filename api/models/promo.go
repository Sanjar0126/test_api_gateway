package models

type Language struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
	En string `json:"en"`
}

type GetPromoModel struct {
	Title       Language `json:"title"`
	Description Language `json:"description"`
	IsActive    bool     `json:"is_active"`
	Image       string   `json:"image"`
	StartTime   string   `json:"start_time"`
	EndTime     string   `json:"end_time"`
	CreatedAt   string   `json:"created_at"`
}

type GetAllPromosModel struct {
	Promos []GetPromoModel `json:"promos"`
	Count  int64           `json:"count"`
}
