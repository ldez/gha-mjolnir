package main

import "log"

func main() {
	displayVersion()

	err := action(true)
	if err != nil {
		log.Fatal(err)
	}
}
