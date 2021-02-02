package main

import (
	"go_ypc/bootstrap"
	"go_ypc/config"
)

//go:generate echo arch=$GOARCH os=$GOOS pkg=$GOPACKAGE file=$GOFILE line=$GOLINE

// 初始化配置信息
func init() {
	config.Initialize()
}

func main() {
	// 设置应用引导服务
	r := bootstrap.SetupApp()

	// 监听并在 .env 配置的端口上启动服务
	_ = r.Run(config.GetAppPort())
}
