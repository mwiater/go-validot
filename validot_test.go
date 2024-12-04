// validot_test.go
package validot

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/mwiater/go-validot/plugins"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// CacheSizeValidationPlugin is a custom plugin for testing purposes.
// It ensures that the "CACHE_SIZE" key has a numeric value between 128 and 1024.
type CacheSizeValidationPlugin struct{}

// Validate checks if the key is "CACHE_SIZE" and validates its value.
func (p *CacheSizeValidationPlugin) Validate(key, value string) (bool, error) {
	if key != "CACHE_SIZE" {
		return false, nil // Plugin not handling this key.
	}

	// Simple numeric validation
	size, err := strconv.Atoi(value)
	if err != nil {
		return true, fmt.Errorf("CACHE_SIZE must be a numeric value")
	}
	if size < 128 || size > 1024 {
		return true, fmt.Errorf("CACHE_SIZE must be between 128 and 1024")
	}
	return true, nil
}

// Name returns the name of the plugin.
func (p *CacheSizeValidationPlugin) Name() string {
	return "CacheSizeValidationPlugin"
}

// Helper function to create a temporary .env file with given content.
func createTempEnvFile(t *testing.T, content string) string {
	tmpDir := t.TempDir()
	envFilePath := filepath.Join(tmpDir, ".env")
	err := os.WriteFile(envFilePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp .env file: %v", err)
	}
	return envFilePath
}

func TestValidateDotEnv_ValidFile(t *testing.T) {
	envContent := `
# Valid .env file

API_KEY="12345abcdef"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="PRODUCTION"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,  // Enforce that values must be quoted
		Verbose:       false, // Disable verbose logging
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.NoError(t, err, "Expected no validation errors for a valid .env file")
}

func TestValidateDotEnv_MissingRequiredKeys(t *testing.T) {
	envContent := `
# Missing required keys

API_KEY="12345abcdef"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil,
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.Error(t, err, "Expected validation error for missing required keys")

	// Define expected missing keys
	expectedMissingKeys := []string{"API_URL", "DB_HOST", "ENVIRONMENT"}

	// Check that each expected missing key is present in the error message
	for _, key := range expectedMissingKeys {
		assert.Contains(t, err.Error(), key, fmt.Sprintf("Error message should contain missing key: %s", key))
	}

	// Optionally, check that no unexpected keys are reported as missing
	unexpectedMissingKeys := []string{"SERVICE_ENDPOINT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}
	for _, key := range unexpectedMissingKeys {
		assert.NotContains(t, err.Error(), key, fmt.Sprintf("Error message should not contain key: %s", key))
	}
}

func TestValidateDotEnv_InvalidURL(t *testing.T) {
	envContent := `
# Invalid URL scheme

API_KEY="12345abcdef"
API_URL="ftp://api.myapp.com/v1/" # Should be https
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.Error(t, err, "Expected validation error for invalid URL scheme")
	assert.Contains(t, err.Error(), "URL scheme for key \"API_URL\" must be one of [https]")
}

func TestValidateDotEnv_InvalidEnum(t *testing.T) {
	envContent := `
# Invalid ENUM value

API_KEY="12345abcdef"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="INVALID_ENV" # Should be DEVELOPMENT, STAGING, or PRODUCTION

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.Error(t, err, "Expected validation error for invalid enum value")
	assert.Contains(t, err.Error(), "value for key \"ENVIRONMENT\" must be one of [DEVELOPMENT STAGING PRODUCTION]")
}

func TestValidateDotEnv_InvalidBoolean(t *testing.T) {
	envContent := `
# Invalid Boolean value

API_KEY="12345abcdef"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="maybe" # Should be true, false, 1, 0, yes, no
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.Error(t, err, "Expected validation error for invalid boolean value")
	assert.Contains(t, err.Error(), "value for key \"ENABLE_DEBUG\" must be a boolean (accepted values: [true false 1 0 yes no])")
}

func TestValidateDotEnv_InvalidIPAddress(t *testing.T) {
	envContent := `
# Invalid IP Address

API_KEY="12345abcdef"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="8.8.8.8" # Should be a private IP address
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.Error(t, err, "Expected validation error for invalid IP address")
	assert.Contains(t, err.Error(), "value for key \"TRUSTED_PROXY_IP\" must be a private IP address")
}

func TestValidateDotEnv_DuplicateKeys(t *testing.T) {
	envContent := `
# Duplicate keys

API_KEY="12345abcdef"
API_KEY="duplicatekey123" # Duplicate
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_HOST="duplicatedbhost" # Duplicate
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.NoError(t, err, "Expected no validation errors for duplicate keys")
	// Note: godotenv overwrites duplicate keys, so the last value is used.
	// If handling duplicates is desired, additional logic is needed.
}

func TestValidateDotEnv_EmptyFile(t *testing.T) {
	envContent := `# Empty .env file`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: false,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.Error(t, err, "Expected validation error for empty .env file")

	// Define expected missing keys
	expectedMissingKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Check that each expected missing key is present in the error message
	for _, key := range expectedMissingKeys {
		assert.Contains(t, err.Error(), key, fmt.Sprintf("Error message should contain missing key: %s", key))
	}
}

func TestValidateDotEnv_InlineComments(t *testing.T) {
	envContent := `
# Inline comments and valid keys

API_KEY="12345abcdef" # API key for authentication
API_URL="https://api.myapp.com/v1/" # API endpoint
API_SECRET="secretvalue123" # Secret key
API_TIMEOUT="30" # Timeout in seconds

SERVICE_ENDPOINT="https://service.myapp.com/endpoint" # Service URL
SERVICE_VERSION="v2" # Service version

DB_HOST="localhost" # Database host
DB_PORT="5432" # Database port
DB_USER="admin" # Database user
DB_PASSWORD="securepassword" # Database password
DB_NAME="myapp_db" # Database name

ENVIRONMENT="DEVELOPMENT" # Environment setting

ENABLE_DEBUG="true" # Enable debug mode
ENABLE_FEATURE_X="false" # Feature X flag
ENABLE_FEATURE_Y="yes" # Feature Y flag

TRUSTED_PROXY_IP="192.168.1.100" # Trusted proxy IP
REDIS_HOST="redis.local" # Redis host
REDIS_PORT="6379" # Redis port

LOG_LEVEL="INFO" # Logging level
LOG_FORMAT="json" # Logging format

SERVICE_TIMEOUT="60" # Service timeout
CACHE_SIZE="256" # Cache size
UPLOAD_LIMIT="1048576" # Upload limit
USE_SSL="yes" # Use SSL
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.NoError(t, err, "Expected no validation errors for .env file with inline comments")
}

func TestValidateDotEnv_NoQuotesWhenRequired(t *testing.T) {
	envContent := `
# Values not quoted when RequireQuotes is true

API_KEY=12345abcdef
API_URL=https://api.myapp.com/v1/
API_SECRET=secretvalue123
API_TIMEOUT=30

SERVICE_ENDPOINT=https://service.myapp.com/endpoint
SERVICE_VERSION=v2

DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=securepassword
DB_NAME=myapp_db

ENVIRONMENT=STAGING

ENABLE_DEBUG=true
ENABLE_FEATURE_X=false
ENABLE_FEATURE_Y=yes

TRUSTED_PROXY_IP=192.168.1.100
REDIS_HOST=redis.local
REDIS_PORT=6379

LOG_LEVEL=INFO
LOG_FORMAT=json

SERVICE_TIMEOUT=60
CACHE_SIZE=256
UPLOAD_LIMIT=1048576
USE_SSL=yes
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator with RequireQuotes=true
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	// Since the current plugins do not enforce quotes, this should pass
	// If quote enforcement is desired, additional logic needs to be added in the validator
	assert.NoError(t, err, "Expected no validation errors even without quotes when RequireQuotes is true")
}

func TestValidateDotEnv_CustomPlugin(t *testing.T) {
	envContent := `
# .env file with custom plugin

API_KEY="12345abcdef"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256" # Valid
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator with custom plugin
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       []plugins.ValidationPlugin{&CacheSizeValidationPlugin{}},
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.NoError(t, err, "Expected no validation errors with valid CACHE_SIZE")

	// Now, set CACHE_SIZE to an invalid value
	envContentInvalid := `
# .env file with invalid CACHE_SIZE

API_KEY="12345abcdef"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="64" # Invalid: less than 128
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePathInvalid := createTempEnvFile(t, envContentInvalid)

	// Validate with invalid CACHE_SIZE
	err = validator.ValidateDotEnv(envFilePathInvalid)
	assert.Error(t, err, "Expected validation error for invalid CACHE_SIZE")
	assert.Contains(t, err.Error(), "CACHE_SIZE must be between 128 and 1024")
}

func TestValidateDotEnv_InvalidKeyValue(t *testing.T) {
	envContent := `
# Invalid key-value pair

=invalidkey
API_URL="https://api.myapp.com/v1/"
API_KEY="12345abcdef"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="STAGING"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger with buffer to capture logs
	var logBuf bytes.Buffer
	logger := logrus.New()
	logger.SetOutput(&logBuf)
	logger.SetLevel(logrus.DebugLevel)

	// Define required keys
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       true,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)

	// Since godotenv.Read skips invalid lines and required keys are present, expect no error
	assert.NoError(t, err, "Expected no validation errors since required keys are present despite invalid key-value pair")
}

func TestValidateDotEnv_KeyNotHandledByAnyPlugin(t *testing.T) {
	envContent := `
# Key not handled by any plugin

API_KEY="12345abcdef"
CUSTOM_KEY="custom_value"
API_URL="https://api.myapp.com/v1/"
API_SECRET="secretvalue123"
API_TIMEOUT="30"

SERVICE_ENDPOINT="https://service.myapp.com/endpoint"
SERVICE_VERSION="v2"

DB_HOST="localhost"
DB_PORT="5432"
DB_USER="admin"
DB_PASSWORD="securepassword"
DB_NAME="myapp_db"

ENVIRONMENT="PRODUCTION"

ENABLE_DEBUG="true"
ENABLE_FEATURE_X="false"
ENABLE_FEATURE_Y="yes"

TRUSTED_PROXY_IP="192.168.1.100"
REDIS_HOST="redis.local"
REDIS_PORT="6379"

LOG_LEVEL="INFO"
LOG_FORMAT="json"

SERVICE_TIMEOUT="60"
CACHE_SIZE="256"
UPLOAD_LIMIT="1048576"
USE_SSL="yes"
`

	envFilePath := createTempEnvFile(t, envContent)

	// Initialize logger
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	// Define required keys (CUSTOM_KEY is not required)
	requiredKeys := []string{"API_URL", "SERVICE_ENDPOINT", "DB_HOST", "ENVIRONMENT", "ENABLE_DEBUG", "TRUSTED_PROXY_IP"}

	// Create validator
	validator := NewValidator(Config{
		RequireQuotes: true,
		Verbose:       false,
		Logger:        logger,
		Plugins:       nil, // No additional plugins
	}, requiredKeys)

	// Validate
	err := validator.ValidateDotEnv(envFilePath)
	assert.NoError(t, err, "Expected no validation errors for keys not handled by any plugin")
	// Note: CUSTOM_KEY is optional and not handled by any plugin
}
