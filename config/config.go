package config

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Radium     Radium     `yaml:"radium" env-prefix:"radium"`
		Wave       Wave       `yaml:"wave" env-prefix:"wave"`
		PG         PG         `yaml:"postgres"`
		Storage    Storage    `yaml:"storage"`
		Smtp       Smtp       `yaml:"smtp"`
		Centrifugo Centrifugo `yaml:"centrifugo"`
		JWT        JWT
	}

	Radium struct {
		HTTP           HTTP      `yaml:"http" env-prefix:"http"`
		Auth           Auth      `yaml:"auth" env-prefix:"auth"`
		DefaultGroupID uuid.UUID `env:"DEFAULT_GROUP_ID" env-default:"81af02da-bf9e-4769-aa07-36903517733d"`
	}

	Wave struct {
		HTTP HTTP `yaml:"http" env-prefix:"http"`
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
		PasswordSaltSha1       string        `env:"PASSWORD_SALT_SHA1" env-default:"ГОД НАЗАД Я СИДЕЛ НА СКАМЕЙКЕ В ОБЩЕСТВЕННОМ ПАРКЕ"`
		PasswordCostBcrypt     int           `env-required:"true" yaml:"password_cost_bcrypt"`
		AccessTokenTTL         time.Duration `env-required:"true" yaml:"access_token_ttl"`
		RefreshTokenTTL        time.Duration `env-required:"true" yaml:"refresh_token_ttl"`
		LengthVerificationCode int           `env-required:"true" yaml:"length_verification_code"`
	}

	Storage struct {
		PrivateEndpoint string `env:"STORAGE_PRIVATE_ENDPOINT" env-default:"minio:9000"`
		Endpoint        string `env:"STORAGE_ENDPOINT" env-default:"localhost:9000"`
		Id              string `env:"STORAGE_ID" env-default:"useruseruser"`
		Secret          string `env:"STORAGE_SECRET" env-default:"useruseruser"`
		Region          string `env:"STORAGE_REGION" env-default:"RU"`
	}

	Smtp struct {
		Host string `env-required:"true" yaml:"host" env:"SMTP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"SMTP_PORT"`

		Email    string `env:"SMTP_EMAIL" env-default:"wdkadwadwakpklrbjb@urfu.me"`
		Password string `env:"SMTP_PASSWORD" env-default:"паыауыаыулажыдула"`
		Username string `env:"SMTP_USERNAME" env-default:"noreply@khostya.online"`
	}

	Centrifugo struct {
		Token  string `env:"CENTRIFUGO_TOKEN" env-default:"5d12cdf5-252f-4b04-a4e1-37aec016ef5c"`
		ApiKey string `env:"CENTRIFUGO_API_KEY" env-default:"373b4d58-fe26-40a8-b5fc-00785ffc8450"`
	}

	JWT struct {
		SigningKey string `env:"JWT_SIGNING_KEY" env-default:"wdkadwadwakpklrbjb"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	configPath := "./config/config.yml"
	path, exists := os.LookupEnv("CONFIG_PATH")
	if exists {
		configPath = path
	}

	err := cleanenv.ReadConfig(configPath, cfg)
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
