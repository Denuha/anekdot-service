package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port      int    `envconfig:"PORT" required:"true"`
	DBConnStr string `envconfig:"DB_CONN_STR" required:"true"`

	Debug bool `envconfig:"AUTH_DEBUG" default:"false"`
}

func InitConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("", &conf)
	return &conf, err
}
