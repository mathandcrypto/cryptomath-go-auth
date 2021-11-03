package redisConfig

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
	redisViper := viper.New()

	redisViper.AddConfigPath("configs/redis")
	redisViper.SetConfigName("config")
	redisViper.SetConfigType("env")
	redisViper.AutomaticEnv()

	if err := redisViper.ReadInConfig(); err != nil {
		return nil, &configErrors.ReadConfigError{ConfigName: "redis", ViperErr: err}
	}

	//	Load redis host
	redisHost := redisViper.GetString("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost"
	}

	//	Load redis port
	redisPort := redisViper.GetInt("REDIS_PORT")
	if redisPort == 0 {
		redisPort = 6379
	}

	return &Config{
		Host: redisHost,
		Port: int16(redisPort),
	}, nil
}
