package config

import (
	"flag"
	"time"
)

type Config struct {
	RunAddress  string
	DatabaseURI string
	TokenTTL    time.Duration
}

func UseServerStartParams() *Config {
	var c *Config

	flag.StringVar(&c.RunAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.DatabaseURI, "d", "", "connection string to postgres db")
	flag.DurationVar(&c.TokenTTL, "t", 10*time.Hour, "token TTL")

	return c
}
