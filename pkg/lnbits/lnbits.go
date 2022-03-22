package lnbits

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type LNbits struct {
	URL        string
	InvoiceKey string
}

func New(url, invoiceKey string) *LNbits {
	return &LNbits{
		URL:        url,
		InvoiceKey: invoiceKey,
	}
}

type invoice struct {
	PaymentHash    string `json:"payment_hash"`
	PaymentRequest string `json:"payment_request"`
}

func (ln *LNbits) CreateInvoice(ctx context.Context, amount int64, memo string) (hash string, request string, err error) {
	body, _ := json.Marshal(map[string]interface{}{
		"amount": amount,
		"memo":   memo,
	})
	u, _ := url.Parse(ln.URL)
	u.Path = path.Join(u.Path, "/api/v1/payments")
	req, err := http.NewRequestWithContext(ctx, "POST", u.String(), bytes.NewBuffer(body))
	if err != nil {
		err = errors.Wrap(err, "failed to create request")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", ln.InvoiceKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "failed to send request")
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read response body")
		return
	}
	if resp.StatusCode != http.StatusCreated {
		err = errors.New("failed to create invoice")
		return
	}
	var inv invoice
	err = json.Unmarshal(body, &inv)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal response")
		return
	}

	return inv.PaymentHash, inv.PaymentRequest, nil
}

type invoiceStatus struct {
	Paid bool `json:"paid"`
}

func (ln *LNbits) CheckInvoice(ctx context.Context, hash string) (paid bool, err error) {
	u, _ := url.Parse(ln.URL)
	u.Path = path.Join(u.Path, "/api/v1/payments/"+hash)
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		err = errors.Wrap(err, "failed to create request")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", ln.InvoiceKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "failed to send request")
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.Wrap(err, "failed to read response body")
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.Errorf("failed to check invoice: %s", string(body))
		return
	}
	var inv invoiceStatus
	err = json.Unmarshal(body, &inv)
	if err != nil {
		err = errors.Wrap(err, "failed to unmarshal response")
		return
	}
	return inv.Paid, nil
}
