package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Env struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	AppEnv      string `mapstructure:"APP_ENV"`
	AppAddr     string `mapstructure:"APP_ADDR"`
}

func NewEnv() (Env, error) {
	viper.SetDefault("APP_ENV", "development")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(".env")
	_ = viper.ReadInConfig()

	var env Env

	if err := viper.Unmarshal(&env); err != nil {
		return Env{}, err
	}
	if env.DatabaseURL == "" {
		return Env{}, fmt.Errorf("DATABASE_URL is required")
	}

	return env, nil
}
