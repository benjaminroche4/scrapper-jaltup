package main

import "log"

var (
	version = "unknown"
)

func main() {
	if err := parseCLI(); err != nil {
		log.Fatal(err)
	}
}
