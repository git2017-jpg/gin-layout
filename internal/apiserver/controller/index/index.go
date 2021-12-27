package index

import (
	"github.com/BooeZhang/gin-layout/internal/apiserver/service"
	"github.com/BooeZhang/gin-layout/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	srv service.Service
}

func NewIndexController(store store.Factory) *Controller {
	return &Controller{
		srv: service.NewService(store),
	}
}

func (i *Controller) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (i *Controller) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
