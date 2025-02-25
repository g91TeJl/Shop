package model

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func requiredIf(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}
		return errors.New("error")
	}
}
