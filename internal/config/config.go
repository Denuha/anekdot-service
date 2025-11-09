package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      int    `envconfig:"PORT" required:"true"`
	DBConnStr string `envconfig:"DB_CONN_STR" required:"true"`

	Debug         bool   `envconfig:"AUTH_DEBUG" default:"false"`
	TelegramToken string `envconfig:"TELEGRAM_TOKEN" required:"true"`
	TelegramOn    bool   `envconfig:"TELEGRAM_ON" default:"true"`

	UserPasswordSalt string `envconfig:"USER_PASSWORD_SALT" default:"AAA"`

	TokenSignedKey string        `envconfig:"TOKEN_SIGNED_KEY" default:"AAA"`
	TokenExpires   time.Duration `envconfig:"TOKEN_EXPIRES" default:"15m"`

	TokenRefreshSignedKey string        `envconfig:"TOKEN_REFRESH_SIGNING_KEY" default:"BBB"`
	TokenRefreshExpires   time.Duration `envconfig:"TOKEN_REFRESH_EXPIRES" default:"24h"`

	SwaggerBasePath string `envconfig:"SWAGGER_BASE_PATH" default:""`
	SwaggerHost     string `envconfig:"SWAGGER_HOST" default:""`
	SwaggerVersion  string `envconfig:"SWAGGER_VERSION" default:""`
}

func InitConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("", &conf)
	return &conf, err
}
