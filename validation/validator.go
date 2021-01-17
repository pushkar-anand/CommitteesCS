package validation

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var idRegex = regexp.MustCompile(`^[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}$`)

var validate *validator.Validate

func init() {
	validate = validator.New()

	// register function to get tag name from json tags.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "" {
			name = strings.SplitN(fld.Tag.Get("schema"), ",", 2)[0]
		}
		if name == "-" {
			return ""
		}
		return name
	})

	AddStringModifier("to_lower", strings.ToLower)
	AddStringModifier("title_upper", func(str string) string {
		return strings.Title(strings.ToLower(str))
	})
	AddStringModifier("trim", strings.TrimSpace)
}

// DoValidation validates the input as per the set validation tags.
//
// Returns:
//
// bool - true if input is valid
//
// []string - list of all fields that failed validation
//
// error - error if validation fails
//
func DoValidation(data interface{}) (bool, []string, error) {
	err := validate.Struct(data)

	if err != nil {
		errFields, invalidErr := HandleValidationError(err)
		return false, errFields, invalidErr
	}

	return true, nil, nil
}

// IsValidUUID checks if the given string is a valid UUID
func IsValidUUID(str string) bool {
	return idRegex.MatchString(str)
}
