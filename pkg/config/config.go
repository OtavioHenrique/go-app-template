package config

type Config struct {
	AppName  string
	HttpPort string
}

// This class should hold all config, related to load secrets, envs and any other thing.
// Should be the first thing to be created to "build" application
func NewConfig(appName string) *Config {
	config := new(Config)

	config.AppName = appName
	config.HttpPort = ":8081"

	return config
}
