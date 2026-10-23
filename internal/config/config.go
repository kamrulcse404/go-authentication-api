package config

type Config struct {
	AppName string
	Port    string
}

func Load() Config {
	cfg := Config{
		AppName: "Authentication API",
		Port:    "8080",
	}
	return cfg
}
