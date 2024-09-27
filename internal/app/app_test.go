package app_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/valanced/currency-converter/internal/api"
	"github.com/valanced/currency-converter/internal/app"
)

func TestHandleConvert_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConverter := NewMockCurrencyConverter(ctrl)

	app := app.New(mockConverter)

	ctx := context.Background()
	amountStr := "10.0"
	from := "USD"
	to := "EUR"
	amount, _ := decimal.NewFromString(amountStr)
	expectedResult := decimal.NewFromFloat(8.4)

	mockConverter.EXPECT().Convert(gomock.Any(), amount, from, to).Return(expectedResult, nil)

	result, err := app.HandleConvert(ctx, amountStr, from, to)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult.String(), result)
}

func TestHandleConvert_InvalidAmount(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConverter := NewMockCurrencyConverter(ctrl)

	app := app.New(mockConverter)

	ctx := context.Background()
	invalidAmountStr := "invalid"
	from := "USD"
	to := "EUR"

	result, err := app.HandleConvert(ctx, invalidAmountStr, from, to)
	assert.Error(t, err)
	assert.Equal(t, "", result)
	assert.Equal(t, "can't convert invalid to decimal", err.Error())
}

func TestHandleConvert_ApiError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConverter := NewMockCurrencyConverter(ctrl)

	app := app.New(mockConverter)

	ctx := context.Background()

	amountStr := "10.0"
	from := "USD"
	to := "EUR"
	amount, _ := decimal.NewFromString(amountStr)

	expectedApiError := api.Error{Code: 0, Message: "some error"}
	expectedErrorText := "internal error"

	mockConverter.EXPECT().Convert(gomock.Any(), amount, from, to).Return(decimal.Decimal{}, expectedApiError)

	result, err := app.HandleConvert(ctx, amountStr, from, to)
	assert.Error(t, err)
	assert.Equal(t, expectedErrorText, err.Error())
	assert.Equal(t, "", result)
}
