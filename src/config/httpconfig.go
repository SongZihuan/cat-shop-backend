package config

import "strings"

type HttpConfig struct {
	Address         string      `yaml:"address"`
	DebugMsg        bool        `yaml:"debugmsg"`
	BaseAPI         string      `yaml:"baseapi"`
	ResourceBaseAPI string      `yaml:"resourcebaseapi"`
	TestApi         bool        `yaml:"testapi"`
	Proxy           ProxyConfig `yaml:"proxy"`
}

func (h *HttpConfig) setDefault() {
	if h.Address == "" {
		h.Address = "localhost:2689"
	}

	h.BaseAPI = strings.TrimSpace(h.BaseAPI)

	if h.BaseAPI == "" {
		h.BaseAPI = "/api"
	} else {
		if !strings.HasPrefix(h.BaseAPI, "/") {
			h.BaseAPI = "/" + h.BaseAPI
		}

		if strings.HasSuffix(h.BaseAPI, "/") {
			h.BaseAPI = strings.TrimRight(h.BaseAPI, "/")
		}
	}

	if h.ResourceBaseAPI == "" {
		h.ResourceBaseAPI = h.BaseAPI
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
	if !strings.HasSuffix(h.ResourceBaseAPI, h.BaseAPI) {
		_ = NewConfigWarning("resource base api has not suffix (base api)")
	}

	err := h.Proxy.check()
	if err != nil && err.IsError() {
		return err
	}

	return nil
}
