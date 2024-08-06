package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Targets []string `yaml:"targets"`
}

func LoadConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
