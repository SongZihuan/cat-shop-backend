package config

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

const CorsMaxAgeSec = 86400
const CorsDefaultMaxAgeSec = CorsMaxAgeSec

type CorsConfig struct {
	AllowCors      utils.StringBool `json:"allowcors"`
	AllowOrigin    []string         `json:"alloworigin"`
	AllowOriginReg []string         `json:"alloworiginres"`
	MaxAgeSec      int              `json:"maxagesec"`
}

func (c *CorsConfig) setDefault() {
	c.AllowCors.SetDefaultDisable()
	if c.AllowCors.IsEnable() && c.MaxAgeSec == 0 {
		c.MaxAgeSec = CorsMaxAgeSec
	}
}

func (c *CorsConfig) check(co *CorsOrigin) ConfigError {
	if c.AllowCors.IsEnable() {
		if c.MaxAgeSec <= 0 || c.MaxAgeSec > CorsMaxAgeSec {
			return NewConfigError(fmt.Sprintf("cors maxagesec %d is invalid", c.MaxAgeSec))
		}

		co.SetString(c.AllowOrigin)

		for _, r := range c.AllowOriginReg {
			_ = co.ApplyReg(r)
		}
	}
	return nil
}

func (c *CorsConfig) Enable() bool {
	return c.AllowCors.IsEnable()
}

func (c *CorsConfig) Disable() bool {
	return c.AllowCors.IsDisable()
}
