package plugins

import (
	"fmt"
	"strings"
)

// BooleanValidationPlugin validates that the value of a specific environment variable
// key conforms to one of the accepted boolean representations. Optionally, it can
// standardize the value to a canonical form ("true" or "false").
type BooleanValidationPlugin struct {
	Key            string   // The key of the environment variable to validate.
	AcceptedValues []string // A list of accepted boolean representations (e.g., "true", "false", "1", "0").
	Standardize    bool     // If true, standardizes the value to "true" or "false".
}

// Validate verifies if the value for the specified key is a valid boolean representation.
// It checks against the plugin's `Key` and `AcceptedValues`.
//
// Parameters:
//   - key: The key of the environment variable being validated.
//   - value: The value of the environment variable to validate.
//
// Returns:
//   - bool: Indicates whether this plugin handled the validation.
//   - error: An error if the value is invalid or nil if it passes validation.
func (p *BooleanValidationPlugin) Validate(key, value string) (bool, error) {
	if key != p.Key {
		return false, nil // Plugin does not handle this key.
	}

	normalizedValue := strings.ToLower(strings.TrimSpace(value))
	valid := false
	for _, accepted := range p.AcceptedValues {
		if normalizedValue == strings.ToLower(accepted) {
			valid = true
			break
		}
	}

	if !valid {
		return true, fmt.Errorf("value for key %q must be a boolean (accepted values: %v)", key, p.AcceptedValues)
	}

	if p.Standardize {
		// Example of standardization logic can be implemented here if needed.
	}

	return true, nil
}

// Name provides the name of the plugin.
//
// Returns:
//   - string: The name of the plugin.
func (p *BooleanValidationPlugin) Name() string {
	return "BooleanValidationPlugin"
}
