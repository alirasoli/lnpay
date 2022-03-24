package service

import "context"

type Payment interface {
	InitPayment(ctx context.Context, amount int64, description string, webhook string) (hash string, invoice string, err error)
	StartWorker(ctx context.Context) error
}
