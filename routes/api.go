package routes

import (
	"github.com/gin-gonic/gin"
	"go_ypc/app/http/controllers"
	"go_ypc/app/http/middleware"
)

// API 路由
func RegisterApiRoutes(r *gin.Engine) {
	// V1 分组
	v1 := r.Group("/v1").Use(middleware.AddTraceId())
	{
		homeCateEndpoint := new(controllers.HomeCate)
		v1.GET("/getHomeCate", homeCateEndpoint.Index)
		v1.GET("/homeCate/:id", homeCateEndpoint.Show)
		v1.POST("/createHomeCate", homeCateEndpoint.Create)
		v1.POST("/updateHomeCate/:id", homeCateEndpoint.Update)
		v1.POST("/deleteById/:id", homeCateEndpoint.Delete)
		v1.GET("/testRedis", homeCateEndpoint.TestRedis)
		v1.POST("/testCreate", homeCateEndpoint.TestCreate)


	}
}
