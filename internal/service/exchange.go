package service

import (
	"context"
	"lnpay/internal/entity/constants"
)

type Exchange interface {
	GetBtcPrice(ctx context.Context, currency constants.Currency) (float64, error)
	GetPriceInSats(ctx context.Context, price float64, currency constants.Currency) (int64, error)
}
