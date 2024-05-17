package dto

type PageReq struct {
	PageIndex int64 `json:"page_index"`
	PageSize  int64 `json:"page_size"`
}

type PageResp[T any] struct {
	PageIndex int64 `json:"page_index"`
	PageSize  int64 `json:"page_size"`
	TotalPage int64 `json:"total_page"`
	Total     int64 `json:"total"`
	DataList  []T   `json:"data_list"`
}
