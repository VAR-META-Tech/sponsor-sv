package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type EnvConfig struct {
	HomeKeybase string `yaml:"homekeybase:`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}

func GetEnv(pathToConfig string) (EnvConfig, error) {
	var cfg EnvConfig
	f, err := os.Open(pathToConfig)
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
