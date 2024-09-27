package converter_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valanced/currency-converter/internal/converter"
)

func TestConverter_Convert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRater := NewMockRater(ctrl)

	converter := converter.New(mockRater)

	ctx := context.Background()
	amount := decimal.NewFromInt(10)
	from := "USD"
	to := "EUR"
	expectedRate := 0.84
	expectedResult, err := decimal.NewFromString("8.4")
	require.NoError(t, err)

	mockRater.EXPECT().FetchRate(ctx, from, to).Return(expectedRate, nil)

	result, err := converter.Convert(ctx, amount, from, to)
	assert.NoError(t, err)
	assert.True(t, expectedResult.Equal(result), "Expected: %s, got: %s", expectedResult, result)
}

func TestConverter_Convert_FetchRateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRater := NewMockRater(ctrl)

	converter := converter.New(mockRater)

	ctx := context.Background()
	amount := decimal.NewFromInt(10)
	from := "USD"
	to := "EUR"
	expectedError := errors.New("fetch rate error")

	mockRater.EXPECT().FetchRate(ctx, from, to).Return(0.0, expectedError)

	_, err := converter.Convert(ctx, amount, from, to)
	assert.Error(t, err)                // Проверяем, что вернулась ошибка
	assert.Equal(t, expectedError, err) // Проверяем, что ошибка соответствует ожидаемой
}
