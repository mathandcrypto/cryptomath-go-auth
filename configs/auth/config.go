package authConfig

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	configErrors "github.com/mathandcrypto/cryptomath-go-auth/internal/common/errors/config"
)

type Config struct {
	AccessSessionMaxAge int32	`mapstructure:"ACCESS_SESSION_MAX_AGE" validate:"required,gte=1,lte=60"`
	RefreshSessionMaxAge int32	`mapstructure:"REFRESH_SESSION_MAX_AGE" validate:"required,gte=1,lte=30"`
	MaxRefreshSessions	int64	`mapstructure:"MAX_REFRESH_SESSIONS" validate:"required,gte=1,lte=10"`
}

func (c *Config) RefreshSessionExpirationDate() time.Time {
	now := time.Now()
	now.Add((-24) * time.Duration(c.RefreshSessionMaxAge) * time.Hour)

	return now
}

func New() (*Config, error) {
	authViper := viper.New()
	authValidate := validator.New()

	authViper.SetDefault("MAX_REFRESH_SESSIONS", 5)

	authViper.AddConfigPath("configs/auth")
	authViper.SetConfigName("config")
	authViper.SetConfigType("env")
	authViper.AutomaticEnv()

	if err := authViper.ReadInConfig(); err != nil {
		return nil, &configErrors.ReadConfigError{ConfigName: "auth", ViperErr: err}
	}

	var authConfig Config
	if err := authViper.Unmarshal(&authConfig); err != nil {
		return nil, &configErrors.UnmarshalError{ConfigName: "auth", ViperErr: err}
	}

	if err := authValidate.Struct(authConfig); err != nil {
		return nil, &configErrors.ValidationError{ConfigName: "auth", ValidateErr: err}
	}

	return &authConfig, nil
}