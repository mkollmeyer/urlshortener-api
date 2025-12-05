package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mkollmeyer/urlshortener-api/app"
)

func main() {
	app := app.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start: ", err)
	}
}
