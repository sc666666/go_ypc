package bootstrap

import (
	"go_ypc/pkg/config"
	"go_ypc/pkg/model"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	// 连接数据库
	db := model.ConnectDB()

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.connections.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.connections.mysql.max_idle_connections"))
	// 设置每个链接过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.connections.mysql.max_life_seconds")) * time.Second)
}
