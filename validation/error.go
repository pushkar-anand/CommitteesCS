package validation

import (
	"github.com/go-playground/validator/v10"
)

// HandleValidationError accumulates all the struct fields that failed validation.
//
// It returns a slice of all the invalid fields.
// In case the validation fails, returns error.
func HandleValidationError(err error) ([]string, error) {
	fields := make([]string, 0)

	if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
		return nil, invalidErr
	}

	for _, validationErr := range err.(validator.ValidationErrors) {
		fields = append(fields, validationErr.Field())
	}

	return fields, nil
}
