package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jnhu76/dwz/pkg/app"
	"github.com/jnhu76/dwz/pkg/e"
	"github.com/jnhu76/dwz/pkg/logging"
	"github.com/jnhu76/dwz/pkg/shorten"
	"github.com/jnhu76/dwz/service/url_service"
)

type Url struct {
	OriginUrl  string `json:"origin" binding:"required" valid:"http_url"`
	ShortenUrl string `json:"shorten" bingding:"required"`
}

type AcceptUrl struct {
	OriginUrl string
}

// // @Summary Hello jwt
// // @Produce json
// // @Security Bearer
// // @Success 200 {object} app.Response
// // @Failure 500 {object} app.Response
// // @Router /api/v1/jwt [get]
// func GetJwt(c *gin.Context) {
// 	appG := app.Gin{C: c}

// 	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
// 		"Hello": "jwt",
// 	})
// }

// @Summary Create URL
// @Accept json
// @Produce json
// @Param url body Url true "OriginUrl"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/add [post]
func AddUrl(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		url  Url
	)

	httpCode, errCode := app.BindAndValidJson(c, &url)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	urlService := url_service.Url_Service{OriginUrl: url.OriginUrl}
	exists, err := urlService.ExistsByOrigin()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_URL_FAIL, nil)
		return
	}

	if exists {
		urlRecord, err := urlService.GetUrlByOrigin()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_URL_FAIL, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, urlRecord.ShorternUrl)
		return
	} else {

		// else create a new record.

		urlService.ShorternUrl = shorten.Shorten(urlService.OriginUrl)
		urlService.CreatedBy = 1

		if err = urlService.Add(); err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_ADD_URL_FAIL, nil)
			return
		}

		appG.Response(http.StatusOK, e.SUCCESS, urlService.ShorternUrl)
	}
}

// @Summary Delete URL
// @Produce json
// @Param url path string true "ShortenUrl"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/{shorten} [delete]
func DeleteUrl(c *gin.Context) {
	appG := app.Gin{C: c}

	url := c.Param("url")

	urlService := url_service.Url_Service{ShorternUrl: url}

	exists, err := urlService.ExistByShort()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_URL_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_URL, nil)
	}

	err = urlService.DeleteByShortern()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_URL_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary Get URL
// @Param url path string true "ShortenUrl"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/{shorten} [get]
func GetUrl(c *gin.Context) {
	appG := app.Gin{C: c}

	urlService := url_service.Url_Service{ShorternUrl: c.Param("url")}
	logging.Info(c.Params)

	exists, err := urlService.ExistByShort()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_URL_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusNotFound, e.ERROR_EXIST_URL_FAIL, nil)
		return
	}

	// get by shorten ? id ???
	url, err := urlService.Get(urlService.ShorternUrl)
	appG.Response(http.StatusOK, e.SUCCESS, url.OriginUrl)
	return
}
