package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/internal/domain/info/service"
	"service/internal/interface/vo"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton
type InfoHandler struct {
	InfoService *service.InfoService `singleton:"service/internal/domain/info/service.InfoService"`
}

func (i *InfoHandler) SaveInfo(c *gin.Context) {
	i.InfoService.Save()
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}
	c.AsciiJSON(http.StatusOK, data)
}

func (i *InfoHandler) GetInfo() *vo.InfoVo {
	return nil
}

func (i *InfoHandler) QueryInfoList() {

}
