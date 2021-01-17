package validation

import (
	"github.com/go-playground/validator/v10"
)

type validatorFn func(fl validator.FieldLevel) bool

type stringModifierFn func(str string) string

// AddCustomValidator adds a custom validator tag
func AddCustomValidator(tagName string, f validatorFn) {
	_ = validate.RegisterValidation(tagName, validator.Func(f))
}

// AddStructLevelValidation adds a struct level validation
func AddStructLevelValidation(fn validator.StructLevelFunc, structType interface{}) {
	validate.RegisterStructValidation(fn, structType)
}

// AddStringModifier is used to add tags that modifies the field value.
// Works only if the field type is string
func AddStringModifier(tagName string, fn stringModifierFn) {
	AddCustomValidator(tagName, func(fl validator.FieldLevel) bool {
		if fl.Field().Type().String() == "string" {
			str := fn(fl.Field().String())
			fl.Field().SetString(str)
		}
		return true
	})
}
