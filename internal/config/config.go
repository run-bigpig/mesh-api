package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

var (
	globalConfig *Config
	once         sync.Once
)

type Config struct {
	App     string     `yaml:"App"`
	Listen  string     `yaml:"Listen"`
	Adapter []*Adapter `yaml:"Adapter"`
	Mysql   *Mysql     `yaml:"Mysql"`
}

type Adapter struct {
	Name string `yaml:"Name"`
	Mark string `yaml:"Mark"`
	Host string `yaml:"Host"`
}

type Mysql struct {
	DSN string `yaml:"DSN"`
}

func Set(configPath string) {
	once.Do(func() {
		file, err := os.ReadFile(configPath)
		if err != nil {
			return
		}

		var config Config
		if err := yaml.Unmarshal(file, &config); err != nil {
			return
		}
		globalConfig = &config
	})
}

func Get() *Config {
	return globalConfig
}
