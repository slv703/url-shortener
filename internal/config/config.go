package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string  `env-required:"true" yaml:"env"`
	Server *Server `env-required:"true" yaml:"server"`
}

type Server struct {
	Host    string        `env-required:"true" yaml:"host"`
	Port    string        `env-required:"true" yaml:"port"`
	Timeout time.Duration `env-default:"10s"   yaml:"timeout"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return nil, fmt.Errorf("read config error: %w", err)
	}
	return &config, nil
}
