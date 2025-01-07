package config

import "strings"

func processURL(url string, defaultUrl ...string) string {
	if len(url) == 0 && len(defaultUrl) == 1 {
		url = defaultUrl[0]
	}

	url = strings.TrimSpace(url)
	url = strings.TrimRight(url, "/")
	if !strings.HasPrefix(url, "/") {
		url = "/" + url
	}

	return url
}
