package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_ypc/app/http/requests"
	"go_ypc/app/models/app_home_cate"
	"go_ypc/pkg/model"
	"go_ypc/pkg/response"
	"strconv"
)

type HomeCate struct {
}

// 列表
func (*HomeCate) Index(c *gin.Context) {
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	data, err := app_home_cate.GetInfo(size, page)
	if err != nil {

	}
	// 格式化
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

//创建数据
func (h *HomeCate) Create(c *gin.Context) {
	var appHomeCate app_home_cate.AppHomeCate
	//appHomeCate = []*app_home_cate.AppHomeCate{}
	err := c.ShouldBindJSON(&appHomeCate)
	//err := c.BindJSON(&appHomeCate)
	fmt.Print(err)
	//c.ShouldBindBodyWith(&appHomeCate)
	data, err := app_home_cate.CreateInfo(&appHomeCate)
	if err != nil {

	}
	response.Ok(data, c)
}

//更新数据
//func (h *HomeCate) Update(c *gin.Context) {
//	id := c.Param("id")
//	var appHomeCate app_home_cate.AppHomeCate
//	err := c.ShouldBindJSON(&appHomeCate)
//
//	data, err := app_home_cate.SaveInfo(&appHomeCate, id)
//	if err != nil {
//
//	}
//
//	response.Ok(data, c)
//}

////删除数据
//func (*HomeCate) Delete(c *gin.Context) {
//	id := c.Param("id")
//
//	data, err := app_home_cate.GetInfo(size, page)
//	if err != nil {
//
//	}
//
//	response.Ok(data, c)
//}

func (*HomeCate) TestRedis(c *gin.Context) {
	err := model.RDB.Set(model.CTX, "testKey", "testValue", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := model.RDB.Get(model.CTX, "testKey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("testKey :", val)
}

func (*HomeCate) TestCreate(c *gin.Context) {
	form := requests.ValidateHomeCate(c)
	fmt.Println(form.Title)
	// app_home_cate2.ValidateStruct(c)
}
