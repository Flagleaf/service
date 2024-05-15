package impl

import (
	"fmt"
	"service/internal/infrastructure/facade"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type KeyServiceFacadeCachedImpl struct {
	KeyServiceFacade facade.KeyServiceFacade `singleton:"service/internal/infrastructure/facade/impl.KeyServiceFacadeImpl"`
}

func (f *KeyServiceFacadeCachedImpl) GetById(id int64) (string, error) {
	fmt.Println("cached")
	return f.KeyServiceFacade.GetById(id)
}
