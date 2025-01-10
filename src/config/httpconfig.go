package config

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type HttpConfig struct {
	Address        string           `yaml:"address"`
	DebugMsg       utils.StringBool `yaml:"debugmsg"`
	BasePath       string           `yaml:"basepath"`
	ApiPath        string           `yaml:"apipath"`
	ResourcePath   string           `yaml:"resourcepath"`
	EnableTestAPI  string           `yaml:"enabletestapi"`
	Proxy          ProxyConfig      `yaml:"proxy"`
	StopSecret     string           `yaml:"stopsecret"`
	StopWaitSecond int              `yaml:"stopwaitsecond"`
	Cors           CorsConfig       `yaml:"cors"`
}

func (h *HttpConfig) setDefault(global *GlobalConfig) {
	if h.Address == "" {
		h.Address = "localhost:2689"
	}

	if global.IsDebug() || global.IsTest() {
		h.DebugMsg.SetDefaultEanble()
	} else {
		h.DebugMsg.SetDefaultDisable()
	}

	h.BasePath = utils.ProcessPath(h.BasePath)
	h.ApiPath = utils.ProcessPath(h.ApiPath, "/api")
	h.ResourcePath = utils.ProcessPath(h.ResourcePath, "/resource")

	if h.StopSecret == "" {
		h.StopSecret = utils.RandStr(8)
		NewConfigWarning(fmt.Sprintf("Auto set http stop secret %s\n", h.StopSecret))
	}

	if h.StopWaitSecond <= 0 {
		h.StopWaitSecond = 10
	}

	h.Proxy.setDefault(global)
	h.Cors.setDefault()
}

func (h *HttpConfig) check(co *CorsOrigin) ConfigError {
	if len(h.ApiPath) == 0 {
		return NewConfigError("Api Path is empty")
	}

	if len(h.ResourcePath) == 0 {
		return NewConfigError("Resource Path is empty")
	}

	if !utils.IsValidURLPath(h.BasePath) {
		return NewConfigError("http base path is not valid")
	}

	if !utils.IsValidURLPath(h.ResourcePath) {
		return NewConfigError("http resource path is not valid")
	}

	if !utils.IsValidURLPath(h.ApiPath) {
		return NewConfigError("http api path is not valid")
	}

	if !utils.IsValidURLPath(h.BasePath + h.ResourcePath) {
		return NewConfigError("http resource path is not valid")
	}

	if !utils.IsValidURLPath(h.BasePath + h.ApiPath) {
		return NewConfigError("http api path is not valid")
	}

	err := h.Proxy.check()
	if err != nil && err.IsError() {
		return err
	}

	err = h.Cors.check(co)
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
