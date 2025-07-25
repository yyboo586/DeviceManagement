package model

var (
	DefaultPageSize = 10
)

type Author struct {
	Authorization string `p:"Authorization" v:"required" in:"header" dc:"Bearer {{token}}"`
}

type PageReq struct {
	Page     int `json:"page" dc:"页码"`
	PageSize int `json:"page_size" dc:"每页条数"`
}

type PageRes struct {
	Total       int `json:"total" dc:"总条数"`
	CurrentPage int `json:"current_page" dc:"当前页码"`
}
