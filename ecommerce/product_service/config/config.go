package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Postgres            PostgresConfig
	ProductServiceHost  string
	ProductServicePort  string
	ServiceHost         string
	ServicePort         string
}

func Load(path string) Config {
	godotenv.Load(path + "/.env")

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		Postgres: PostgresConfig{
			Host:     conf.GetString("POSTGRES_HOST"),
			Port:     conf.GetString("POSTGRES_PORT"),
			User:     conf.GetString("POSTGRES_USER"),
			Password: conf.GetString("POSTGRES_PASSWORD"),
			Database: conf.GetString("POSTGRES_DATABASE"),
		},
		ProductServiceHost: conf.GetString("PRODUCTSERVICE_HOST"),
		ProductServicePort: conf.GetString("PRODUCTSERVICE_PORT"),
		ServiceHost:        conf.GetString("SERVICE_HOST"),
		ServicePort:        conf.GetString("SERVICE_PORT"),
	}

	return cfg
}
