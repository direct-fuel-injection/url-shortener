package config

import "os"

type Config struct {
	Port string
}

func Read() Config {
	var config Config

	port, exists := os.LookupEnv("PORT")

	if exists {
		config.Port = port
	}

	return config
}