package facade

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type KeyServiceFacadeImpl struct {
}

func (f *KeyServiceFacadeImpl) GetById(id int64) (string, error) {
	return "key", nil
}
