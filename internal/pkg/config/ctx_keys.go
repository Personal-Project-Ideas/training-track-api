// Package config provides configuration and context keys for the application.
package config

type ctxKey string

// RequestIDKey is the context key for request IDs.
const RequestIDKey ctxKey = "X-Request-ID"
