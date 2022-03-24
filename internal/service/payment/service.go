package payment

import (
	"context"
	"github.com/pkg/errors"
	"lnpay/internal/config"
	"lnpay/internal/data"
	"lnpay/internal/entity/model"
	"lnpay/internal/service"
	"lnpay/pkg/lnbits"
	"net/http"
	"net/url"
)

type payment struct {
	cfg *config.Payment
	ln  *lnbits.LNbits
	db  data.Database
}

func New(cfg *config.Payment, db data.Database) service.Payment {
	return &payment{
		cfg: cfg,
		ln:  lnbits.New(cfg.LNbits.URL, cfg.LNbits.InvoiceKey),
		db:  db,
	}
}

func (s *payment) InitPayment(ctx context.Context, amount int64, description string, webhook string) (hash string, invoice string, err error) {
	hash, invoice, err = s.ln.CreateInvoice(ctx, amount, description)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize payment")
		return
	}

	err = s.db.CreatePayment(ctx, &model.Payment{
		Hash:    hash,
		Invoice: invoice,
		Amount:  amount,
		Webhook: webhook,
	})
	if err != nil {
		err = errors.Wrap(err, "failed to create payment")
		return
	}

	return
}

func (s *payment) callWebhook(ctx context.Context, webhook string, hash string) error {
	u, _ := url.Parse(webhook)
	q := u.Query()
	q.Set("hash", hash)
	u.RawQuery = q.Encode()
	// TODO use context
	_, err := http.Get(u.String())
	if err != nil {
		return errors.Wrap(err, "failed to call webhook")
	}
	return nil
}
