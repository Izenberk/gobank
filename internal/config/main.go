package config

type Config struct {
	Host 			string
	Port 			int
}

type Option func(*Config)

func WithHost (host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort (port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func NewConfig (opts ...Option) *Config {

	cfg := &Config{
		Host:			"localhost",
		Port: 		8080,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}