package service

import "context"

type Payment interface {
	InitPayment(ctx context.Context, amount int64, currency string, description string) (hash string, invoice string, err error)
}
