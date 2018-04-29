package main

import (
	"log"

	"github.com/wung-s/firebase_authentication/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
