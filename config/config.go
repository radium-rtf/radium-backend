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
		Smtp    `yaml:"smtp"`
	}

	HTTP struct {
		Port           string        `env-required:"true" yaml:"port"`
		ReadTimeout    time.Duration `env-required:"true" yaml:"read_timeout"`
		WriteTimeout   time.Duration `env-required:"true" yaml:"write_timeout"`
		IdleTimeout    time.Duration `env-required:"true" yaml:"idle_timeout"`
		MaxHeaderBytes int           `env-required:"true" yaml:"max_header_bytes"`
	}

	PG struct {
		URL             string        `env:"PG_URL" env-default:"postgres://user:password@postgres:5432/postgres?sslmode=disable"`
		MaxOpenConns    int           `env-required:"true" yaml:"max_open_conns"`
		MaxIdleConns    int           `env-required:"true" yaml:"max_idle_conns"`
		ConnMaxIdleTime time.Duration `env-required:"true" yaml:"conn_max_idle_time"`
		ConnMaxLifetime time.Duration `env-required:"true" yaml:"conn_max_lifetime"`
	}

	Auth struct {
		SigningKey         string        `env:"JWT_SIGNING_KEY" env-default:"wdkadwadwakpklrbjb"`
		PasswordSaltSha1   string        `env:"PASSWORD_SALT_SHA1" env-default:"ГОД НАЗАД Я СИДЕЛ НА СКАМЕЙКЕ В ОБЩЕСТВЕННОМ ПАРКЕ"`
		PasswordCostBcrypt int           `env-required:"true" yaml:"password_cost_bcrypt"`
		AccessTokenTTL     time.Duration `env-required:"true" yaml:"access_token_ttl"`
		RefreshTokenTTL    time.Duration `env-required:"true" yaml:"refresh_token_ttl"`
	}

	Storage struct {
		PrivateEndpoint string `env:"STORAGE_PRIVATE_ENDPOINT" env-default:"minio:9000"`
		Endpoint        string `env:"STORAGE_ENDPOINT" env-default:"localhost:9000"`
		Id              string `env:"STORAGE_ID" env-default:"useruseruser"`
		Secret          string `env:"STORAGE_SECRET" env-default:"useruseruser"`
		Region          string `env:"STORAGE_REGION" env-default:"RU"`
	}

	Smtp struct {
		Host                   string `env-required:"true" yaml:"host" env:"SMTP_HOST"`
		Port                   int    `env-required:"true" yaml:"port" env:"SMTP_PORT"`
		LengthVerificationCode int    `env-required:"true" yaml:"length_verification_code"`

		Email    string `env:"SMTP_EMAIL" env-default:"wdkadwadwakpklrbjb@urfu.me"`
		Password string `env:"SMTP_PASSWORD" env-default:"паыауыаыулажыдула"`
		Username string `env:"SMTP_USERNAME" env-default:"noreply@khostya.online"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func MustConfig() *Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
