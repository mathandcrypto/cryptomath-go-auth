package databaseConfig

import (
	"fmt"
	configErrors "github.com/mathandcrypto/cryptomath-go-auth/internal/common/errors/config"
	"github.com/spf13/viper"
)

type Config struct {
	Host	string
	Port	int16
	User	string
	Password	string
	Database	string
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Host, c.User, c.Password, c.Database, c.Port)
}

func New() (*Config, error) {
	dbViper := viper.New()

	dbViper.AddConfigPath("configs/database")
	dbViper.SetConfigName("config")
	dbViper.SetConfigType("env")
	dbViper.AutomaticEnv()

	if err := dbViper.ReadInConfig(); err != nil {
		return nil, &configErrors.ReadConfigError{ConfigName: "database", ViperErr: err}
	}

	//	Load database host
	dbHost := dbViper.GetString("DATABASE_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	//	Load database port
	dbPort := dbViper.GetInt("DATABASE_PORT")
	if dbPort == 0 {
		dbPort = 5432
	}

	//	Load database user
	dbUser := dbViper.GetString("POSTGRES_USER")
	if dbUser == "" {
		return nil, &configErrors.EmptyKeyError{Key: "POSTGRES_USER"}
	}

	//	Load database password
	dbPassword := dbViper.GetString("POSTGRES_PASSWORD")
	if dbPassword == "" {
		return nil, &configErrors.EmptyKeyError{Key: "POSTGRES_PASSWORD"}
	}

	//	Load database name
	dbName := dbViper.GetString("POSTGRES_DB")
	if dbName == "" {
		return nil, &configErrors.EmptyKeyError{Key: "POSTGRES_DB"}
	}

	return &Config{
		Host: dbHost,
		Port: int16(dbPort),
		User: dbUser,
		Password: dbPassword,
		Database: dbName,
	}, nil
}
