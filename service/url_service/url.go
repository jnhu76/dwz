package url_service

import (
	"encoding/json"

	"github.com/jnhu76/dwz/models"
	"github.com/jnhu76/dwz/pkg/gredis"
	"github.com/jnhu76/dwz/pkg/logging"
	"github.com/jnhu76/dwz/service/cache_service"
)

type Url_Service struct {
	ID          int
	OriginUrl   string
	ShorternUrl string
	CreatedBy   int
}

func (u *Url_Service) Add() error {
	url := map[string]interface{}{
		"origin_url":  u.OriginUrl,
		"shorter_url": u.ShorternUrl,
		"created_by":  u.CreatedBy,
	}

	if err := models.AddUrl(url); err != nil {
		return err
	}

	return nil
}

// id or short_url
func (u *Url_Service) Get(shorten string) (*models.Url, error) {
	var cacheUrl *models.Url

	cache := cache_service.UrlCache{ShorterUrl: u.ShorternUrl}
	key := cache.GetUrlKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUrl)
			return cacheUrl, nil
		}
	}

	url, err := models.GetUrlByShort(shorten)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, url, 3600)
	return url, nil
}

func (u *Url_Service) GetAll() ([]*models.Url, error) {
	var (
		urls, cacheUrls []*models.Url
	)

	cache := cache_service.UrlCache{
		ID:         u.ID,
		OriginUrl:  u.OriginUrl,
		ShorterUrl: u.ShorternUrl,
		Creator:    u.CreatedBy,
	}

	key := cache.GetUrlsKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUrls)
			return cacheUrls, nil
		}
	}

	urls, err := models.GetUrls(u.CreatedBy)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, urls, 3600)
	return urls, nil
}

func (u *Url_Service) DeleteByShortern() error {
	return models.DeleteUrl(u.ShorternUrl)
}

func (u *Url_Service) ExistByShort() (bool, error) {
	return models.ExistUrlByShortern(u.ShorternUrl)
}

func (u *Url_Service) ExistsByOrigin() (bool, error) {
	return models.ExistUrlByOrigin(u.OriginUrl)
}

func (u *Url_Service) GetUrlByOrigin() (*models.Url, error) {
	var cacheUrl *models.Url

	cache := cache_service.UrlCache{OriginUrl: u.OriginUrl}
	key := cache.GetUrlOriginKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUrl)
			return cacheUrl, nil
		}
	}
	url, err := models.GetShortByOrigin(u.OriginUrl)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, url, 3600)
	return url, nil
}
