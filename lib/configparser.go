package lib

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Token       string   `json:"bot_token"`
	DatabaseURL string   `json:"db_url"`
	Roles       []string `json:"admin_roles"`
	Name        string   `json:"name"`
	LogoURL     string   `json:"logo_url"`
	Activities  []string `json:"activities"`
}

func LoadConfig() *Config {
	// Read config file
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// json data
	var obj Config

	// Parse config file
	err = json.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}

	return &obj
}
