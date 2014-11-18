package config

import (
	"github.com/derekdowling/mamba"
	"os"
	"path"
	"runtime"
	"strings"
)

func init() {
	LoadConfig()
}

// Publicly accessible Mamba Configs
var Server *mamba.Config
var DB *mamba.Config

// So we don't overwrite our existing configs
var loaded = false

// Loads our app configuration files into place
func LoadConfig() {

	if loaded == false {
		// Get Load Path
		baseDir := getLoadPath()

		// LOAD APP CONFIG
		Server = LoadServer(baseDir)

		// LOAD DB CONFIG
		DB = LoadDB(baseDir)

		loaded = true
	}
}

func getLoadPath() string {
	// Some magic to get the abs path of the file
	_, filename, _, _ := runtime.Caller(1)
	baseDir := strings.Join([]string{path.Dir(filename), "yml"}, "/")
	return baseDir
}

// Detect the BURSA_ENV variable and user it to determine the final config name.
// E.g. server.production. We default to server.development if the ENV isn't set.
func getConfigName(baseDir string) string {
	config_parts := []string{
		"server",
	}

	if part := os.Getenv("BURSA_ENV"); part != "" {
		config_parts = append(config_parts, part)
	} else {
		config_parts = append(config_parts, "development")
	}

	return strings.Join(config_parts, ".")
}

func LoadServer(baseDir string) *mamba.Config {
	server := mamba.NewConfig()
	server.SetConfigName(
		getConfigName("server"),
	)
	server.AddConfigPath(baseDir)
	server.ReadInConfig()
	return server
}

func LoadDB(baseDir string) *mamba.Config {
	database := mamba.NewConfig()
	database.SetConfigName("database")
	database.AddConfigPath(baseDir)
	database.ReadInConfig()
	return database
}
