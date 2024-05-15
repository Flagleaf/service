package _interface

import (
	"github.com/gin-gonic/gin"
	"service/internal/interface/handler"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type Router struct {
	InfoHandler *handler.InfoHandler `singleton:"service/internal/interface/handler.InfoHandler"`
}

func (route *Router) Setup() error {
	r := gin.Default()

	r.GET("/someJSON", route.InfoHandler.SaveInfo)

	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
