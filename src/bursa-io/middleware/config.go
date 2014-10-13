package middleware

import (
	"github.com/spf13/viper"
	"net/http"
)

type ConfigMiddleware struct{}

// Loads our app configuration files into place
func loadConfig() {
	viper.SetConfigName("base")
	viper.AddConfigPath("/bursa/src/bursa-io/config/yml/")
	viper.ReadInConfig()
}

func (c *ConfigMiddleware) GetHandler() Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		loadConfig()
	}
}
