package config

import (
	"fmt"
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"strings"
)

type HttpConfig struct {
	Address           string      `yaml:"address"`
	DebugMsg          bool        `yaml:"debugmsg"`
	ApiBaseAPI        string      `yaml:"apibaseapi"`
	ResourceBaseAPI   string      `yaml:"resourcebaseapi"`
	TestApi           bool        `yaml:"testapi"`
	Proxy             ProxyConfig `yaml:"proxy"`
	StopSecret        string      `yaml:"stopsecret"`
	StopWaitSecond    int         `yaml:"stopwaitsecond"`
	RestartWaitSecond int         `yaml:"restartwaitsecond"`
	WaitStartSecond   int         `yaml:"waitstartsecond"`
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

	if h.StopSecret == "" {
		h.StopSecret = utils.RandStr(8)
		fmt.Printf("Auto set http stop secret %s\n", h.StopSecret)
	}

	if h.StopWaitSecond <= 0 {
		h.StopWaitSecond = 10
	}

	if h.RestartWaitSecond <= 0 {
		h.StopWaitSecond = 20
	}

	if h.WaitStartSecond <= 0 {
		h.WaitStartSecond = 20
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

	if h.RestartWaitSecond <= h.StopWaitSecond {
		_ = NewConfigWarning("RestartWaitSecond is greater than StopWaitSecond")
	}

	return nil
}

func (h *HttpConfig) CheckStopSecret(secret string) bool {
	return h.StopSecret == secret
}
