package config

import "strings"

type HttpConfig struct {
	Address         string      `yaml:"address"`
	DebugMsg        bool        `yaml:"debugmsg"`
	ApiBaseAPI      string      `yaml:"apibaseapi"`
	ResourceBaseAPI string      `yaml:"resourcebaseapi"`
	TestApi         bool        `yaml:"testapi"`
	Proxy           ProxyConfig `yaml:"proxy"`
}

func (h *HttpConfig) setDefault() {
	if h.Address == "" {
		h.Address = "localhost:2689"
	}

	h.ApiBaseAPI = strings.TrimSpace(h.ApiBaseAPI)

	if h.ApiBaseAPI == "" {
		h.ApiBaseAPI = "/api"
	} else {
		if !strings.HasPrefix(h.ApiBaseAPI, "/") {
			h.ApiBaseAPI = "/" + h.ApiBaseAPI
		}

		if strings.HasSuffix(h.ApiBaseAPI, "/") {
			h.ApiBaseAPI = strings.TrimRight(h.ApiBaseAPI, "/")
		}
	}

	if h.ResourceBaseAPI == "" {
		h.ResourceBaseAPI = "/file"
	} else {
		if !strings.HasPrefix(h.ResourceBaseAPI, "/") {
			h.ResourceBaseAPI = "/" + h.ResourceBaseAPI
		}

		if strings.HasSuffix(h.ResourceBaseAPI, "/") {
			h.ResourceBaseAPI = strings.TrimRight(h.ResourceBaseAPI, "/")
		}
	}

	h.Proxy.setDefault()
}

func (h *HttpConfig) check() ConfigError {
	err := h.Proxy.check()
	if err != nil && err.IsError() {
		return err
	}

	return nil
}
