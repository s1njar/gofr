package main

import "gofr.dev/pkg/gofr"

func main() {
	app := gofr.New()

	app.GET("/", func(ctx *gofr.Context) (any, error) {
		return "Hello World!", nil
	})

	app.Run()
}
