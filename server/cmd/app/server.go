package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"server/internal/app/server"
	"sync"
)

var (
	configPath string
	mu         *sync.Mutex
	wg         sync.WaitGroup
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "Path to config file")
}

func main() {

	flag.Parse()
	config := server.NewConfig()
	meta, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	//check undecoded conf...
	if len(meta.Undecoded()) != 0 {
		log.Fatal("Undecoded configs param: ", meta.Undecoded())
	}

	if err := server.New(config).Start(); err != nil {
		log.Fatal(err)
	}
}
