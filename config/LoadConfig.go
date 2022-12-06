package config

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	Port  uint16 `json:"port"` //端口
	Mysql struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
		Port     uint16 `json:"port"`
	} `json:"mysql"`
	CollectionAddress string `json:"collection_address"`
}

var Config = config{}

func LoadConfig() error {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		return err
	}

	return nil
}
