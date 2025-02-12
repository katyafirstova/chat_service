package main

import (
	"context"
	"log"

	"github.com/katyafirstova/chat_service/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("Failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Failed to start app: %s", err.Error())
	}
}
