package main

import (
	"fiber/src/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	handler.Init(app)

	app.Run()
}
