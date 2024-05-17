package _interface

import (
	"github.com/gin-gonic/gin"
	"service/internal/infrastructure/middleware"
	"service/internal/interface/handler"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type Router struct {
	InfoHandler *handler.InfoHandler `singleton:"service/internal/interface/handler.InfoHandler"`
}

func (route *Router) Setup() error {
	r := gin.Default()

	r.Use(middleware.ResponseMiddleware())

	r.POST("/info", route.InfoHandler.SaveInfo)
	r.DELETE("/info/:id", route.InfoHandler.Remove)
	r.PUT("/info", route.InfoHandler.Update)
	r.GET("/info/:id", route.InfoHandler.GetInfo)
	r.GET("/info", route.InfoHandler.QueryInfoList)

	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
