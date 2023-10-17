package model

type PageInfo struct {
	PageSize int         `json:"pageSize"`
	PageNum  int         `json:"pageNum"`
	Total    int         `json:"total"`
	LastPage int         `json:"lastPage"`
	Order    string      `json:"order"`
	Offset   int         `json:"offset"`
	Data     interface{} `json:"data"`
}
