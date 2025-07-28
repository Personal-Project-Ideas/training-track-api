package dtos

import "time"

// UserInputDTO represents the input data transfer object for creating or updating a user.
type UserInputDTO struct {
	FullName string `json:"full_name" validate:"required,min=2,max=255"`
	Age      int    `json:"age" validate:"gte=0,lte=150"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// UserUpdateInputDTO represents the input data transfer object for updating a user.
type UserUpdateInputDTO struct {
	FullName *string `json:"full_name,omitempty" validate:"omitempty,min=2,max=255"`
	Age      *int    `json:"age,omitempty" validate:"omitempty,gte=0,lte=150"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Password *string `json:"password,omitempty" validate:"omitempty,min=8"`
}

// UserOutputDTO represents the output data transfer object for a user.
type UserOutputDTO struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
