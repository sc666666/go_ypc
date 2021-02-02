package controllers

import (
	"github.com/gin-gonic/gin"
	"go_ypc/app/models/app_home_cate"
	"go_ypc/pkg/response"
)

type HomeCate struct {
}

// 列表
func (*HomeCate) Index(c *gin.Context) {

}

// 详情
func (*HomeCate) Show(c *gin.Context) {
	id := c.Param("id")

	data, err := app_home_cate.GetInfoById(id)
	if err != nil {

	}

	response.Ok(data, c)
}
