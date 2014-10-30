package config

import (
	"fmt"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

func init() {
	LoadConfig()
}

// Loads our app configuration files into place
func LoadConfig() {

	// Some magic to get the abs path of the file
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), "../../../config")
	viper.AddConfigPath(filepath)

	// looking for base.yml
	viper.SetConfigName("base")
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
