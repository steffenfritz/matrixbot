package matrixbot

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	User     string
	Password string
	RoomID   string
}

// ReadConfig reads the configuration of a toml file into a Config struct
func ReadConfig(confFile string) Config {
	var config Config
	if _, err := toml.DecodeFile(confFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func checkConfig(config Config) bool {
	if len(config.Server) == 0 {
		return false
	}

	return true
}
