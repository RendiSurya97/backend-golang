package config

var (
	cfg Config
)

type (
	PortConfig struct {
		HTTP string `gcfg:"http"`
		GRPC string `gcfg:"grpc"`
	}

	Config struct {
		Port PortConfig
	}
)

// Initialize config
func InitConfig() error {
	cfg = Config{
		Port: PortConfig{
			HTTP: "8080",
			GRPC: "8081",
		},
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
