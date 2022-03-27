package exchange

import (
	"context"
	"github.com/pkg/errors"
	"lnpay/internal/entity/constants"
	"lnpay/internal/service"
	"lnpay/pkg/nobitex"
)

type exchange struct {
}

func New() service.Exchange {
	return &exchange{}
}

func (s *exchange) GetBtcPrice(ctx context.Context, currency constants.Currency) (float64, error) {
	nobi := nobitex.New()
	var c nobitex.Currency
	if currency == constants.IRT {
		c = nobitex.IRR
	}
	price, err := nobi.GetBtcPrice(ctx, c)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get btc price")
	}

	if currency == constants.IRT {
		price = price / 10
	}

	return price, nil
}

func (s *exchange) GetPriceInSats(ctx context.Context, price float64, currency constants.Currency) (int64, error) {
	btcPrice, err := s.GetBtcPrice(ctx, currency)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get btc price")
	}

	return int64(price / (btcPrice / 100_000_000)), nil
}
