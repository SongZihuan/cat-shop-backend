package config

type PasswordConfig struct {
	Front   string `yaml:"front"`
	Backend string `yaml:"backend"`
}

func (p *PasswordConfig) setDefault() {
	if p.Front == "" && p.Backend != "" {
		p.Front = p.Backend
	} else if p.Backend == "" && p.Front != "" {
		p.Backend = p.Front
	}
}

func (p *PasswordConfig) check() ConfigError {
	if p.Front == "" {
		return NewConfigError("password hash front salt must be set")
	}

	if p.Front == "" {
		return NewConfigError("password hash backend salt must be set")
	}
	return nil
}
