# go-validot

**go-validot** is a robust Go package designed to validate `.env` files, ensuring that your application's environment configurations are correctly set and adhere to specified standards. With support for custom plugins, comprehensive configuration options, and a suite of examples, `go-validot` offers a flexible solution for managing environment variables effectively.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Configuration](#configuration)
- [Plugins](#plugins)
- [Best Practices](#best-practices)
- [Running the Tests](#running-the-tests)

## Features

- **Validation of Required Keys:** Ensure that essential environment variables are present.
- **Customizable Validation Rules:** Define and enforce specific rules for environment variable values.
- **Plugin Architecture:** Extend functionality with custom validation plugins.
- **Verbose Logging:** Gain insights into the validation process with detailed logs.
- **Graceful Handling of Invalid Lines:** Skip malformed lines without halting the validation process.
- **Support for Multiple Data Types:** Validate URLs, enums, IP addresses, booleans, and more.

## Installation

To install `go-validot`, use the `go get` command:

```bash
go get github.com/mwiater/go-validot
```

Ensure that you have [Go](https://golang.org/dl/) installed and properly configured on your system.

## Usage

Using `go-validot` is straightforward. Below is a basic example of how to integrate it into your Go project:

1. **Import the Package:**

   ```go
   import (
       "github.com/mwiater/go-validot"
       "github.com/sirupsen/logrus"
   )
   ```

2. **Configure and Initialize the Validator:**

   ```go
   // Initialize logger
   logger := logrus.New()
   logger.SetOutput(os.Stdout)
   logger.SetLevel(logrus.InfoLevel)

   // Define required keys
   requiredKeys := map[string]bool{
       "API_URL":          true,
       "SERVICE_ENDPOINT": true,
       "DB_HOST":          true,
       "ENVIRONMENT":      true,
       "ENABLE_DEBUG":     true,
       "TRUSTED_PROXY_IP": true,
   }

   // Create validator with desired configuration
   validator := validot.NewValidator(validot.Config{
       RequireQuotes: true,
       Verbose:       true,
       Logger:        logger,
       Plugins:       []validot.ValidationPlugin{&validot.CacheSizeValidationPlugin{}},
   }, requiredKeys)

   // Validate the .env file
   err := validator.ValidateDotEnv(".env")
   if err != nil {
       logger.Fatalf("Validation failed: %v", err)
   }

   logger.Info(".env file is valid.")
   ```

3. **Run Your Application:**

   Ensure that your `.env` file is present in the expected location and run your application as usual. The validator will check the `.env` file against the defined rules and log the results accordingly.

## Examples

`go-validot` comes with a set of example projects that demonstrate various validation scenarios. Each example resides in the `examples/` directory and showcases how to implement specific validation rules.

### Available Examples

1. **Boolean Validation**

   - **Description:** Validates boolean environment variables, ensuring they accept only predefined boolean values.
   - **How to Run:**
     ```bash
     cd examples/boolean_validation
     go run .
     ```
   - **Expected Output:**
      ```
      INFO[2024-12-03T10:17:43-08:00] Validator Configuration:
      INFO[2024-12-03T10:17:43-08:00]   RequireQuotes: true
      INFO[2024-12-03T10:17:43-08:00]   Verbose: true
      INFO[2024-12-03T10:17:43-08:00]   Number of Plugins: 4
      INFO[2024-12-03T10:17:43-08:00] End of Configuration
      INFO[2024-12-03T10:17:43-08:00] Starting validation for file: .env
      INFO[2024-12-03T10:17:43-08:00] Processing key: REDIS_PORT
      INFO[2024-12-03T10:17:43-08:00]   REDIS_PORT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: API_KEY
      INFO[2024-12-03T10:17:43-08:00]   API_KEY is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: ENABLE_FEATURE_Y
      INFO[2024-12-03T10:17:43-08:00]   ENABLE_FEATURE_Y is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: REDIS_HOST
      INFO[2024-12-03T10:17:43-08:00]   REDIS_HOST is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: ENABLE_DEBUG
      INFO[2024-12-03T10:17:43-08:00]   ENABLE_DEBUG is a required variable.
      INFO[2024-12-03T10:17:43-08:00]   [Validated by: BooleanValidationPlugin]
      INFO[2024-12-03T10:17:43-08:00] Processing key: USE_SSL
      INFO[2024-12-03T10:17:43-08:00]   USE_SSL is a required variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: DB_PORT
      INFO[2024-12-03T10:17:43-08:00]   DB_PORT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: ENVIRONMENT
      INFO[2024-12-03T10:17:43-08:00]   ENVIRONMENT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00]   [Validated by: EnumValidationPlugin]
      INFO[2024-12-03T10:17:43-08:00] Processing key: ENABLE_FEATURE_X
      INFO[2024-12-03T10:17:43-08:00]   ENABLE_FEATURE_X is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: TRUSTED_PROXY_IP
      INFO[2024-12-03T10:17:43-08:00]   TRUSTED_PROXY_IP is an optional variable.
      INFO[2024-12-03T10:17:43-08:00]   [Validated by: IPAddressValidationPlugin]
      INFO[2024-12-03T10:17:43-08:00] Processing key: API_SECRET
      INFO[2024-12-03T10:17:43-08:00]   API_SECRET is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: SERVICE_TIMEOUT
      INFO[2024-12-03T10:17:43-08:00]   SERVICE_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: SERVICE_ENDPOINT
      INFO[2024-12-03T10:17:43-08:00]   SERVICE_ENDPOINT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: SERVICE_VERSION
      INFO[2024-12-03T10:17:43-08:00]   SERVICE_VERSION is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: DB_PASSWORD
      INFO[2024-12-03T10:17:43-08:00]   DB_PASSWORD is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: LOG_LEVEL
      INFO[2024-12-03T10:17:43-08:00]   LOG_LEVEL is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: CACHE_SIZE
      INFO[2024-12-03T10:17:43-08:00]   CACHE_SIZE is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: DB_HOST
      INFO[2024-12-03T10:17:43-08:00]   DB_HOST is a required variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: UPLOAD_LIMIT
      INFO[2024-12-03T10:17:43-08:00]   UPLOAD_LIMIT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: DB_USER
      INFO[2024-12-03T10:17:43-08:00]   DB_USER is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: API_URL
      INFO[2024-12-03T10:17:43-08:00]   API_URL is an optional variable.
      INFO[2024-12-03T10:17:43-08:00]   [Validated by: URLValidationPlugin]
      INFO[2024-12-03T10:17:43-08:00] Processing key: API_TIMEOUT
      INFO[2024-12-03T10:17:43-08:00]   API_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: LOG_FORMAT
      INFO[2024-12-03T10:17:43-08:00]   LOG_FORMAT is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: DB_NAME
      INFO[2024-12-03T10:17:43-08:00]   DB_NAME is an optional variable.
      INFO[2024-12-03T10:17:43-08:00] Processing key: FEATURE_FLAG_NEW_UI
      INFO[2024-12-03T10:17:43-08:00]   FEATURE_FLAG_NEW_UI is a required variable.
      INFO[2024-12-03T10:17:43-08:00] .env file is valid.
      ```

2. **Default Usage**

   - **Description:** Demonstrates the basic usage of `go-validot` without any custom plugins.
   - **How to Run:**
     ```bash
     cd examples/default_usage
     go run .
     ```
   - **Expected Output:**
     ```
     INFO[2024-12-03T10:20:10-08:00] Starting validation for file: .env
     INFO[2024-12-03T10:20:10-08:00] .env file is valid.
     ```

3. **Enum Validation**

   - **Description:** Ensures that specific environment variables match one of the allowed enumerated values.
   - **How to Run:**
     ```bash
     cd examples/enum_validation
     go run .
     ```
   - **Expected Output:**
      ```
      INFO[2024-12-03T10:20:38-08:00] Validator Configuration:
      INFO[2024-12-03T10:20:38-08:00]   RequireQuotes: true
      INFO[2024-12-03T10:20:38-08:00]   Verbose: true
      INFO[2024-12-03T10:20:38-08:00]   Number of Plugins: 4
      INFO[2024-12-03T10:20:38-08:00] End of Configuration
      INFO[2024-12-03T10:20:38-08:00] Starting validation for file: .env
      INFO[2024-12-03T10:20:38-08:00] Processing key: SERVICE_TIMEOUT
      INFO[2024-12-03T10:20:38-08:00]   SERVICE_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: DB_PORT
      INFO[2024-12-03T10:20:38-08:00]   DB_PORT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: USE_SSL
      INFO[2024-12-03T10:20:38-08:00]   USE_SSL is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: API_SECRET
      INFO[2024-12-03T10:20:38-08:00]   API_SECRET is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: SERVICE_VERSION
      INFO[2024-12-03T10:20:38-08:00]   SERVICE_VERSION is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: TRUSTED_PROXY_IP
      INFO[2024-12-03T10:20:38-08:00]   TRUSTED_PROXY_IP is an optional variable.
      INFO[2024-12-03T10:20:38-08:00]   [Validated by: IPAddressValidationPlugin]
      INFO[2024-12-03T10:20:38-08:00] Processing key: REDIS_PORT
      INFO[2024-12-03T10:20:38-08:00]   REDIS_PORT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: DB_USER
      INFO[2024-12-03T10:20:38-08:00]   DB_USER is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: REDIS_HOST
      INFO[2024-12-03T10:20:38-08:00]   REDIS_HOST is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: ENABLE_FEATURE_Y
      INFO[2024-12-03T10:20:38-08:00]   ENABLE_FEATURE_Y is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: ENVIRONMENT
      INFO[2024-12-03T10:20:38-08:00]   ENVIRONMENT is a required variable.
      INFO[2024-12-03T10:20:38-08:00]   [Validated by: EnumValidationPlugin]
      INFO[2024-12-03T10:20:38-08:00] Processing key: LOG_FORMAT
      INFO[2024-12-03T10:20:38-08:00]   LOG_FORMAT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: DB_HOST
      INFO[2024-12-03T10:20:38-08:00]   DB_HOST is a required variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: SERVICE_ENDPOINT
      INFO[2024-12-03T10:20:38-08:00]   SERVICE_ENDPOINT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: CACHE_SIZE
      INFO[2024-12-03T10:20:38-08:00]   CACHE_SIZE is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: API_KEY
      INFO[2024-12-03T10:20:38-08:00]   API_KEY is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: API_URL
      INFO[2024-12-03T10:20:38-08:00]   API_URL is an optional variable.
      INFO[2024-12-03T10:20:38-08:00]   [Validated by: URLValidationPlugin]
      INFO[2024-12-03T10:20:38-08:00] Processing key: ENABLE_FEATURE_X
      INFO[2024-12-03T10:20:38-08:00]   ENABLE_FEATURE_X is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: LOG_LEVEL
      INFO[2024-12-03T10:20:38-08:00]   LOG_LEVEL is a required variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: UPLOAD_LIMIT
      INFO[2024-12-03T10:20:38-08:00]   UPLOAD_LIMIT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: API_TIMEOUT
      INFO[2024-12-03T10:20:38-08:00]   API_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: DB_PASSWORD
      INFO[2024-12-03T10:20:38-08:00]   DB_PASSWORD is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: DB_NAME
      INFO[2024-12-03T10:20:38-08:00]   DB_NAME is an optional variable.
      INFO[2024-12-03T10:20:38-08:00] Processing key: ENABLE_DEBUG
      INFO[2024-12-03T10:20:38-08:00]   ENABLE_DEBUG is an optional variable.
      INFO[2024-12-03T10:20:38-08:00]   [Validated by: BooleanValidationPlugin]
      INFO[2024-12-03T10:20:38-08:00] .env file is valid.
      ```

4. **IP Address Validation**

   - **Description:** Validates that IP address environment variables are within private IP ranges.
   - **How to Run:**
     ```bash
     cd examples/ip_address_validation
     go run .
     ```
   - **Expected Output:**
      ```
      INFO[2024-12-03T10:22:56-08:00] Validator Configuration:
      INFO[2024-12-03T10:22:56-08:00]   RequireQuotes: true
      INFO[2024-12-03T10:22:56-08:00]   Verbose: true
      INFO[2024-12-03T10:22:56-08:00]   Number of Plugins: 4
      INFO[2024-12-03T10:22:56-08:00] End of Configuration
      INFO[2024-12-03T10:22:56-08:00] Starting validation for file: .env
      INFO[2024-12-03T10:22:56-08:00] Processing key: DATABASE_IP
      INFO[2024-12-03T10:22:56-08:00]   DATABASE_IP is a required variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: DB_USER
      INFO[2024-12-03T10:22:56-08:00]   DB_USER is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: API_KEY
      INFO[2024-12-03T10:22:56-08:00]   API_KEY is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: SERVICE_VERSION
      INFO[2024-12-03T10:22:56-08:00]   SERVICE_VERSION is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: DB_PASSWORD
      INFO[2024-12-03T10:22:56-08:00]   DB_PASSWORD is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: CACHE_SIZE
      INFO[2024-12-03T10:22:56-08:00]   CACHE_SIZE is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: API_SECRET
      INFO[2024-12-03T10:22:56-08:00]   API_SECRET is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: DB_HOST
      INFO[2024-12-03T10:22:56-08:00]   DB_HOST is a required variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: API_URL
      INFO[2024-12-03T10:22:56-08:00]   API_URL is an optional variable.
      INFO[2024-12-03T10:22:56-08:00]   [Validated by: URLValidationPlugin]
      INFO[2024-12-03T10:22:56-08:00] Processing key: TRUSTED_PROXY_IP
      INFO[2024-12-03T10:22:56-08:00]   TRUSTED_PROXY_IP is a required variable.
      INFO[2024-12-03T10:22:56-08:00]   [Validated by: IPAddressValidationPlugin]
      INFO[2024-12-03T10:22:56-08:00] Processing key: DB_PORT
      INFO[2024-12-03T10:22:56-08:00]   DB_PORT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: LOG_FORMAT
      INFO[2024-12-03T10:22:56-08:00]   LOG_FORMAT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: UPLOAD_LIMIT
      INFO[2024-12-03T10:22:56-08:00]   UPLOAD_LIMIT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: ENABLE_DEBUG
      INFO[2024-12-03T10:22:56-08:00]   ENABLE_DEBUG is an optional variable.
      INFO[2024-12-03T10:22:56-08:00]   [Validated by: BooleanValidationPlugin]
      INFO[2024-12-03T10:22:56-08:00] Processing key: USE_SSL
      INFO[2024-12-03T10:22:56-08:00]   USE_SSL is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: ENABLE_FEATURE_X
      INFO[2024-12-03T10:22:56-08:00]   ENABLE_FEATURE_X is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: ENABLE_FEATURE_Y
      INFO[2024-12-03T10:22:56-08:00]   ENABLE_FEATURE_Y is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: REDIS_PORT
      INFO[2024-12-03T10:22:56-08:00]   REDIS_PORT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: LOG_LEVEL
      INFO[2024-12-03T10:22:56-08:00]   LOG_LEVEL is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: REDIS_HOST
      INFO[2024-12-03T10:22:56-08:00]   REDIS_HOST is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: SERVICE_TIMEOUT
      INFO[2024-12-03T10:22:56-08:00]   SERVICE_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: DB_NAME
      INFO[2024-12-03T10:22:56-08:00]   DB_NAME is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: API_TIMEOUT
      INFO[2024-12-03T10:22:56-08:00]   API_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: SERVICE_ENDPOINT
      INFO[2024-12-03T10:22:56-08:00]   SERVICE_ENDPOINT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: FEATURE_FLAG_NEW_UI
      INFO[2024-12-03T10:22:56-08:00]   FEATURE_FLAG_NEW_UI is an optional variable.
      INFO[2024-12-03T10:22:56-08:00] Processing key: ENVIRONMENT
      INFO[2024-12-03T10:22:56-08:00]   ENVIRONMENT is an optional variable.
      INFO[2024-12-03T10:22:56-08:00]   [Validated by: EnumValidationPlugin]
      INFO[2024-12-03T10:22:56-08:00] .env file is valid.
      ```

5. **URL Validation**

   - **Description:** Ensures that URL environment variables use secure schemes like `https`.
   - **How to Run:**
     ```bash
     cd examples/url_validation
     go run .
     ```
   - **Expected Output:**
      ```
      INFO[2024-12-03T10:23:26-08:00] Validator Configuration:
      INFO[2024-12-03T10:23:26-08:00]   RequireQuotes: true
      INFO[2024-12-03T10:23:26-08:00]   Verbose: true
      INFO[2024-12-03T10:23:26-08:00]   Number of Plugins: 4
      INFO[2024-12-03T10:23:26-08:00] End of Configuration
      INFO[2024-12-03T10:23:26-08:00] Starting validation for file: .env
      INFO[2024-12-03T10:23:26-08:00] Processing key: ENABLE_FEATURE_X
      INFO[2024-12-03T10:23:26-08:00]   ENABLE_FEATURE_X is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: LOG_LEVEL
      INFO[2024-12-03T10:23:26-08:00]   LOG_LEVEL is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: USE_SSL
      INFO[2024-12-03T10:23:26-08:00]   USE_SSL is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: DB_PASSWORD
      INFO[2024-12-03T10:23:26-08:00]   DB_PASSWORD is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: TRUSTED_PROXY_IP
      INFO[2024-12-03T10:23:26-08:00]   TRUSTED_PROXY_IP is an optional variable.
      INFO[2024-12-03T10:23:26-08:00]   [Validated by: IPAddressValidationPlugin]
      INFO[2024-12-03T10:23:26-08:00] Processing key: REDIS_PORT
      INFO[2024-12-03T10:23:26-08:00]   REDIS_PORT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: DB_NAME
      INFO[2024-12-03T10:23:26-08:00]   DB_NAME is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: REDIS_HOST
      INFO[2024-12-03T10:23:26-08:00]   REDIS_HOST is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: DB_PORT
      INFO[2024-12-03T10:23:26-08:00]   DB_PORT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: LOG_FORMAT
      INFO[2024-12-03T10:23:26-08:00]   LOG_FORMAT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: SERVICE_ENDPOINT
      INFO[2024-12-03T10:23:26-08:00]   SERVICE_ENDPOINT is a required variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: SERVICE_VERSION
      INFO[2024-12-03T10:23:26-08:00]   SERVICE_VERSION is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: API_URL
      INFO[2024-12-03T10:23:26-08:00]   API_URL is a required variable.
      INFO[2024-12-03T10:23:26-08:00]   [Validated by: URLValidationPlugin]
      INFO[2024-12-03T10:23:26-08:00] Processing key: DB_USER
      INFO[2024-12-03T10:23:26-08:00]   DB_USER is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: ENABLE_FEATURE_Y
      INFO[2024-12-03T10:23:26-08:00]   ENABLE_FEATURE_Y is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: SERVICE_TIMEOUT
      INFO[2024-12-03T10:23:26-08:00]   SERVICE_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: DB_HOST
      INFO[2024-12-03T10:23:26-08:00]   DB_HOST is a required variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: ENABLE_DEBUG
      INFO[2024-12-03T10:23:26-08:00]   ENABLE_DEBUG is an optional variable.
      INFO[2024-12-03T10:23:26-08:00]   [Validated by: BooleanValidationPlugin]
      INFO[2024-12-03T10:23:26-08:00] Processing key: API_SECRET
      INFO[2024-12-03T10:23:26-08:00]   API_SECRET is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: ENVIRONMENT
      INFO[2024-12-03T10:23:26-08:00]   ENVIRONMENT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00]   [Validated by: EnumValidationPlugin]
      INFO[2024-12-03T10:23:26-08:00] Processing key: API_TIMEOUT
      INFO[2024-12-03T10:23:26-08:00]   API_TIMEOUT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: UPLOAD_LIMIT
      INFO[2024-12-03T10:23:26-08:00]   UPLOAD_LIMIT is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: API_KEY
      INFO[2024-12-03T10:23:26-08:00]   API_KEY is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] Processing key: CACHE_SIZE
      INFO[2024-12-03T10:23:26-08:00]   CACHE_SIZE is an optional variable.
      INFO[2024-12-03T10:23:26-08:00] .env file is valid.
      ```

## Configuration

`go-validot` offers a flexible configuration system to tailor the validation process to your project's needs. Below are the primary configuration options:

- **RequireQuotes (`bool`):**  
  Enforces that all environment variable values must be enclosed in quotes.  
  *Default:* `false`

- **Verbose (`bool`):**  
  Enables detailed logging of the validation process, providing insights into each validation step.  
  *Default:* `false`

- **Logger (`*logrus.Logger`):**  
  Allows you to specify a custom logger instance. By default, `go-validot` uses `logrus`.  
  *Default:* Initialized with `logrus.New()`

- **Plugins (`[]ValidationPlugin`):**  
  A slice of custom validation plugins to extend `go-validot`'s functionality. Plugins can enforce additional rules beyond the core validations.

### Example Configuration

```go
validator := validot.NewValidator(validot.Config{
    RequireQuotes: true,
    Verbose:       true,
    Logger:        logger,
    Plugins:       []validot.ValidationPlugin{&validot.CacheSizeValidationPlugin{}},
}, requiredKeys)
```

## Plugins

`go-validot` supports a plugin architecture, allowing developers to create and integrate custom validation rules seamlessly. Below are descriptions of the available plugins:

### 1. **BooleanValidationPlugin**

- **Description:**
  
  Validates that specified environment variables are boolean values. Ensures that variables like `ENABLE_DEBUG` accept only predefined boolean representations (e.g., `true`, `false`, `1`, `0`, `yes`, `no`).

- **Usage:**
  
  Integrate the plugin into the validator's configuration to enforce boolean constraints on relevant keys.

- **Example Behavior:**
  
  - **Valid:**
    - `ENABLE_DEBUG="true"`
    - `ENABLE_FEATURE_X="1"`
  
  - **Invalid:**
    - `ENABLE_DEBUG="maybe"`
    - `ENABLE_FEATURE_X="enabled"`

### 2. **EnumValidationPlugin**

- **Description:**
  
  Ensures that specific environment variables match one of the allowed enumerated values. Useful for variables like `ENVIRONMENT` which should be `DEVELOPMENT`, `STAGING`, or `PRODUCTION`.

- **Usage:**
  
  Configure the plugin with the allowed values for each key you wish to validate.

- **Example Behavior:**
  
  - **Valid:**
    - `ENVIRONMENT="PRODUCTION"`
    - `LOG_LEVEL="INFO"`
  
  - **Invalid:**
    - `ENVIRONMENT="TESTING"`
    - `LOG_LEVEL="VERBOSE"`

### 3. **IPAddressValidationPlugin**

- **Description:**
  
  Checks that IP address environment variables are within private IP ranges, enhancing security by preventing the use of public IPs where inappropriate.

- **Usage:**
  
  Integrate the plugin and specify which keys should be validated as private IP addresses.

- **Example Behavior:**
  
  - **Valid:**
    - `TRUSTED_PROXY_IP="192.168.1.100"`
  
  - **Invalid:**
    - `TRUSTED_PROXY_IP="8.8.8.8"`

### 4. **URLValidationPlugin**

- **Description:**
  
  Ensures that URL environment variables use secure schemes such as `https`. This plugin can be configured to enforce specific schemes for different keys.

- **Usage:**
  
  Add the plugin to the validator and specify which keys should adhere to certain URL schemes.

- **Example Behavior:**
  
  - **Valid:**
    - `API_URL="https://api.example.com"`
  
  - **Invalid:**
    - `API_URL="http://api.example.com"` (Insecure scheme)

### Creating Custom Plugins

To create a custom plugin, implement the `ValidationPlugin` interface:

```go
type ValidationPlugin interface {
    Validate(key, value string) (handled bool, err error)
    Name() string
}
```

**Example:**

```go
type CustomPlugin struct{}

func (p *CustomPlugin) Validate(key, value string) (bool, error) {
    if key == "CUSTOM_KEY" {
        // Implement custom validation logic
        return true, nil
    }
    return false, nil
}

func (p *CustomPlugin) Name() string {
    return "CustomPlugin"
}
```

Integrate the custom plugin into the validator:

```go
validator := validot.NewValidator(validot.Config{
    RequireQuotes: true,
    Verbose:       true,
    Logger:        logger,
    Plugins:       []validot.ValidationPlugin{&CustomPlugin{}},
}, requiredKeys)
```

## Best Practices

To maximize the effectiveness of `go-validot`, consider the following best practices:

1. **Define Clear Validation Rules:**
   - Clearly outline which environment variables are required and what constraints they should adhere to. This clarity helps in setting up accurate validation rules.

2. **Use Verbose Logging During Development:**
   - Enable verbose logging to gain detailed insights into the validation process, making it easier to debug and refine your validation rules.

3. **Leverage Plugins for Custom Validations:**
   - Utilize the plugin architecture to enforce specific rules that are unique to your application's requirements. This extensibility ensures that `go-validot` remains adaptable to various scenarios.

4. **Regularly Update Validation Rules:**
   - As your application evolves, revisit and update your validation rules to accommodate new environment variables or changes in existing ones.

5. **Handle Sensitive Information Securely:**
   - Ensure that sensitive environment variables, such as API keys and passwords, are handled securely. Consider integrating additional plugins or mechanisms to validate and protect such data.

6. **Integrate with CI/CD Pipelines:**
   - Incorporate `go-validot` into your continuous integration and deployment pipelines to automatically validate environment configurations before deploying applications.

## Running the Tests

To ensure the integrity and reliability of `go-validot`, a comprehensive test suite is provided. Running the tests verifies that all validation rules and plugins function as expected.

### How to Run the Tests

Execute the following command in your terminal:

```bash
go test -v ./...
```

### Expected Output

Upon running the tests, you should see output indicating the status of each test case. Below is an example of a successful test run:

```
?       github.com/mwiater/go-validot/examples/boolean_validation       [no test files]
?       github.com/mwiater/go-validot/examples/default_usage    [no test files]
=== RUN   TestValidateDotEnv_ValidFile
--- PASS: TestValidateDotEnv_ValidFile (0.00s)
=== RUN   TestValidateDotEnv_MissingRequiredKeys
--- PASS: TestValidateDotEnv_MissingRequiredKeys (0.00s)
=== RUN   TestValidateDotEnv_InvalidURL
--- PASS: TestValidateDotEnv_InvalidURL (0.00s)
=== RUN   TestValidateDotEnv_InvalidEnum
--- PASS: TestValidateDotEnv_InvalidEnum (0.00s)
=== RUN   TestValidateDotEnv_InvalidBoolean
--- PASS: TestValidateDotEnv_InvalidBoolean (0.00s)
=== RUN   TestValidateDotEnv_InvalidIPAddress
--- PASS: TestValidateDotEnv_InvalidIPAddress (0.00s)
=== RUN   TestValidateDotEnv_DuplicateKeys
--- PASS: TestValidateDotEnv_DuplicateKeys (0.00s)
=== RUN   TestValidateDotEnv_EmptyFile
--- PASS: TestValidateDotEnv_EmptyFile (0.00s)
=== RUN   TestValidateDotEnv_InlineComments
--- PASS: TestValidateDotEnv_InlineComments (0.00s)
=== RUN   TestValidateDotEnv_NoQuotesWhenRequired
--- PASS: TestValidateDotEnv_NoQuotesWhenRequired (0.00s)
=== RUN   TestValidateDotEnv_CustomPlugin
--- PASS: TestValidateDotEnv_CustomPlugin (0.00s)
=== RUN   TestValidateDotEnv_InvalidKeyValue
--- PASS: TestValidateDotEnv_InvalidKeyValue (0.00s)
=== RUN   TestValidateDotEnv_KeyNotHandledByAnyPlugin
--- PASS: TestValidateDotEnv_KeyNotHandledByAnyPlugin (0.00s)
PASS
ok      github.com/mwiater/go-validot   (cached)
?       github.com/mwiater/go-validot/examples/enum_validation  [no test files]
?       github.com/mwiater/go-validot/examples/ip_address_validation    [no test files]
?       github.com/mwiater/go-validot/examples/url_validation   [no test files]
?       github.com/mwiater/go-validot/plugins   [no test files]
```

*Note:* The exact timing (`0.XXXs`) will vary based on system performance.

### Interpreting the Results

- **`PASS` Status:**  
  Indicates that the test case has successfully passed without any issues.

- **`FAIL` Status:**  
  Signifies that the test case encountered an error or did not meet the expected conditions. Review the error messages to identify and address the underlying issues.

By regularly running the test suite, you can ensure that `go-validot` remains reliable and continues to validate your `.env` files effectively as your project evolves.

---

For more information, contributions, or support, please refer to the [GitHub repository](https://github.com/mwiater/go-validot).