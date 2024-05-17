package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	do2 "service/internal/domain/info/do"
	"service/internal/domain/info/service"
	error2 "service/internal/infrastructure/error"
	"service/internal/interface/dto"
	"strconv"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type InfoHandler struct {
	InfoService *service.InfoService `singleton:"service/internal/domain/info/service.InfoService"`
}

func (i *InfoHandler) SaveInfo(c *gin.Context) {
	var createReq dto.InfoCreateReq
	if err := c.ShouldBindJSON(&createReq); err != nil {
		c.AbortWithError(http.StatusOK, err)
		return
	}
	do := do2.InfoDo{}
	i.InfoService.Save(do)
	c.Status(http.StatusOK)
}

func (i *InfoHandler) Remove(c *gin.Context) {
	id := c.Param("id")
	iid, _ := strconv.ParseInt(id, 0, 0)
	i.InfoService.Delete(iid)
	c.Status(http.StatusOK)
}

func (i *InfoHandler) Update(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (i *InfoHandler) GetInfo(c *gin.Context) {
	do := i.InfoService.GetInfo(1)
	c.Status(http.StatusOK)
	c.Set("response", do)
}

func (i *InfoHandler) QueryInfoList(c *gin.Context) {
	if 11 == 12 {
		c.AbortWithError(http.StatusOK, error2.NewServiceError(1, ""))
		return
	}
	dos := i.InfoService.QueryInfoList()
	c.Status(http.StatusOK)
	c.Set("response", dos)
}
