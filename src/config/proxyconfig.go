package config

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"strings"
)

type ProxyConfig struct {
	Proxy      string   `json:"proxy"`
	TrustedIPs []string `json:"trustedips"`
}

func (p *ProxyConfig) setDefault() {
	if p.Proxy == "" {
		p.Proxy = "enable"
	}

	p.Proxy = strings.ToLower(p.Proxy)
}

func (p *ProxyConfig) check() ConfigError {
	if p.Proxy != "enable" && p.Proxy != "disable" {
		return NewConfigError("proxy must be enable/disable")
	}

	if p.Proxy == "enable" {
		if len(p.TrustedIPs) == 0 {
			_ = NewConfigWarning("proxy trusts ips will be ignore because proxy is disabled")
		} else {
			for _, ip := range p.TrustedIPs {
				if !utils.ValidIPv4(ip) && !utils.ValidIPv6(ip) && !utils.IsValidIPv4CIDR(ip) && !utils.IsValidIPv6CIDR(ip) {
					return NewConfigError(fmt.Sprintf("bad proxy trusts ip address: %s", ip))
				}
			}
		}
	} else {
		_ = NewConfigWarning("You trusted all proxies, this is NOT safe. We recommend you to set a value.")
	}

	return nil
}

func (p *ProxyConfig) Enable() bool {
	return p.Proxy == "enable"
}
