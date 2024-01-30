package config

import (
	"time"
)

type Config struct {
	RunAddress  string
	DatabaseURI string
	TokenTTL    time.Duration
}

func UseServerStartParams() *Config {
	return &Config{
		RunAddress:  "localhost:8080",
		DatabaseURI: "host='localhost' user='postgres' password='123' dbname='postgres' sslmode=disable",
		TokenTTL:    10 * time.Hour,
	}
}
