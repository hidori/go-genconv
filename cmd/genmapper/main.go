package main

import (
	"log"

	app "github.com/hidori/go-genmapper/app/genmapper"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
