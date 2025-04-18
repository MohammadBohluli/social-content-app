package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/MohammadBohluli/social-content-app/types"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func New() *validator.Validate {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}

func GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", fe.Field())
	case "max":
		return fmt.Sprintf("%s must be less than %s characters.", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be greater than %s characters.", fe.Field(), fe.Param())
	default:
		return fmt.Sprintf("%s not valid.", fe.Field())
	}
}

func ParseValidationErrors(err error) types.FieldErrors {
	var errors types.FieldErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			errors = append(errors, types.FieldError{
				Field: fieldErr.Field(),
				Error: GetErrorMessage(fieldErr),
			})
		}
	}
	return errors
}
