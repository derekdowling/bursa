package config

import (
	"github.com/derekdowling/mamba"
	"path"
	"runtime"
	"strings"
)

func init() {
	LoadConfig()
}

var Server *mamba.Config
var DB     *mamba.Config

// Loads our app configuration files into place
func LoadConfig() {

	// Get Load Path
	baseDir := getLoadPath()

	// LOAD APP CONFIG
	Server = LoadServer(baseDir)

	// LOAD DB CONFIG
	DB = LoadDB(baseDir)
}

func getLoadPath() string {
	// Some magic to get the abs path of the file
	_, filename, _, _ := runtime.Caller(1)
	baseDir := strings.Join([]string{path.Dir(filename), "yml"}, "/")
	return baseDir
}

func LoadServer(baseDir string) *mamba.Config {
	server := mamba.NewConfig()
	server.SetConfigName("server")
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
