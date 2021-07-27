package models

type GetCategoryModel struct {
	ID              string               `json:"id"`
	Name            string               `json:"name"`
	Slug            string               `json:"slug"`
	ParentID        string               `json:"parent_id"`
	Description     string               `json:"description"`
	ChildCategories []ChildCategoryModel `json:"child_categories"`
	Image           string               `json:"image"`
	OrderNo         int64                `json:"order_no,string"`
	Products        []GetProductModel    `json:"products"`
	Title           Language             `json:"title"`
	DescriptionV2   Language             `json:"description_v2"`
}

type ChildCategoryModel struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	ParentID      string            `json:"parent_id"`
	Description   string            `json:"description"`
	Image         string            `json:"image"`
	OrderNo       int64             `json:"order_no,string"`
	Products      []GetProductModel `json:"products"`
	Title         Language          `json:"title"`
	DescriptionV2 Language          `json:"description_v2"`
}

type GetAllCategoriesModel struct {
	Categories []GetCategoryModel `json:"categories"`
	Count      int64              `json:"count,string"`
}
