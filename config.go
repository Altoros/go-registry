package main

import (
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

var defaultConfig = Config{
	Port: 3000,
}

func Default() *Config {
	config := defaultConfig

	return &config
}

func (c *Config) Initialize(configYAML []byte) error {
	return yaml.Unmarshal(configYAML, &c)
}

func (c *Config) PortStr() string {
	return strconv.Itoa(int(c.Port))
}

func InitFromFile(path string) (*Config, error) {
	var config *Config = Default()
	var err error

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = config.Initialize(bytes)
	if err != nil {
		return nil, err
	}

	return config, nil
}
