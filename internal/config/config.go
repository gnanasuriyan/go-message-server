package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type IConfig interface {
	GetConfig() *Config
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Config struct {
	Port     uint     `default:"3000" yaml:"port"`
	DBConfig DBConfig `yaml:"database"`
}

var cfg Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		// Load config from file
		yamlFile, err := os.ReadFile("config.yml")
		if err != nil {
			panic(err)
		}
		// Unmarshal the yaml file into the Config struct
		err = yaml.Unmarshal(yamlFile, &cfg)
		if err != nil {
			panic(err)
		}
	})
	return &cfg
}

func (c *Config) GetPort() uint {
	return c.Port
}

func (c *Config) GetDatabaseConfig() DBConfig {
	return c.DBConfig
}
