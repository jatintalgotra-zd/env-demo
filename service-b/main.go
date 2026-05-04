package main

import (
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/hello", func(ctx *gofr.Context) (any, error) {
		return map[string]string{
			"service": "service-b",
			"message": "hello from service-b",
		}, nil
	})

	app.Run()
}
