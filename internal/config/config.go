package config

import "os"

type Config struct {
	AppName     string
	Port        string
	DatabaseURL string
}

func Load() Config {
	cfg := Config{
		AppName:     "Authentication API",
		Port:        "8080",
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
	return cfg
}
