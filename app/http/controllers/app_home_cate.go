package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ypc/app/models/app_home_cate"
	"go_ypc/pkg/model"
	"go_ypc/pkg/response"
)

type HomeCate struct {
}

// 列表
func (*HomeCate) Index(c *gin.Context) {

	data, err := app_home_cate.GetInfo()
	if err != nil {

	}
	response.Ok(data, c)
}

// 详情
func (*HomeCate) Show(c *gin.Context) {
	id := c.Param("id")

	data, err := app_home_cate.GetInfoById(id)
	if err != nil {

	}

	response.Ok(data, c)
}

func (*HomeCate) GetHomeCate(c *gin.Context) {

}

func (*HomeCate) TestRedis(c *gin.Context) {
	err := model.RDB.Set(model.CTX, "testKey", "testValue", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := model.RDB.Get(model.CTX, "testKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("testKey:", val)
}
