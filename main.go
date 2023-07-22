package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hello World")
}
