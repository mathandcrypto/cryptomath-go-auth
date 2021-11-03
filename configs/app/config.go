package appConfig

import (
	"fmt"
	configErrors "github.com/mathandcrypto/cryptomath-go-auth/internal/common/errors/config"
	"github.com/spf13/viper"
)

type Config struct {
	Host	string
	Port	int16
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func New() (*Config, error) {
	appViper := viper.New()

	appViper.AddConfigPath("configs/app")
	appViper.SetConfigName("config")
	appViper.SetConfigType("env")
	appViper.AutomaticEnv()

	if err := appViper.ReadInConfig(); err != nil {
		return nil, &configErrors.ReadConfigError{ConfigName: "app", ViperErr: err}
	}

	//	Load app host
	appHost := appViper.GetString("APP_HOST")
	if appHost == "" {
		appHost = "localhost"
	}

	//	Load app port
	appPort := appViper.GetInt("APP_PORT")
	if appPort == 0 {
		return nil, &configErrors.EmptyKeyError{Key: "APP_PORT"}
	}

	return &Config{
		Host: appHost,
		Port: int16(appPort),
	}, nil
}
