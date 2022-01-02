package config

import (
	"log"

	"github.com/pkg/errors"
	"gopkg.in/gcfg.v1"
)

var (
	cfg Config
)

type (
	Config struct {
		Port      PortConfig
		URL       URLConfig
		SecretKey map[string]*SecretKeyConf
	}

	PortConfig struct {
		HTTP string `gcfg:"http"`
		GRPC string `gcfg:"grpc"`
	}

	//url for other services
	URLConfig struct {
		OmdbURL string
	}

	SecretKeyConf struct {
		Value string
	}
)

// Initialize config
func InitConfig() error {
	err := gcfg.ReadFileInto(&cfg, "configs/etc/searchapp.development.ini")
	if err != nil {
		log.Printf("Error read from configs/etc folder, err: %v\n", err)
		return errors.Wrap(err, "failed to read config")
	}
	return nil
}

// Get curent config value
func Get() Config {
	return cfg
}

// Set new config replacing current config
func Set(config Config) {
	cfg = config
}
