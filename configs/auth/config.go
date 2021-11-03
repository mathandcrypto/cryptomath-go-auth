package authConfig

import (
	"github.com/go-playground/validator/v10"
	configErrors "github.com/mathandcrypto/cryptomath-go-auth/internal/common/errors/config"
	"github.com/spf13/viper"
)

type Config struct {
	AccessSessionMaxAge int32	`validate:"gte=1,lte=60"`
	RefreshSessionMaxAge int32	`validate:"gte=1,lte=30"`
	MaxRefreshSessions	int64	`validate:"gte=1,lte=10"`
}

func New() (*Config, error) {
	authViper := viper.New()
	authValidate := validator.New()

	authViper.AddConfigPath("configs/auth")
	authViper.SetConfigName("config")
	authViper.SetConfigType("env")
	authViper.AutomaticEnv()

	if err := authViper.ReadInConfig(); err != nil {
		return nil, &configErrors.ReadConfigError{ConfigName: "auth", ViperErr: err}
	}

	//	Load auth access session max age
	authAccessSessionMaxAge := authViper.GetInt32("ACCESS_SESSION_MAX_AGE")
	if authAccessSessionMaxAge == 0 {
		return nil, &configErrors.EmptyKeyError{Key: "ACCESS_SESSION_MAX_AGE"}
	}

	//	Load auth refresh session max age
	authRefreshSessionMaxAge := authViper.GetInt32("REFRESH_SESSION_MAX_AGE")
	if authRefreshSessionMaxAge == 0 {
		return nil, &configErrors.EmptyKeyError{Key: "REFRESH_SESSION_MAX_AGE"}
	}

	//	Load max refresh sessions
	authMaxRefreshSessions := authViper.GetInt64("MAX_REFRESH_SESSIONS")
	if authMaxRefreshSessions == 0 {
		authMaxRefreshSessions = 5
	}

	authConfig := &Config{
		AccessSessionMaxAge: authAccessSessionMaxAge,
		RefreshSessionMaxAge: authRefreshSessionMaxAge,
		MaxRefreshSessions: authMaxRefreshSessions,
	}

	if err := authValidate.Struct(authConfig); err != nil {
		return nil, &configErrors.ValidationError{ConfigName: "auth", ValidateErr: err}
	}

	return authConfig, nil
}