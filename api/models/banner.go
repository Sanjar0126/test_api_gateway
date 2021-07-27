package models

type CreateBannerModel struct {
	Title    Language `json:"title"`
	Image    string   `json:"image"`
	Position string   `json:"position"`
}

type UpdateBannerModel struct {
	Title    Language `json:"title"`
	Image    string   `json:"image"`
	Position string   `json:"position"`
	Active   bool     `json:"active"`
}

type GetBannerModel struct {
	ID         string   `json:"id"`
	Title      Language `json:"title"`
	Image      string   `json:"image"`
	Position   string   `json:"position"`
	Active     bool     `json:"active"`
	Url        string   `json:"url"`
	Created_at string   `json:"created_at"`
	Updated_at string   `json:"updated_at"`
	Shipper_id string   `json:"shipper_id"`
}

type GetAllBannersModel struct {
	Banners []GetBannerModel `json:"banners"`
	Count   int64            `json:"count"`
}
