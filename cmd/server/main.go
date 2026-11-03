package main

import (
	"log"

	"github.com/serge1997/go-ddd/internal/app"
)

func main() {
	app := app.New()
	log.Fatal(app.Start())
}
