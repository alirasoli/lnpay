package main

import (
	"context"
	"lnpay/internal/transport/http/fiber"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	f := fiber.New()
	err := f.Serve(ctx, ":3000")
	if err != nil {
		cancel()
		log.Fatal(err)
	}
}
