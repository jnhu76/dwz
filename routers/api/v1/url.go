package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnhu76/dwz/pkg/app"
	"github.com/jnhu76/dwz/pkg/e"
)

// @Summary Hello jwt
// @Produce json
// @Security Bearer
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/jwt [get]
func GetJwt(c *gin.Context) {
	appG := app.Gin{C: c}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"Hello": "jwt",
	})
}
