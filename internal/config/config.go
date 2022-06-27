package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Name          string `json:"name"`
	Port          string `json:"port"`
	LogDir        string `json:"log_dir"`
	LogFile       string `json:"log_file_name"`
	LogLevel      string `json:"log_level"`
	MySqlDSN      string `json:"mysql_dsn"`
	RedisAddr     string `json:"redis_addr"`
	RedisPassword string `json:"redis_password"`
	RedisDB       int    `json:"redis_db"`
}

var dConfig = &Config{
	Name:          "x-url service",
	Port:          ":8080",
	LogDir:        "log",
	LogFile:       "xurl.log",
	LogLevel:      "debug",
	MySqlDSN:      "root:password@(localhost:3306)/service",
	RedisAddr:     "localhost:6379",
	RedisPassword: "",
	RedisDB:       0,
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
