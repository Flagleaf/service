package message

import (
	"fmt"
	"service/internal/domain/info/service"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type Consumer struct {
	InfoService *service.InfoService `singleton:"service/internal/domain/info/service.InfoService"`
}

func (i *Consumer) Start() error {
	fmt.Println("init consumer.")
	i.InfoService.Save()
	return nil
}
