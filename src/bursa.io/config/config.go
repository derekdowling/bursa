package config

import (
	"github.com/spf13/viper"
)

// Loads our app configuration files into place
func LoadConfig() {
	viper.SetConfigName("base")
	viper.AddConfigPath("/bursa/src/bursa.io/config/yml/")
	viper.ReadInConfig()
}

// Little wrapper so we don't have to load viper & the config in places, also
// decouples us slightly
func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}
