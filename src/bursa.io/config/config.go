package config

import (
	"github.com/spf13/viper"
)

// Loads our app configuration files into place
func LoadConfig() {
	viper.SetConfigName("base")
	viper.AddConfigPath("config/")
	viper.ReadInConfig()
}

// Little wrapper so we don't have to load both viper & the config in places,
// also decouples us slightly

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}
