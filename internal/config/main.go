package config

import "time"

type Config struct {
	Host 					string
	Port 					int
	ReadTimeout		time.Duration
	WriteTimeout	time.Duration
	IdleTimeout		time.Duration
}

type Option func(*Config)

func WithHost(host string) Option {
	return func(c *Config) {
		c.Host = host
	}
}

func WithPort(port int) Option {
	return func(c *Config) {
		c.Port = port
	}
}

func WithReadTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.WriteTimeout = t
	}
}

func WithIdleTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.IdleTimeout = t
	}
}

func NewConfig(opts ...Option) *Config {

	cfg := &Config{
		Host:					"localhost",
		Port: 				9090,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 	120 * time.Second,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}