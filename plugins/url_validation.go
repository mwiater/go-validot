package plugins

import (
	"fmt"
	"net/url"
	"strings"
)

// URLValidationPlugin validates that the value of a specific environment variable
// key is a well-formed URL. It can optionally enforce that the URL's scheme is in a
// predefined set of allowed schemes (e.g., "https").
type URLValidationPlugin struct {
	Key            string   // The key of the environment variable to validate.
	AllowedSchemes []string // A list of allowed URL schemes, e.g., "http", "https". Optional.
}

// Validate checks if the value associated with the given key is a valid URL
// and optionally enforces that its scheme is one of the allowed schemes.
//
// Parameters:
//   - key: The key of the environment variable being validated.
//   - value: The value of the environment variable to validate.
//
// Returns:
//   - bool: Indicates whether this plugin handled the validation.
//   - error: An error if the value is invalid or nil if it passes validation.
func (p *URLValidationPlugin) Validate(key, value string) (bool, error) {
	if key != p.Key {
		return false, nil // Plugin does not handle this key.
	}

	parsedURL, err := url.Parse(value)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return true, fmt.Errorf("value for key %q must be a valid URL", key)
	}

	if len(p.AllowedSchemes) > 0 {
		validScheme := false
		for _, scheme := range p.AllowedSchemes {
			if strings.EqualFold(parsedURL.Scheme, scheme) { // Use EqualFold here
				validScheme = true
				break
			}
		}
		if !validScheme {
			return true, fmt.Errorf("URL scheme for key %q must be one of %v", key, p.AllowedSchemes)
		}
	}

	return true, nil
}

// Name provides the name of the plugin.
//
// Returns:
//   - string: The name of the plugin.
func (p *URLValidationPlugin) Name() string {
	return "URLValidationPlugin"
}
