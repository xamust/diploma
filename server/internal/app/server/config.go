package server

import (
	"server/internal/app/collect"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Collect  *collect.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080", //default param
		LogLevel: "info",  //default param
		Collect:  collect.NewConfig(),
	}
}
