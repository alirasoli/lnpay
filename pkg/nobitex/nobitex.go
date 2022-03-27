package nobitex

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Nobitex struct {
	URL string
}

type Currency string

const (
	IRR Currency = "rls"
)

func New() *Nobitex {
	return &Nobitex{
		URL: "https://api.nobitex.ir",
	}
}

type statResponse struct {
	Stats struct {
		BTC_RLS struct {
			Latest string `json:"latest"`
		} `json:"btc-rls"`
	} `json:"stats"`
}

func (n *Nobitex) GetBtcPrice(ctx context.Context, currency Currency) (float64, error) {
	u, _ := url.Parse(n.URL)
	u.Path = path.Join(u.Path, "/market/stats")

	resp, err := http.PostForm(u.String(), url.Values{
		"srcCurrency": {"btc"},
		"dstCurrency": {string(currency)},
	})
	if err != nil {
		return 0, errors.Wrap(err, "failed to get btc price")
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read response body")
	}

	var res statResponse
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, errors.Wrap(err, "failed to unmarshal response")
	}

	price, err := strconv.ParseFloat(res.Stats.BTC_RLS.Latest, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse price")
	}
	return price, nil
}
