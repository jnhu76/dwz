package services

import "net/url"

func ParseUrl(str string) bool {
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
