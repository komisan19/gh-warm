package main

import (
	"log"
	"os"

	"github.com/komisan19/gh-warm/config"
)

func main() {
	app := config.Execute()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
