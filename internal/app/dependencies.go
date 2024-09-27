package app

import (
	"context"

	"github.com/shopspring/decimal"
)

//go:generate mockgen -source=./dependencies.go -destination=./mocks_test.go -package=app_test

type CurrencyConverter interface {
	Convert(_ context.Context, amount decimal.Decimal, from, to string) (decimal.Decimal, error)
}
