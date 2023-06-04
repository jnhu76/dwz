package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/jnhu76/dwz/pkg/logging"
)

func MakeErrors(errors validator.ValidationErrors) {
	for _, err := range errors {
		logging.Info(err.Field(), err.Error())
	}
	return
}
