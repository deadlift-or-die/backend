package config

// Config represents the config for the environment
type Config struct {
	Host string
	Port int
}

// NewConfig returns the config for the environment
func NewConfig() (*Config, error) {
	return &Config{
		"localhost",
		8080,
	}, nil
}
