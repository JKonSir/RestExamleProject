package app

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	Service struct {
		Postgres struct {
			USERNAME string `yaml:"user_name"`
			PASSWORD string `yaml:"password"`
			DB_NAME string `yaml:"db_name"`
		}
	}
}

func (conf *Conf) ReadConfig() {

	data, err := ioutil.ReadFile("src/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if yaml.Unmarshal([]byte(data), &conf) != nil {
		log.Fatal(err)
	}
}