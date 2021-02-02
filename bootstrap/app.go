package bootstrap

import (
	"github.com/gin-gonic/gin"
)

// 应用程序设置
func SetupApp() *gin.Engine {
	// 日志
	SetupLogger()

	// 数据库
	SetupDB()

	// 路由
	return SetupRouter()
}
