package client

import (
	"github.com/conformal/btcrpcclient"
	_ "github.com/derekdowling/bursa/config"
	// TODO how to ensure our config is loaded?
	"github.com/spf13/viper"
)

// TODO g_ prefix? What is the proper way of initalizing global variables in go?
var g_client *btcrpcclient.Client

// Return a connected client
func Get() *btcrpcclient.Client {
	// NOTE I was really hoping to have bursa.io/config loaded for it's side
	// effects via the underscore modifier, but it seems to not read the
	// configuration data when done that way.
	if g_client == nil {
		bitcoin_config := viper.GetStringMap("bitcoin")
		g_client, _ = btcrpcclient.New(&btcrpcclient.ConnConfig{
			// There appears to be a bug with nested booleans and viper.GetBool.
			// HttpPostMode: viper.GetBool("bitcoin.HttpPostMode"),
			HttpPostMode: bitcoin_config["HttpPostMode"].(bool),
			DisableTLS:   bitcoin_config["DisableTLS"].(bool),
			Host:         bitcoin_config["Host"].(string),
			User:         bitcoin_config["User"].(string),
			Pass:         bitcoin_config["Pass"].(string),
		}, nil)
	}
	return g_client
}
