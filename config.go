package matrixbot

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server   string
	User     string
	Password string
	RoomID   string
}

func ReadConfig(confFile string) Config {
	var config Config
	if _, err := toml.DecodeFile(confFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
