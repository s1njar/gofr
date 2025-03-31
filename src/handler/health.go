package handler

import "gofr.dev/pkg/gofr"

func Health(ctx *gofr.Context) (any, error) {
	return "Hello World!", nil
}
