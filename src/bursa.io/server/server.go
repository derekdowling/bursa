package main

import (
	"bursa.io/kernel"
	"flag"
)

var prod_flag bool

// Sets our server flags
func init() {
	// listens for a production flag, assume dev mode by default
	flag.Bool(prod_flag, "Production", false, "Starts the server in production mode")
}

func main() {
	// Parse in flags
	flag.Parse()

	// starts the kernel in production mode
	kernel.start(*prod_flag)
}
