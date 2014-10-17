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
