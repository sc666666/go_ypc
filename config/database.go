package config

import "go_ypc/pkg/config"

func init() {
	config.Append("database", config.Configuration{

		// -----------------------------------------------------------------------------
		// Database Connections
		// -----------------------------------------------------------------------------
		"connections": map[string]interface{}{
			"mysql": map[string]interface{}{
				"host":      config.Env("DB_HOST", "127.0.0.1"),
				"port":      config.Env("DB_PORT", "3306"),
				"database":  config.Env("DB_DATABASE", "forge"),
				"username":  config.Env("DB_USERNAME", "forge"),
				"password":  config.Env("DB_PASSWORD", ""),
				"prefix":    config.Env("DB_PREFIX", ""),
				"charset":   "utf8mb4",
				"collation": "utf8mb4_unicode_ci",
				"strict":    true,

				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
		},

		// -----------------------------------------------------------------------------
		// Redis Databases
		// -----------------------------------------------------------------------------
		"redis": map[string]interface{}{
			"default": map[string]interface{}{
				"host":     config.Env("REDIS_HOST", "127.0.0.1"),
				"password": config.Env("REDIS_PASSWORD", nil),
				"port":     config.Env("REDIS_PORT", "6379"),
				"database": config.Env("REDIS_DB", "0"),
			},
		},
	})
}
