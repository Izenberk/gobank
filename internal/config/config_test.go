package config

import (
	"testing"
	"time"
)

func TestNewConfig_defaults(t *testing.T) {
	got := NewConfig()
	
	want := &Config{
		Host:					"localhost",
		Port: 				9090,
		ReadTimeout: 	10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 	120 * time.Second,
	}

		if got.Host != want.Host {
			t.Errorf("NewConfig() default setting: Host is %v; expected %v", got.Host, want.Host)
		}

		if got.Port != want.Port {
			t.Errorf("NewConfig() default setting: Port is %v; expected %v", got.Port, want.Port)
		}

		if got.ReadTimeout != want.ReadTimeout {
			t.Errorf("NewConfig() default setting: ReadTimeout is %v; expected %v", got.ReadTimeout, want.ReadTimeout)
		}

		if got.WriteTimeout != want.WriteTimeout {
			t.Errorf("NewConfig() default setting: WriteTimeout is %v; expected %v", got.WriteTimeout, want.WriteTimeout)
		}

		if got.IdleTimeout != want.IdleTimeout {
			t.Errorf("NewConfig() default setting: IdleTimeout is %v; expected %v", got.IdleTimeout, want.IdleTimeout)
		}
}

func TestNewConfig_withOption(t *testing.T) {
	got := NewConfig(
		WithHost("localhost"),
		WithPort(8081),
		WithReadTimeout(5 * time.Second),
		WithWriteTimeout(5 * time.Second),
		WithIdleTimeout(150 * time.Second),
	)

	want := &Config{
		Host:					"localhost",
		Port: 				8081,
		ReadTimeout: 	5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 	150 * time.Second,
	}

		if got.Host != want.Host {
			t.Errorf("NewConfig() setting: Host is %v; expected %v", got.Host, want.Host)
		}

		if got.Port != want.Port {
			t.Errorf("NewConfig() setting: Port is %v; expected %v", got.Port, want.Port)
		}

		if got.ReadTimeout != want.ReadTimeout {
			t.Errorf("NewConfig() setting: ReadTimeout is %v; expected %v", got.ReadTimeout, want.ReadTimeout)
		}

		if got.WriteTimeout != want.WriteTimeout {
			t.Errorf("NewConfig() setting: WriteTimeout is %v; expected %v", got.WriteTimeout, want.WriteTimeout)
		}

		if got.IdleTimeout != want.IdleTimeout {
			t.Errorf("NewConfig() setting: IdleTimeout is %v; expected %v", got.IdleTimeout, want.IdleTimeout)
		}
}