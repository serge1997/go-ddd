package main

import (
	"log"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/serge1997/go-ddd/internal/app"
)

func main() {
	app := app.New()
	log.Fatal(app.Start())
}
