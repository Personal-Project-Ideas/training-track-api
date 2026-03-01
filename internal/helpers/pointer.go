// Package helpers contains  functions for working with pointers in Go.
package helpers

// ToPointer returns a pointer to the given value of any type.
func ToPointer[T any](v T) *T {
	return &v
}
