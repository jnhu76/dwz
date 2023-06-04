package service

import (
	"strconv"
	"strings"

	"github.com/jnhu76/dwz/pkg/e"
)

type UrlCache struct {
	ID         int
	OriginUrl  string
	ShorterUrl string
	Creator    int
}

func (u *UrlCache) GetUrlKey() string {
	return e.CACHE_URL + "_" + u.ShorterUrl
}

func (u *UrlCache) GetUrlsKey() string {
	keys := []string{
		e.CACHE_URL,
		"LIST",
	}

	if u.ID > 0 {
		keys = append(keys, strconv.Itoa(u.ID))
	}

	return strings.Join(keys, "_")
}
