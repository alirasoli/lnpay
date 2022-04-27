//go:build integration

package sqlite

import (
	"context"
	"lnpay/internal/config"
	"lnpay/internal/data/sqlite"
	"lnpay/internal/entity/model"
	"os"
	"testing"
)

func createDatabase() *sqlite.Sqlite {
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
		return
	}

	row := db.DB.QueryRow("SELECT hash, invoice, amount, webhook, paid_at FROM payment")
	var p model.Payment
	var paidAt *string
	if err := row.Scan(&p.Hash, &p.Invoice, &p.Amount, &p.Webhook, &paidAt); err != nil {
		t.Errorf("Error scanning payment: %v", err)
		return
	}

	if p.Hash != payment.Hash {
		t.Errorf("Expected hash %s, got %s", payment.Hash, p.Hash)
	}
	if p.Invoice != payment.Invoice {
		t.Errorf("Expected invoice %s, got %s", payment.Invoice, p.Invoice)
	}
	if p.Amount != payment.Amount {
		t.Errorf("Expected amount %d, got %d", payment.Amount, p.Amount)
	}
	if p.Webhook != payment.Webhook {
		t.Errorf("Expected webhook %s, got %s", payment.Webhook, p.Webhook)
	}
	if paidAt != nil {
		t.Errorf("Expected nil, got %v", paidAt)
	}
}
