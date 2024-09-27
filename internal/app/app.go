package app

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/shopspring/decimal"
	"github.com/valanced/currency-converter/internal/api"
	"go.uber.org/zap"
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

func (a *App) HandleConvert(ctx context.Context, amountstr, from, to string) (empty string, _ error) {
	amount, err := decimal.NewFromString(amountstr)
	if err != nil {
		return empty, err
	}

	bctx, _ := context.WithTimeout(ctx, timeout) // todo: test task hint: not really necessary here, just for example

	result, err := a.converter.Convert(bctx, amount, from, to)
	if err != nil {
		var apiErr api.Error
		if errors.As(err, &apiErr) && apiErr.Code == 0 { // todo: test task hint: in our example code > 0 means external service error and internal otherwise
			ctxzap.Extract(ctx).Error("converter.Convert apiErr", zap.Error(err))
			return empty, fmt.Errorf("internal error")
		}
		return empty, err
	}

	return result.String(), nil
}
