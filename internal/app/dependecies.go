package app

import (
	"context"

	"github.com/shopspring/decimal"
)

type CurrencyConverter interface {
	Convert(_ context.Context, amount decimal.Decimal, from, to string) (decimal.Decimal, error)
}
