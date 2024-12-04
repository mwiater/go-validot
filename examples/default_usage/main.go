// examples/default_usage/main.go
package main

import (
	"github.com/mwiater/go-validot"
	"github.com/sirupsen/logrus"
)

func main() {
	// Define required keys
	requiredKeys := []string{"API_KEY", "DB_HOST", "API_URL", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Initialize a custom logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel) // Default log level

	// Create a new validator with default settings
	validator := validot.NewValidator(validot.Config{
		RequireQuotes: false,  // Do not enforce that values must be quoted
		Verbose:       false,  // Disable verbose logging
		Logger:        logger, // Use the custom logger
		Plugins:       nil,    // No additional plugins; built-in plugins are included automatically
	}, requiredKeys)

	// Validate the .env file
	_ = validator.ValidateDotEnv(".env") // No need to log success here
}
