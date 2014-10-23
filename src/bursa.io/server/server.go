package main

import (
	"bursa.io/kernel"
	"flag"
)

var production = flag.Bool("prod", false, "Starts the server in production mode")

func main() {
	// Parse in flags
	flag.Parse()

	// starts the kernel in production mode
	kernel.Start(*production)
}
