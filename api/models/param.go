package models

type GetAllParams struct {
	Limit      int    `json:"limit" binding:"required" default:"10"`
	Page       int    `json:"page" binding:"required" default:"1"`
	Search     string `json:"search"`
	SortByDate string `json:"sort_by_date" binding:"required,oneof=asc desc none"`
	SortByName string `json:"sort_by_name" binding:"required,oneof=asc desc none"`
}

type GetAllResponse struct {
	Students []*Student `json:"students"`
	Count    int       `json:"count"`
}
