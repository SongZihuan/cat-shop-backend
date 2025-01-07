package config

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type HttpConfig struct {
	Address        string      `yaml:"address"`
	DebugMsg       StringBool  `yaml:"debugmsg"`
	BaseURL        string      `yaml:"baseurl"`
	ApiURL         string      `yaml:"apiurl"`
	ResourceURL    string      `yaml:"resourceurl"`
	EnableTestAPI  string      `yaml:"enabletestapi"`
	Proxy          ProxyConfig `yaml:"proxy"`
	StopSecret     string      `yaml:"stopsecret"`
	StopWaitSecond int         `yaml:"stopwaitsecond"`
}

func (h *HttpConfig) setDefault() {
	if h.Address == "" {
		h.Address = "localhost:2689"
	}

	h.DebugMsg.SetDefault(Disable)

	h.BaseURL = processURL(h.BaseURL)
	h.ApiURL = processURL(h.ApiURL, "/api")
	h.ResourceURL = processURL(h.ResourceURL, "/resource")

	if h.StopSecret == "" {
		h.StopSecret = utils.RandStr(8)
		fmt.Printf("Auto set http stop secret %s\n", h.StopSecret)
	}

	if h.StopWaitSecond <= 0 {
		h.StopWaitSecond = 10
	}

	h.Proxy.setDefault()
}

func (h *HttpConfig) check() ConfigError {
	err := h.Proxy.check()
	if err != nil && err.IsError() {
		return err
	}

	if len(h.StopSecret) < 8 {
		return NewConfigError("StopSecret length less than 8")
	}

	return nil
}

func (h *HttpConfig) CheckStopSecret(secret string) bool {
	return h.StopSecret == secret
}
