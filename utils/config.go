package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	db struct {
		host string `yaml:"host"`
		port string `yaml:"port"`
		user string `yaml:"user"`
		pass string `yaml:"pass"`
		name string `yaml:"name"`
	}
}

func GetDataBaseUrl() string {
	file, err := ioutil.ReadFile("example2.yaml")

	if err != nil {
		panic(err)
	}

	var cfg ConfigStruct

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}
	db_url := "postgresql://" + cfg.db.host + ":" + cfg.db.port + "/" + cfg.db.name + "?user=" + cfg.db.user + "&password=" + cfg.db.pass
	return db_url
}
