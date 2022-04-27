//go:build integration

package sqlite

import (
	"context"
	"lnpay/internal/config"
	"lnpay/internal/data"
	"lnpay/internal/data/sqlite"
	"lnpay/internal/entity/model"
	"os"
	"testing"
)

func createDatabase() data.Database {
	cfg := config.SQLite{
		Path: "/tmp/lnpay/lnpay.db",
	}
	os.MkdirAll("/tmp/lnpay", os.ModePerm)

	db, err := sqlite.New(cfg)
	if err != nil {
		panic(err)
	}

	return db
}

func teardown() {
	os.Remove("/tmp/lnpay/lnpay.db")
}

func TestCreatePayment(t *testing.T) {
	db := createDatabase()
	defer teardown()

	payment := &model.Payment{
		Hash:    "hash",
		Invoice: "invoice",
		Amount:  1,
		Webhook: "webhook",
		Paid:    false,
	}

	err := db.CreatePayment(context.Background(), payment)
	if err != nil {
		t.Errorf("Error creating payment: %v", err)
	}
}
