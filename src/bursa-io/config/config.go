package config

import (
	"github.com/spf13/viper"
)

func init() {
	LoadConfig()
}

func LoadConfig() {
	viper.SetConfigName("base")
	viper.AddConfigPath("/bursa/src/bursa-io/config/yml/")
	viper.ReadInConfig()
}
