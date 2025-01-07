package config

import (
	"fmt"
	"github.com/SongZihuan/cat-shop-backend/src/utils"
)

type ProxyConfig struct {
	Proxy      StringBool `json:"proxy"`
	TrustedIPs []string   `json:"trustedips"`
}

func (p *ProxyConfig) setDefault() {
	p.Proxy.SetDefault(Disable)
}

func (p *ProxyConfig) check() ConfigError {
	if p.Proxy.Is(Enable) {
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
	return p.Proxy.Is(Enable)
}
