package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/link"
	"gohub/pkg/response"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	//links := link.All() //从数据库拉数据
	//response.Data(c, links)
	response.Data(c, link.AllCached()) //从cache中拉数据
}
