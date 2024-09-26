package converter

import (
	"context"

	"github.com/shopspring/decimal"
)

type Converter struct {
	rater Rater
}

func New(f Rater) *Converter {
	return &Converter{
		rater: f,
	}
}

func (c *Converter) Convert(ctx context.Context, amount decimal.Decimal, from, to string) (decimal.Decimal, error) {
	rate, err := c.rater.FetchRate(ctx, from, to)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return amount.Mul(decimal.NewFromFloat(rate)), nil
}
