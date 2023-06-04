package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jnhu76/dwz/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		MakeErrors(err.(validator.ValidationErrors))
		return http.StatusInternalServerError, e.ERROR
	}

	return http.StatusOK, e.SUCCESS
}
