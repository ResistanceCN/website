package service

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config struct {
	Debug   bool
	Listen  string

	Postgres struct {
		Host      string  `yaml:"host"`
		Database  string  `yaml:"database"`
		User      string  `yaml:"user"`
		Password  string  `yaml:"password"`
		SSL       bool    `yaml:"ssl"`
	} `yaml:"postgres"`

	Redis struct {
		Addr      string  `yaml:"addr"`
		Password  string  `yaml:"password"`
		Database  string  `yaml:"database"`
		Prefix    string  `yaml:"prefix"`
	} `yaml:"redis"`

	Google struct {
		ClientID      string  `yaml:"client_id"`
		ClientSecret  string  `yaml:"client_secret"`
		MapKey        string  `yaml:"map_key"`
	} `yaml:"google"`
}

func init() {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	f.Close()

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
}
