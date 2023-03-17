package models

type PageInfo struct {
	Total    int64       `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	Data     interface{} `json:"data"`
}
