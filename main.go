package main

import "log"

func main() {
	displayVersion()

	err := action(false)
	if err != nil {
		log.Fatal(err)
	}
}
