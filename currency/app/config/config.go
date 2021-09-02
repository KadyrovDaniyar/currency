package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port string
	MsSql Database
	CurrencyURL string
	Swagger string
}

type Database struct {
	Server string
	User string
	Password string
	Port int
	Db string
}

var (
	config Config
)

func Get() *Config {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := new(Config)
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
	return config
}