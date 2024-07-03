package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type PostgresConfig struct {
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     string `envconfig:"POSTGRES_PORT"`
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	Database string `envconfig:"POSTGRES_DATABASE"`
}

type Config struct {
	Postgres        PostgresConfig
	UserServiceHost string `envconfig:"USERSERVICE_HOST"`
	UserServicePort string `envconfig:"USERSERVICE_PORT"`
	ServiceHost     string `envconfig:"SERVICE_HOST"`
	ServicePort     string `envconfig:"SERVICE_PORT"`
}

func LoadConfig(path string) Config {
	err := godotenv.Load(path + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("Error processing envconfig: %v", err)
	}

	return cfg
}
