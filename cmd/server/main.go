package main

import (
	"context"
	"lnpay/internal/config"
	"lnpay/internal/transport/http/fiber"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Fatal(err)
	}

	f := fiber.New()
	if err := f.Serve(ctx, ":"+cfg.Server.Http.Port); err != nil {
		cancel()
		log.Fatal(err)
	}
}
