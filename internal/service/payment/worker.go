package payment

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"time"
)

func (s *payment) StartWorker(ctx context.Context) error {
	// TODO get from config
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

L:
	for {
		select {
		case <-ctx.Done():
			break L
		case <-ticker.C:
			err := s.checkPaymentsStatus(ctx)
			// TODO what to do with error???
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *payment) checkPaymentsStatus(ctx context.Context) error {
	// TODO get from config
	payments, err := s.db.GetActivePayments(ctx, 5*time.Minute)
	if err != nil {
		return errors.Wrap(err, "failed to get active payments")
	}

	for _, payment := range payments {
		paid, err := s.ln.CheckInvoice(ctx, payment.Hash)
		if err != nil {
			return errors.Wrap(err, "failed to check invoice")
		}
		if paid {
			err = s.db.SetPaymentPaid(ctx, payment.Hash)
			if err != nil {
				return errors.Wrap(err, "failed to set payment paid")
			}

			go func() {
				if err := s.callWebhook(ctx, payment.Webhook, payment.Hash); err != nil {
					// TODO properly handle error
					log.Println(err)
				}
			}()
		}
	}

	return nil
}
