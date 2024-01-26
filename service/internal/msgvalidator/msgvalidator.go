package msgvalidator

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type MessageValidator struct {
	validator *validator.Validate
}

func New(validator *validator.Validate) *MessageValidator {
	return &MessageValidator{
		validator: validator,
	}
}

func (v MessageValidator) Validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}
	if errs := v.validator.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var e ErrorResponse
			e.FailedField = err.Field()
			e.Tag = err.Tag()
			e.Value = err.Value()
			e.Error = true
			validationErrors = append(validationErrors, e)
		}
	}
	return validationErrors
}
