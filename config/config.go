package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP    `yaml:"http"`
		PG      `yaml:"postgres"`
		Auth    `yaml:"auth"`
		Storage `yaml:"storage"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
	}

	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max"`
		URL     string `env-required:"true" yaml:"pg_url"`
	}

	Auth struct {
		SigningKey      string        `env-required:"true" yaml:"jwt_signing_key"`
		PasswordSalt    string        `env-required:"true" yaml:"password_salt"`
		AccessTokenTTL  time.Duration `env-required:"true" yaml:"access_token_ttl"`
		RefreshTokenTTL time.Duration `env-required:"true" yaml:"refresh_token_ttl"`
	}

	Storage struct {
		Endpoint string `env-required:"true" yaml:"endpoint"`
		Id       string `env-required:"true" yaml:"id"`
		Secret   string `env-required:"true" yaml:"secret"`
		Region   string `env-required:"true" yaml:"region"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
