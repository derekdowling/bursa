package logger

// Provides us a layer of abstraction so we can easily swap out logging systems in
// future when we want to move to something like elastic search

import (
	"log"
)

func Println(error string) {
	log.Println(error)
}
