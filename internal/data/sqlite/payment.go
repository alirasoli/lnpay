package sqlite

import (
	"context"
	"lnpay/internal/entity/model"
	"time"
)

func (s *Sqlite) CreatePayment(ctx context.Context, p *model.Payment) error {
	_, err := s.DB.ExecContext(ctx, `
		INSERT INTO payment (
		    hash,
		    invoice,
			amount,
		    webhook
		) VALUES (?, ?, ?, ?)
	`, p.Hash, p.Invoice, p.Amount, p.Webhook)
	return err
}

func (s *Sqlite) GetActivePayments(ctx context.Context, exp time.Duration) ([]*model.Payment, error) {
	rows, err := s.DB.QueryContext(ctx, `
		SELECT
		    hash,
		    invoice,
			amount,
		    webhook
		FROM payment
		WHERE
			paid_at IS NULL
			AND
		    created_at > ?
	`, time.Now().Add(-exp).Unix())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*model.Payment
	for rows.Next() {
		var p model.Payment
		if err := rows.Scan(&p.Hash, &p.Invoice, &p.Amount, &p.Webhook); err != nil {
			return nil, err
		}
		payments = append(payments, &p)
	}
	return payments, nil
}

func (s *Sqlite) SetPaymentPaid(ctx context.Context, hash string) error {
	_, err := s.DB.ExecContext(ctx, `
		UPDATE payment
		SET paid_at = (strftime('%s', 'now'))
		WHERE hash = ?
	`, hash)
	return err
}
