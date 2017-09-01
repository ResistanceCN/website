package service

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config struct {
	Debug               bool
	Listen              string

	PostgresHost        string  `yaml:"postgres_host"`
	PostgresDatabase    string  `yaml:"postgres_database"`
	PostgresUser        string  `yaml:"postgres_user"`
	PostgresPassword    string  `yaml:"postgres_password"`
	PostgresSSL         bool    `yaml:"postgres_ssl"`

	RedisAddr           string  `yaml:"redis_addr"`
	RedisPassword       string  `yaml:"redis_password"`
	RedisDatabase       string  `yaml:"redis_database"`
	RedisPrefix         string  `yaml:"redis_prefix"`

	GoogleClientID      string  `yaml:"google_client_id"`
	GoogleClientSecret  string  `yaml:"google_client_secret"`
	GoogleMapKey        string  `yaml:"google_map_key"`
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
