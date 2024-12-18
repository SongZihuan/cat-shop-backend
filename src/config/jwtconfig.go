package config

import (
	"github.com/SuperH-0630/cat-shop-back/src/utils"
	"os"
	"strings"
	"time"
)

type JwtConfig struct {
	Secret     string        `json:"secret"`
	SecretPath string        `json:"secretpath"`
	SaveSecret bool          `json:"savesecret"`
	Hour       time.Duration `json:"hour"`
	ResetMin   time.Duration `json:"resetmin"`
	Issuer     string        `json:"issuer"`
}

func (j *JwtConfig) setDefault() {
	if j.SecretPath == "" {
		j.SecretPath = "./.jwt.secret"
	}

	if j.Secret == "" {
		tmp := loadJwtSecret(j.SecretPath)
		if tmp == "" {
			j.Secret = utils.RandStr(20)
		} else {
			j.Secret = tmp
		}
	}

	if j.Hour <= 0 {
		j.Hour = 2
	}

	if j.ResetMin*60 > j.Hour {
		j.ResetMin -= 15
	}

	if j.ResetMin <= 0 {
		j.ResetMin = 0
	}

	if j.Issuer == "" {
		j.Issuer = "eShopAsTest"
	}

	j.Issuer = strings.ReplaceAll(j.Issuer, " ", "")
	j.Issuer = strings.ToUpper(j.Issuer)
}

func (j *JwtConfig) check() ConfigError {
	if j.SaveSecret {
		err := saveJwtSecret(j.SecretPath, j.Secret)
		if err != nil {
			_ = NewConfigWarning("jwt secret save warning:" + err.Error())
		}
	}

	return nil
}

func loadJwtSecret(path string) string {
	secret, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(secret)
}

func saveJwtSecret(path string, secret string) error {
	return os.WriteFile(path, []byte(secret), 0644)
}
