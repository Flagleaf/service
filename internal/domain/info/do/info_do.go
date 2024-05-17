package do

type InfoDo struct {
	Id        int64
	Name      string
	Code      string
	CreatedAt int64
	UpdatedAt int64
}

type InfoDoParam struct {
	Criteria *InfoDoParamCriteria
}

type InfoDoParamCriteria struct {
	Condition map[string]interface{}
	Between   map[string][]interface{}
}
