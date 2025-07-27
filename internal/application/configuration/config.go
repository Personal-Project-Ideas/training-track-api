// Package configuration  loads the configuration for the application, the env vars
package configuration

// ctxKey is a type for context keys to avoid collisions with other packages.
type ctxKey string

const (
	// RequestIDKey is the key used to store the request ID in the context.
	RequestIDKey ctxKey = "X-Request-ID"
)

// Config holds the application configuration.
type Config struct{}

// LoadConfig initializes and returns a new Config instance.
func LoadConfig() *Config {

	return &Config{}

}
