package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Name     string
	Port     string
	LogDir   string `json:"log_dir"`
	LogFile  string `json:"log_file_name"`
	LogLevel string `json:"log_level"`
}

var dConfig = &Config{
	Name:     "x-url service",
	Port:     ":8080",
	LogDir:   "log",
	LogFile:  "xurl.log",
	LogLevel: "info",
}

func DefaultConfig() *Config {
	return dConfig
}

func (cfg *Config) LoadFromYAML(path string) error {
	return nil
}

func (cfg *Config) LoadFromJson(path string) error {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileData, cfg)
	return err
}
