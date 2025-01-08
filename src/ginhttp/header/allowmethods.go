package header

import "net/http"

var AllowMethods = []string{
	http.MethodGet,
	http.MethodPost,
	http.MethodOptions,
}

func IsAllowMethods(method string) bool {
	for _, m := range AllowMethods {
		if m == method {
			return true
		}
	}
	return false
}
