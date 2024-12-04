// examples/default_usage/main.go
package main

import (
	"github.com/mwiater/go-validot"
)

func main() {
	// Define required keys
	requiredKeys := []string{"API_KEY", "DB_HOST", "API_URL", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create a new validator with default settings
	validator := validot.NewValidator(validot.Config{}, requiredKeys)

	// Validate the .env file
	_ = validator.ValidateDotEnv(".env") // No need to log success here
}
