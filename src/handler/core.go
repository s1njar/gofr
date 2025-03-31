package handler

import (
	"gofr.dev/pkg/gofr"
)

func Init(app *gofr.App) {
	app.GET("/", Health)
}
