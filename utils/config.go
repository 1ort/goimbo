package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Api struct {
		BaseUrl string `yaml:"base_url"`
	}
	Db struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
}

func ReadConfig(configPath string) *Config {
	config_file, err := os.ReadFile(configPath)
	if err != nil {
		panic(fmt.Sprintf("Can not open config file, error: %s", err))
	}
	var cfg Config
	err = yaml.Unmarshal(config_file, &cfg)
	if err != nil {
		panic(fmt.Sprintf("Can not parse config file, error: %s", err))
	}
	return &cfg
}

func (cfg *Config) GetDataBaseUrl() string {
	db_template := "postgresql://%s:%s/%s?user=%s&password=%s"
	db_url := fmt.Sprintf(db_template, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name, cfg.Db.User, cfg.Db.Pass)
	return db_url
}

func (cfg *Config) GetAppAddr() string {
	addr_template := "%s:%s"
	addr := fmt.Sprintf(addr_template, cfg.App.Host, cfg.App.Port)
	return addr
}

func (cfg *Config) GetBaseApiUrl() string {
	return (cfg.Api.BaseUrl)
}
