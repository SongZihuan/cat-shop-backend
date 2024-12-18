package config

import "strings"

type HttpConfig struct {
	Address  string `yaml:"address"`
	DebugMsg bool   `yaml:"debugmsg"`
	BaseAPI  string `yaml:"baseapi"`
	TestApi  bool   `yaml:"testapi"`
}

func (h *HttpConfig) setDefault() {
	if h.Address == "" {
		h.Address = "localhost:2689"
	}

	h.BaseAPI = strings.TrimSpace(h.BaseAPI)
	h.BaseAPI = strings.ToLower(h.BaseAPI)

	if h.BaseAPI == "" {
		h.BaseAPI = "/api"
	}

	if !strings.HasPrefix(h.BaseAPI, "/") {
		h.BaseAPI = "/" + h.BaseAPI
	}

	if strings.HasSuffix(h.BaseAPI, "/") {
		h.BaseAPI = strings.TrimRight(h.BaseAPI, "/")
	}
}

func (h *HttpConfig) check() ConfigError {
	return nil
}
