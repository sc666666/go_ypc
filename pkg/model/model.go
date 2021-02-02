package model

import (
	"fmt"
	"go_ypc/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	// var err error
	var (
		host     = config.GetString("database.connections.mysql.host")
		port     = config.GetString("database.connections.mysql.port")
		database = config.GetString("database.connections.mysql.database")
		username = config.GetString("database.connections.mysql.username")
		password = config.GetString("database.connections.mysql.password")
		charset  = config.GetString("database.connections.mysql.charset")
	)

	// 想要正确的处理 time.Time ，您需要带上 parseTime 参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})

	var level logger.LogLevel
	if config.GetBool("app.debug") {
		// 开启了 debug 模式
		level = logger.Info
	} else {
		// 关闭了 debug 模式
		level = logger.Error
	}

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second,   // 慢 SQL 阈值
	// 		LogLevel:      logger.Silent, // Log level
	// 		Colorful:      false,         // 禁用彩色打印
	// 	},
	// )

	DB, _ = gorm.Open(gormConfig, &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})

	return DB
}
