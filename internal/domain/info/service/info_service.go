package service

import (
	"service/internal/infrastructure/facade"
	"service/internal/infrastructure/repository"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type InfoService struct {
	InfoRepository   repository.InfoRepository `singleton:"service/internal/infrastructure/repository.InfoRepositoryImpl"`
	KeyServiceFacade facade.KeyServiceFacade   `singleton:"service/internal/infrastructure/facade/impl.KeyServiceFacadeCachedImpl"`
}

func (i *InfoService) Save() {
	i.KeyServiceFacade.GetById(1)
}
