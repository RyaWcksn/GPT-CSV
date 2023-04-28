package validator

import (
	"fmt"
	customerror "github.com/RyaWcksn/nann-e/pkgs/error"
	"github.com/go-playground/validator"
)

func msgForTag(error validator.FieldError) string {
	switch error.Tag() {
	case "required":
		return fmt.Sprintf("The %s field is required", error.Field())
	case "min":
		return fmt.Sprintf("The %s field can't be less than %v", error.Field(), error.Param())
	case "max":
		return fmt.Sprintf("The %s field can't be more than %v", error.Field(), error.Param())
	case "uppercase":
		return fmt.Sprintf("The %s field should all be uppercase", error.Field())
	case "email":
		return fmt.Sprintf("The %s field is invalid", error.Field())
	case "eq=0|eq=1":
		return fmt.Sprintf("The %s field has invalid input", error.Field())
	default:
		return ""
	}
}

func Validate(request interface{}) error {
	var validateErrors []customerror.ErrorValidatorDetails
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errItems := customerror.ErrorValidatorDetails{
				Field:   err.Field(),
				Message: msgForTag(err),
			}
			validateErrors = append(validateErrors, errItems)
		}
		return customerror.GetErrorValidation(customerror.BadRequest, validateErrors)
	}
	return nil
}
