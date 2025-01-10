package config

import (
	"github.com/SongZihuan/cat-shop-backend/src/utils"
	"os"
	"strings"
)

type JwtConfig struct {
	Secret     string           `json:"secret"`
	SecretPath string           `json:"secretpath"`
	SaveSecret utils.StringBool `json:"savesecret"`
	Hour       int64            `json:"hour"`
	ResetMin   int64            `json:"resetmin"`
	Issuer     string           `json:"issuer"`
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

	j.SaveSecret.SetDefaultEanble()

	if j.Hour <= 0 {
		j.Hour = 2
	}

	if j.Issuer == "" {
		j.Issuer = "CatShop"
	}

	j.Issuer = strings.ReplaceAll(j.Issuer, " ", "")
}

func (j *JwtConfig) check() ConfigError {
	if j.SaveSecret.IsEnable() {
		err := saveJwtSecret(j.SecretPath, j.Secret)
		if err != nil {
			_ = NewConfigWarning("jwt secret save warning:" + err.Error())
		}
	}

	if j.ResetMin > j.Hour*60 {
		return NewConfigError("jwt secret reset min must less than expire hour")
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
