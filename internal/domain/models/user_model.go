// Package models provides the domain models for the application.
package models

import (
	"github.com/google/uuid"
)

// User represents a user in the system.
type User struct {
	ID        *uuid.UUID
	FullName  string
	Age       int
	Email     string
	Password  string
	CreatedAt *string
	UpdatedAt *string
}
