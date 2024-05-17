package service

import (
	"service/internal/domain/info/do"
	"service/internal/infrastructure/facade"
	"service/internal/infrastructure/repository"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type InfoService struct {
	InfoRepository   repository.InfoRepository `singleton:"service/internal/infrastructure/repository.InfoRepositoryImpl"`
	KeyServiceFacade facade.KeyServiceFacade   `singleton:"service/internal/infrastructure/facade.KeyServiceFacadeCachedImpl"`
}

func (i *InfoService) Save(do do.InfoDo) {
	i.InfoRepository.Insert(do)
}

func (i *InfoService) Delete(id int64) {
	i.InfoRepository.DeleteByPrimaryKey(id)
}

func (i *InfoService) GetInfo(id int64) do.InfoDo {
	return i.InfoRepository.SelectByPrimaryKey(id)
}

func (i *InfoService) QueryInfoList() []do.InfoDo {
	return i.InfoRepository.SelectByParam(do.InfoDoParam{})
}
