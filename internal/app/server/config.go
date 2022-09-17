package server

import (
	"server/internal/app/casher"
	"server/internal/app/collect"
	"server/internal/app/systemsproject"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Collect  *collect.Config
	Systems  *systemsproject.Config
	Casher   *casher.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080", //default param
		LogLevel: "info",  //default param
		Collect:  collect.NewConfig(),
		Systems:  systemsproject.NewConfig(),
		Casher:   casher.NewConfig(),
	}
}
