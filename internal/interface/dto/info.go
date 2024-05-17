package dto

type InfoCreateReq struct {
	Name string
	Code string
}

type InfoUpdateReq struct {
	Id   int64
	Name string
	Code string
}

type InfoQueryParam struct {
	Name string
	Code string
}
