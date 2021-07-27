package models

type Point struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type GetGeozoneModel struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Points []Point `json:"points"`
}

type GetAllGeozonesModel struct {
	Count    int               `json:"count"`
	Geozones []GetGeozoneModel `json:"geozones"`
}

type CreateGeozoneModel struct {
	Name   string  `json:"name"`
	Points []Point `json:"points"`
}

type UpdateGeozoneModel struct {
	Name   string  `json:"name"`
	Points []Point `json:"points"`
}

type CheckPointInsideGeozoneRequest struct {
	Point Point `json:"point"`
}

type CheckPointInsideGeozoneResponse struct {
	Result bool `json:"result"`
}
