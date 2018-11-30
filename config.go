package config

import (
	"flag"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var confPath string

func init() {
	defaultConfPath := "config/config.yml"

	flag.StringVar(&confPath, "c", defaultConfPath, "Configuration file path.")
	flag.StringVar(&confPath, "config", defaultConfPath, "Configuration file path.")
}

func Conf(conf interface{}) error {
	if !flag.Parsed() {
		flag.Parse()
	}

	configFile, err := ioutil.ReadFile(confPath)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(configFile, conf)
}
