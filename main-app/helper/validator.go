package helper

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Initialize a validator instance
var validate = validator.New()

// ValidationError represents a structured validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateStruct validates the given struct and returns a slice of errors
func ValidateStruct(obj interface{}) []ValidationError {
	var errors []ValidationError

	// Validate the struct
	err := validate.Struct(obj)
	if err != nil {
		// Check if it's a validation error
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errors = append(errors, ValidationError{
					Field:   e.Field(),
					Message: fmt.Sprintf("failed on '%s' validation", e.Tag()),
				})
			}
		}
	}

	return errors
}
