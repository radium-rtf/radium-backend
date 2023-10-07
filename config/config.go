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
		Port           string        `env-required:"true" yaml:"port"`
		ReadTimeout    time.Duration `env-required:"true" yaml:"read_timeout"`
		WriteTimeout   time.Duration `env-required:"true" yaml:"write_timeout"`
		IdleTimeout    time.Duration `env-required:"true" yaml:"idle_timeout"`
		MaxHeaderBytes int           `env-required:"true" yaml:"max_header_bytes"`
	}

	PG struct {
		URL             string        `env-required:"true" yaml:"pg_url"`
		MaxOpenConns    int           `env-required:"true" yaml:"max_open_conns"`
		MaxIdleConns    int           `env-required:"true" yaml:"max_idle_conns"`
		ConnMaxIdleTime time.Duration `env-required:"true" yaml:"conn_max_idle_time"`
		ConnMaxLifetime time.Duration `env-required:"true" yaml:"conn_max_lifetime"`
	}

	Auth struct {
		SigningKey         string        `env-required:"true" yaml:"jwt_signing_key"`
		PasswordSaltSha1   string        `env-required:"true" yaml:"password_salt_sha1"`
		PasswordCostBcrypt int           `env-required:"true" yaml:"password_cost_bcrypt"`
		AccessTokenTTL     time.Duration `env-required:"true" yaml:"access_token_ttl"`
		RefreshTokenTTL    time.Duration `env-required:"true" yaml:"refresh_token_ttl"`
	}

	Storage struct {
		PrivateEndpoint string `env-required:"true" yaml:"private_endpoint"`
		Endpoint        string `env-required:"true" yaml:"endpoint"`
		Id              string `env-required:"true" yaml:"id"`
		Secret          string `env-required:"true" yaml:"secret"`
		Region          string `env-required:"true" yaml:"region"`
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
