package plugins

// ValidationPlugin defines the interface that all validation plugins must implement.
// Validation plugins are used to enforce specific validation rules for key-value pairs
// in environment variable files.
type ValidationPlugin interface {
	// Validate checks whether the provided key-value pair conforms to the plugin's validation rules.
	//
	// Parameters:
	//   - key: The environment variable key being validated.
	//   - value: The value of the environment variable to validate.
	//
	// Returns:
	//   - bool: Indicates whether the plugin handled the validation for the given key.
	//   - error: An error if the value does not satisfy the validation rules, or nil if validation passes.
	Validate(key, value string) (bool, error)

	// Name returns the name of the plugin for identification purposes.
	//
	// Returns:
	//   - string: The name of the plugin.
	Name() string
}
