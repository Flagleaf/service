package facade

type KeyServiceFacade interface {
	GetById(id int64) (string, error)
}
