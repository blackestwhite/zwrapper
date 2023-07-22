package main

import (
	"log"

	"github.com/blackestwhite/zwrapper/config"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Load()
	log.Println("Hello World")
}
