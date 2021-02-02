package config

import "go_ypc/pkg/config"

// Configuration information initialization
func Initialize() {
	// Triggered to load the init method into other files in this directory.
}

// Get application port
func GetAppPort() string {
	return ":" + config.GetString("app.port")
}
