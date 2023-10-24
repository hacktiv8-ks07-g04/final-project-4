package config

import (
	"os"
)

type Config struct {
	Database Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Get() *Config {
	return &Config{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
