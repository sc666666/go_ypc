package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go_ypc/pkg/config"
	requests "go_ypc/pkg/request"
	"go_ypc/routes"
)

// SetupRouter 初始化 Router
func SetupRouter() *gin.Engine {
	_ = requests.InitTrans(config.GetString("app.locale"))

	// gin.Default() 使用 Gin 默认的 Logger 和 Recovery 中间件，自定义 gin.New()
	r := gin.Default()

	// 全局中间件
	// r.Use()

	routes.RegisterApiRoutes(r)

	return r
}
