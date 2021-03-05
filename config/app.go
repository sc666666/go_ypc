package config

import "go_ypc/pkg/config"

func init() {
	config.Append("app", config.Configuration{

		// -----------------------------------------------------------------------------
		// Application Name
		// -----------------------------------------------------------------------------
		// This value is the name of your application.
		"name": config.Env("APP_NAME", "Go"),

		// -----------------------------------------------------------------------------
		// Application Environment
		// -----------------------------------------------------------------------------
		// This value determines the "environment" in which your application is
		// currently running.
		"env": config.Env("APP_ENV", "production"),

		// -----------------------------------------------------------------------------
		// Application Debug Mode
		// -----------------------------------------------------------------------------
		// If enabled, a detailed error message will be displayed on every error
		// that occurs within your application.
		"debug": config.Env("APP_DEBUG", false),

		// -----------------------------------------------------------------------------
		// Application URL
		// -----------------------------------------------------------------------------
		// The URL is used by the application.
		"url": config.Env("APP_URL", "http://localhost"),

		// -----------------------------------------------------------------------------
		// Application Port
		// -----------------------------------------------------------------------------
		// This port is used by the application listening.
		"port": config.Env("APP_PORT", "80"),

		// -----------------------------------------------------------------------------
		// Encryption Key
		// -----------------------------------------------------------------------------
		// This key is used by the encrypt service.
		"key": config.Env("APP_KEY", ""),

		// -----------------------------------------------------------------------------
		// Application Locale
		// -----------------------------------------------------------------------------
		"locale": config.Env("LOCALE", "zh"),

		// -----------------------------------------------------------------------------
		// Application Fallback Locale
		// -----------------------------------------------------------------------------
		"fallback_locale": config.Env("FALLBACK_LOCALE", "en"),
	})
}
