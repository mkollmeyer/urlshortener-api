package main

import (
	"context"
	"fmt"

	"github.com/mkollmeyer/urlshortener-api/app"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start: ", err)
	}
}
