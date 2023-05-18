package services

import (
	"fmt"
	"net/url"
)

func ParseUrl(str string) bool {
	fmt.Println(str)
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	switch u.Scheme {
	case "http", "https", "ftp":
		return true
	}
	return false
}
