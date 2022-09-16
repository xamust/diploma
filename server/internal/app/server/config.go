package server

import (
	"fmt"
	"os"
	"server/internal/app/collect"
	"server/internal/app/systemsProject"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Collect  *collect.Config
	Systems  *systemsProject.Config
}

func NewConfig() *Config {
	PORT := ":8080"
	if os.Getenv("PORT") != "" {
		PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	return &Config{
		BindAddr: PORT,   //default param
		LogLevel: "info", //default param
		Collect:  collect.NewConfig(),
		Systems:  systemsProject.NewConfig(),
	}
}
