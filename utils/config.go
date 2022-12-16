package utils

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigStruct struct {
	Db struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	}
}

func GetDataBaseUrl() string {
	mydir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := ioutil.ReadFile(mydir + "/config.yaml")

	if err != nil {
		panic(err)
	}

	var cfg ConfigStruct

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}
	db_url := "postgresql://" + cfg.Db.Host + ":" + cfg.Db.Port + "/" + cfg.Db.Name + "?user=" + cfg.Db.User + "&password=" + cfg.Db.Pass
	return db_url
}
