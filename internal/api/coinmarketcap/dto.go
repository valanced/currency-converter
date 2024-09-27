package coinmarketcap

import (
	"encoding/json"
	"time"
)

type response struct {
	Data   json.RawMessage `json:"data"`
	Status responseStatus  `json:"status"`
}

type priceConversionResponse map[string]currencyData

type currencyData struct {
	Amount      float64   `json:"amount"`
	ID          string    `json:"id"`
	LastUpdated time.Time `json:"last_updated"`
	Name        string    `json:"name"`
	Quote       quote     `json:"quote"`
	Symbol      string    `json:"symbol"`
}

type quote map[string]currencyQuote

type currencyQuote struct {
	LastUpdated time.Time `json:"last_updated"`
	Price       float64   `json:"price"`
}

type responseStatus struct {
	CreditCount  int       `json:"credit_count"`
	Elapsed      int       `json:"elapsed"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage *string   `json:"error_message"`
	Notice       *string   `json:"notice"`
	Timestamp    time.Time `json:"timestamp"`
	Version      string    `json:"version"`
}
