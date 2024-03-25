package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type ConnectionConfig struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var once sync.Once

func GetConfig() *ConnectionConfig {
	var instance *ConnectionConfig
	once.Do(func() {
		instance = &ConnectionConfig{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			fmt.Println(cleanenv.GetDescription(instance, nil))
		}
	})

	return instance
}
