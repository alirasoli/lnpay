package main

import (
	"context"
	"lnpay/internal/config"
	"lnpay/internal/data/sqlite"
	"lnpay/internal/service/payment"
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

	sqliteDB, err := sqlite.New(cfg.Database.SQLite)
	if err != nil {
		log.Fatal(err)
	}

	paymentService := payment.New(&cfg.Payment, sqliteDB)

	go func() {
		err := paymentService.StartWorker(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	f := fiber.New(paymentService)
	if err := f.Serve(ctx, ":"+cfg.Server.Http.Port); err != nil {
		cancel()
		log.Fatal(err)
	}
}
