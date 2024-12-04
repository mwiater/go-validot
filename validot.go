package validot

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mwiater/go-validot/plugins"
	"github.com/sirupsen/logrus"
)

// Validator is responsible for validating `.env` files based on the provided configuration.
// It includes required keys, configuration options, and validation plugins.
type Validator struct {
	config       Config                     // Configuration settings for the Validator.
	requiredKeys map[string]bool            // A map of keys that are required in the `.env` file.
	plugins      []plugins.ValidationPlugin // List of validation plugins to apply to the `.env` file.
}

// NewValidator initializes and returns a new Validator instance.
//
// Parameters:
//   - config: The configuration settings for the Validator.
//   - requiredKeys: A slice of strings specifying the keys that must be present in the `.env` file.
//
// Returns:
//   - *Validator: A pointer to a newly created Validator instance.
func NewValidator(config Config, requiredKeys []string) *Validator {
	reqKeys := make(map[string]bool)
	for _, key := range requiredKeys {
		reqKeys[key] = false
	}

	builtInPlugins := loadBuiltInPlugins()
	allPlugins := append(builtInPlugins, config.Plugins...)

	return &Validator{
		config:       config,
		requiredKeys: reqKeys,
		plugins:      allPlugins,
	}
}

// loadBuiltInPlugins initializes and returns a slice of built-in validation plugins.
//
// Returns:
//   - []plugins.ValidationPlugin: A list of preconfigured validation plugins.
func loadBuiltInPlugins() []plugins.ValidationPlugin {
	var builtIn []plugins.ValidationPlugin

	urlPlugin := &plugins.URLValidationPlugin{
		Key:            "API_URL",
		AllowedSchemes: []string{"https"},
	}
	builtIn = append(builtIn, urlPlugin)

	enumPlugin := &plugins.EnumValidationPlugin{
		Key:           "ENVIRONMENT",
		AllowedValues: []string{"DEVELOPMENT", "STAGING", "PRODUCTION"},
		CaseSensitive: true,
	}
	builtIn = append(builtIn, enumPlugin)

	boolPlugin := &plugins.BooleanValidationPlugin{
		Key:            "ENABLE_DEBUG",
		AcceptedValues: []string{"true", "false", "1", "0", "yes", "no"},
		Standardize:    true,
	}
	builtIn = append(builtIn, boolPlugin)

	ipPlugin := &plugins.IPAddressValidationPlugin{
		Key:               "TRUSTED_PROXY_IP",
		AllowedIPVersions: []string{"IPv4", "IPv6"},
		MustBePrivate:     true,
	}
	builtIn = append(builtIn, ipPlugin)

	return builtIn
}

// ValidateDotEnv validates the `.env` file at the specified path using the Validator's configuration.
//
// Parameters:
//   - filePath: The path to the `.env` file to validate.
//
// Returns:
//   - error: An error if validation fails, or nil if the `.env` file is valid.
func (v *Validator) ValidateDotEnv(filePath string) error {
	if v.config.Logger == nil {
		v.config.Logger = logrus.New()
		if v.config.Verbose {
			v.config.Logger.SetLevel(logrus.DebugLevel)
		} else {
			v.config.Logger.SetLevel(logrus.InfoLevel)
		}
	}

	if v.config.Verbose {
		v.config.Logger.Infof("Validator Configuration:")
		v.config.Logger.Infof("  RequireQuotes: %v", v.config.RequireQuotes)
		v.config.Logger.Infof("  Verbose: %v", v.config.Verbose)
		v.config.Logger.Infof("  Number of Plugins: %d", len(v.plugins))
		v.config.Logger.Infof("End of Configuration")
	}

	v.config.Logger.Infof("Starting validation for file: %s", filePath)

	envVars, err := loadEnvFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}

	for key, value := range envVars {
		if v.config.Verbose {
			v.config.Logger.Infof("Processing key: %s", key)
		}

		if _, exists := v.requiredKeys[key]; exists {
			v.requiredKeys[key] = true
			if v.config.Verbose {
				v.config.Logger.Infof("  %s is a required variable.", key)
			}
		} else {
			if v.config.Verbose {
				v.config.Logger.Infof("  %s is an optional variable.", key)
			}
		}

		for _, plugin := range v.plugins {
			handled, err := plugin.Validate(key, value)
			if err != nil {
				if v.config.Verbose {
					v.config.Logger.Errorf("Validation error for key %s by %s: %v", key, plugin.Name(), err)
				} else {
					v.config.Logger.Errorf("Validation error for key %s: %v", key, err)
				}
				return err
			}
			if handled && v.config.Verbose {
				v.config.Logger.Infof("  [Validated by: %s]", plugin.Name())
			}
		}
	}

	missingKeys := []string{}
	for key, found := range v.requiredKeys {
		if !found {
			missingKeys = append(missingKeys, key)
		}
	}

	if len(missingKeys) > 0 {
		errMsg := fmt.Sprintf("missing required keys: %v", missingKeys)
		v.config.Logger.Error(errMsg)
		return fmt.Errorf(errMsg)
	}

	v.config.Logger.Infof(".env file is valid.")
	return nil
}

// loadEnvFile reads and parses the `.env` file from the specified path.
//
// Parameters:
//   - filePath: The path to the `.env` file.
//
// Returns:
//   - map[string]string: A map containing the key-value pairs from the `.env` file.
//   - error: An error if reading or parsing the file fails.
func loadEnvFile(filePath string) (map[string]string, error) {
	envMap, err := godotenv.Read(filePath)
	if err != nil {
		return nil, err
	}
	return envMap, nil
}
