// examples/url_validation/main.go
package main

import (
	"github.com/mwiater/go-validot"
	"github.com/sirupsen/logrus"
)

func main() {
	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST"}

	// Initialize a custom logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel) // Default log level

	// Create a new validator
	validator := validot.NewValidator(validot.Config{
		RequireQuotes: true,   // Enforce that values must be quoted
		Verbose:       true,   // Enable verbose logging
		Logger:        logger, // Use the custom logger
		Plugins:       nil,    // Use built-in plugins
	}, requiredKeys)

	// Validate the .env file
	_ = validator.ValidateDotEnv(".env") // No need to log success here
}
