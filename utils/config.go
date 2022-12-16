package utils

import (
	// "io/ioutil"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
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
	db_url := "postgresql://" + cfg.Db.Host + ":" + cfg.Db.Port + "/" + cfg.Db.Name + "?user=" + cfg.Db.User + "&password=" + cfg.Db.Pass
	return db_url
}
