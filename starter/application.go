package starter

import (
	"service/internal/application/message"
	_interface "service/internal/interface"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type Application struct {
	Router       *_interface.Router `singleton:"service/internal/interface.Router"`
	InitConsumer *message.Consumer  `singleton:"service/internal/application/message.Consumer"`
}

func (a *Application) Run() {
	err := a.Router.Setup()
	if err != nil {
		return
	}
	err = a.InitConsumer.Start()
	if err != nil {
		return
	}
}
