package model

type Payment struct {
	Hash    string `json:"hash"`
	Invoice string `json:"invoice"`
	Amount  int64  `json:"amount"`
	Webhook string `json:"webhook"`
	Paid    bool   `json:"paid"`
}
