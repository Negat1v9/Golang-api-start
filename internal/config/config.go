package config

type Config struct {
	ServerPort string
}

// config initialization.
// Edit: get the config when the application starts their environment variables or .env, config.yaml file ...
func NewConfig() *Config {
	return &Config{
		ServerPort: ":8088",
	}
}
