package cache_service

import (
	"crypto/sha256"
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

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return string(bs)
}

func (u *UrlCache) GetUrlOriginKey() string {
	return e.CACHE_URL + "_" + hash(u.OriginUrl)
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
