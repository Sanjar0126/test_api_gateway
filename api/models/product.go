package models

type GetProductModel struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	CategoryID    string   `json:"category_id"`
	Description   string   `json:"description"`
	Price         int64    `json:"price,string"`
	OrderNo       int64    `json:"order_no,string"`
	Image         string   `json:"image"`
	Title         Language `json:"title"`
	DescriptionV2 Language `json:"description_v2"`
}

type GetAllProductsModel struct {
	Products []GetProductModel `json:"products"`
	Count    int64             `json:"count,string"`
}
