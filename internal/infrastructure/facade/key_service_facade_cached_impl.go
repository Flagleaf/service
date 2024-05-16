package facade

import (
	"fmt"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type KeyServiceFacadeCachedImpl struct {
	KeyServiceFacade KeyServiceFacade `singleton:"service/internal/infrastructure/facade.KeyServiceFacadeImpl"`
}

func (f *KeyServiceFacadeCachedImpl) GetById(id int64) (string, error) {
	fmt.Println("cached")
	return f.KeyServiceFacade.GetById(id)
}
