package chaos

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Providers map[string]ProviderConfig `yaml:"providers"`
}

type ProviderConfig struct {
	Enabled bool              `yaml:"enabled"`
	Actions []string          `yaml:"actions"`
	Options map[string]string `yaml:",inline"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
