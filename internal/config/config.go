package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ServerPort int `json:"server_port"`
}

func LoadConfig() *Config {
	var AppConfig Config
	raw, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(raw, &AppConfig); err != nil {
		panic(err)
	}
	return &AppConfig
}
