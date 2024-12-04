package plugins

import (
	"fmt"
	"strings"
)

// EnumValidationPlugin validates that the value of a specific environment variable
// key is within a predefined set of allowed options. The validation can be
// configured to be case-sensitive or case-insensitive.
type EnumValidationPlugin struct {
	Key           string   // The key of the environment variable to validate.
	AllowedValues []string // A list of permissible values for the key.
	CaseSensitive bool     // If true, validation is case-sensitive; otherwise, it is case-insensitive.
}

// Validate verifies if the value for the specified key is within the allowed set of values.
// It checks against the plugin's `Key` and `AllowedValues`.
//
// Parameters:
//   - key: The key of the environment variable being validated.
//   - value: The value of the environment variable to validate.
//
// Returns:
//   - bool: Indicates whether this plugin handled the validation.
//   - error: An error if the value is invalid or nil if it passes validation.
func (p *EnumValidationPlugin) Validate(key, value string) (bool, error) {
	if key != p.Key {
		return false, nil // Plugin does not handle this key.
	}

	for _, allowed := range p.AllowedValues {
		if p.CaseSensitive {
			if value == allowed {
				return true, nil
			}
		} else {
			if strings.EqualFold(value, allowed) {
				return true, nil
			}
		}
	}

	return true, fmt.Errorf("value for key %q must be one of %v", key, p.AllowedValues)
}

// Name provides the name of the plugin.
//
// Returns:
//   - string: The name of the plugin.
func (p *EnumValidationPlugin) Name() string {
	return "EnumValidationPlugin"
}
