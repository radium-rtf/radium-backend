package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
		Auth `yaml:"auth"`
		Storage
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
		Id       string `env-required:"true" env:"STORAGE_ID"`
		Endpoint string `env-required:"true" env:"STORAGE_ENDPOINT"`
		Secret   string `env-required:"true" env:"STORAGE_SECRET"`
		Region   string `env-required:"true" env:"STORAGE_REGION"`
	}
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = cleanenv.ReadEnv(&cfg.Storage)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
