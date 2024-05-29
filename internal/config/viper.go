package config

import (
	"github.com/spf13/viper"
)

func NewViper() *viper.Viper {
	config := viper.New()

	config.SetEnvPrefix("VP")
	config.AutomaticEnv()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath("./")
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	return config
}
