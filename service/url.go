package service

import (
	"encoding/json"

	"github.com/jnhu76/dwz/models"
	"github.com/jnhu76/dwz/pkg/gredis"
	"github.com/jnhu76/dwz/pkg/logging"
)

type Url struct {
	ID         int
	OriginUrl  string
	ShorterUrl string
	CreatedBy  int
}

func (u *Url) Add() error {
	url := map[string]interface{}{
		"origin_url":  u.OriginUrl,
		"shorter_url": u.ShorterUrl,
		"created_by":  u.CreatedBy,
	}

	if err := models.AddUrl(url); err != nil {
		return err
	}

	return nil
}

// id or short_url
func (u *Url) Get() (*models.Url, error) {
	var cacheUrl *models.Url

	cache := UrlCache{ID: u.ID}
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

	url, err := models.GetUrl(u.ID)
	if err != nil {
		return nil, err
	}
	gredis.Set(key, url, 3600)
	return url, nil
}

func (u *Url) GetAll() ([]*models.Url, error) {
	var (
		urls, cacheUrls []*models.Url
	)

	cache := UrlCache{
		ID:         u.ID,
		OriginUrl:  u.OriginUrl,
		ShorterUrl: u.ShorterUrl,
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

func (u *Url) Delete() error {
	return models.DeleteUrl(u.ID)
}

func (u *Url) ExistByShort() (bool, error) {
	return models.ExistUrlByOrigin(u.OriginUrl)
}
