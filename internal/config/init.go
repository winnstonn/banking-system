package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() *Config {
	cfg := &Config{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return cfg
}
