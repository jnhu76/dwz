package models

import (
	"net/http"
	"time"
	"url/services"

	"github.com/gin-gonic/gin"
)

type UrlPost struct {
	Url string `json:"url"`
}

type Url struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	OriginUrl  string    `json:"origin" gorm:"index,unique"`
	ShorterUrl string    `json:"shorter" gorm:"index"`
	IsDeleted  bool      `json:"isdeleted"`
	CreatedAt  time.Time `json:"createdat" gorm:"autoCreateTime"`
}

func GetUrl(c *gin.Context) {
	var url Url

	if err := DB.Where("shorter_url = ?", c.Param("url")).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Url not found",
		})
		return
	}

	// c.String(http.StatusOK, url.OriginUrl)
	c.Redirect(302, url.OriginUrl)
	c.Abort()
}

func PostUrl(c *gin.Context) {
	var origin UrlPost
	var url Url
	var err error

	c.BindJSON(&origin)

	if services.ParseUrl(origin.Url) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid url"})
		return
	}

	// if err = DB.Where("origin_url = ?", origin.Url).First(&url).Error; err == nil {
	records := DB.Where("origin_url = ?", origin.Url).First(&url)
	if records.RowsAffected != 0 {
		c.String(http.StatusOK, url.ShorterUrl)
		return
	}

	url = Url{OriginUrl: origin.Url, ShorterUrl: services.GetUUID()}
	if err = DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.String(http.StatusOK, url.ShorterUrl)
}

func DeleteUrl(c *gin.Context) {
	var url Url

	if err := DB.Where("origin = ?", c.Param("url")).First(&url).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Url not found"})
		return
	}

	if err := DB.Delete(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Url deleted"})
}
