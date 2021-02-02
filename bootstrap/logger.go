package bootstrap

import (
	"go_ypc/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(true, false, true)
}
