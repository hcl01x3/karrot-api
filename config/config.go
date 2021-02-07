package config

import (
	"time"

	"github.com/octo-5/karrot-api/util"
)

type Config struct {
	APIPort      int           `json:"API_PORT" validate:"gte=0,lte=65535"`
	APIDomain    string        `json:"API_DOMAIN" validate:"url"`
	DBDSN        string        `json:"DB_DSN" validate:"url"`
	AccessExpiry time.Duration `json:"ACCESS_EXPIRY_HRS" validate:"gte=0"`
	JWTSecret    string        `json:"JWT_SECRET" validate:"min=20,max=255"`
	LogLevel     string        `json:"LOG_LEVEL" validate:"oneof=debug info warn error"`
	IsProduction bool
}

func Load() *Config {
	return &Config{
		APIPort:      util.MustGetIntEnv("API_PORT"),
		APIDomain:    util.MustGetEnv("API_DOMAIN"),
		DBDSN:        util.MustGetEnv("DB_DSN"),
		AccessExpiry: time.Duration(util.MustGetIntEnv("ACCESS_EXPIRY_MINS")) * time.Minute,
		JWTSecret:    util.MustGetEnv("JWT_SECRET"),
		LogLevel:     util.MustGetEnv("LOG_LEVEL"),
		IsProduction: util.MustGetEnv("LOG_LEVEL") != "debug",
	}
}
