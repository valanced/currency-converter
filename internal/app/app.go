package app

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/valanced/currency-converter/internal/errors"
)

const (
	timeout = 5 * time.Second
)

type App struct {
	converter CurrencyConverter
}

func NewApp(converter CurrencyConverter) *App {
	return &App{converter: converter}
}

func (a *App) HandleConvert(ctx context.Context, args []string) (empty string, _ error) {
	if len(args) != 3 {
		return empty, errors.ErrInvalidArgument
	}

	from := args[1]
	to := args[2]

	amount, err := decimal.NewFromString(args[0])
	if err != nil {
		return empty, errors.ErrInvalidArgument
	}

	bctx, _ := context.WithTimeout(ctx, timeout)

	result, err := a.converter.Convert(bctx, amount, from, to)
	if err != nil {
		return empty, err
	}

	return result.String(), nil
}
