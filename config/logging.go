package config

import "go_ypc/pkg/config"

func init() {
	config.Append("logging", config.Configuration{

		// -----------------------------------------------------------------------------
		// Default Log Channel
		// -----------------------------------------------------------------------------
		"default": config.Env("LOG_CHANNEL", "daily"),

		// -----------------------------------------------------------------------------
		// Log Channels
		// -----------------------------------------------------------------------------
		"channels": map[string]interface{}{
			"daily": map[string]interface{}{
				"path":   config.Env("LOG_PATH", "./runtime/logs"),
				"format": "2006-01-02",
				"ext":    "log",
			},
		},

		// -----------------------------------------------------------------------------
		// Log Cutting
		// -----------------------------------------------------------------------------
		"cutting": map[string]interface{}{
			"max_size":    1,
			"max_backups": 5,
			"max_age":     30,
			"compress":    false,
		},
	})
}
