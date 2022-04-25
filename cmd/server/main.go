package main

import (
	"context"
	"lnpay/internal/config"
	"lnpay/internal/data/sqlite"
	"lnpay/internal/service/exchange"
	"lnpay/internal/service/payment"
	"lnpay/internal/transport/grpc"
	"lnpay/internal/transport/http/fiber"
	"log"
	"os/signal"
	"sync"
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

	exchangeService := exchange.New()
	paymentService := payment.New(&cfg.Payment, sqliteDB)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := paymentService.StartWorker(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		f := fiber.New(paymentService, exchangeService)
		if err := f.Serve(ctx, ":"+cfg.Server.Http.Port); err != nil {
			cancel()
			log.Fatal(err)
		}
	}()

	grpcServer := grpc.New()
	if err := grpcServer.Serve(ctx); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
