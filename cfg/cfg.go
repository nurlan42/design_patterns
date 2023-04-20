package cfg

import (
	"encoding/json"
	"os"
	"time"
)

type (
	Config struct {
		DB   DB
		HTTP HTTP
	}

	DB struct {
		TODO TODO
	}

	TODO struct {
		Host     string
		Port     int
		UserName string
		Password string
		Name     string
		SSLMode  string
		Type     string
	}

	HTTP struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		GinDebugMode bool
	}
)

func NewConfig(cfgPath string) (*Config, error) {
	cfg := new(Config)

	if err := cfg.Parse(cfgPath); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Parse(cfgPath string) error {

	f, err := os.ReadFile(cfgPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(f, c); err != nil {
		return err
	}
	return nil
}
