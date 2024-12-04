package validot

import (
	"github.com/mwiater/go-validot/plugins"
	"github.com/sirupsen/logrus"
)

// Config represents the configuration settings for a Validator.
// This structure defines the behavior of the validation process,
// including logging, verbosity, and custom plugins.
type Config struct {
	RequireQuotes bool                       // If true, enforces that all values in the `.env` file must be quoted.
	Verbose       bool                       // If true, enables detailed logging for the validation process.
	Logger        *logrus.Logger             // A custom logger instance for logging messages; if nil, a default logger will be used.
	Plugins       []plugins.ValidationPlugin // A list of user-defined validation plugins to extend the validation capabilities.
}
