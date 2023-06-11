package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jnhu76/dwz/pkg/app"
	"github.com/jnhu76/dwz/pkg/e"
	"github.com/jnhu76/dwz/pkg/util"
	"github.com/jnhu76/dwz/service/auth_service"
)

type auth struct {
	Username string `validate:"required,max=50"`
	Password string `validate:"required,max=50"`
}

// @Summary Get Auth
// @Produce  json
// @Param username formData string true "userName"
// @Param password formData string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	validate := validator.New()

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	err := validate.Struct(&a)

	if err != nil {
		app.MakeErrors(err.(validator.ValidationErrors))
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth_Service{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	fmt.Println(err)
	fmt.Println(token)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}

// @Summary Hello world
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /hello [get]
func GetHello(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"Hello": "world",
	})
}
