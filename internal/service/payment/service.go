package payment

import (
	"context"
	"github.com/pkg/errors"
	"lnpay/internal/config"
	"lnpay/internal/service"
	"lnpay/pkg/lnbits"
)

type payment struct {
	cfg *config.Payment
	ln  *lnbits.LNbits
}

func New(cfg *config.Payment) service.Payment {
	return &payment{
		cfg: cfg,
		ln:  lnbits.New(cfg.LNbits.URL, cfg.LNbits.InvoiceKey),
	}
}

func (s *payment) InitPayment(ctx context.Context, amount int64, currency string, description string) (hash string, invoice string, err error) {
	hash, invoice, err = s.ln.CreateInvoice(ctx, amount, description)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize payment")
		return
	}
	// TODO insert to database and do the logic

	return
}
