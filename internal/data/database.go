package data

import (
	"context"
	"lnpay/internal/entity/model"
	"time"
)

type Database interface {
	Payment
}

type Payment interface {
	CreatePayment(ctx context.Context, payment *model.Payment) error
	GetActivePayments(ctx context.Context, exp time.Duration) ([]*model.Payment, error)
	SetPaymentPaid(ctx context.Context, hash string) error
}
