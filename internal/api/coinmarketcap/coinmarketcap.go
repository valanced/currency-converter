package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/valanced/currency-converter/internal/api"
	"github.com/valanced/currency-converter/internal/util"
	"go.uber.org/zap"
)

const (
	baseURL            = "https://sandbox-api.coinmarketcap.com/v1" // todo: test task hint: might be configured
	priceConversionURL = "/tools/price-conversion"
)

type Client struct {
	apiKey string

	client *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		client: new(http.Client),
	}
}

func (c *Client) FetchRate(ctx context.Context, from, to string) (float64, error) {
	params := url.Values{}
	params.Set("amount", "1") // todo: test task hint: we can use it, but use 1 in purpose to separate service and client logic
	params.Set("symbol", from)
	params.Set("convert", to)
	url := fmt.Sprint(baseURL + priceConversionURL + "?" + params.Encode())

	result := priceConversionResponse{}
	if err := c.doRequest(ctx, url, &result); err != nil {
		return 0, err
	}

	return result[from].Quote[to].Price, nil
}

func (c *Client) doRequest(ctx context.Context, url string, rsp any) error {
	log := ctxzap.Extract(ctx).Named("request coinmarketcap").With(zap.String("url", url))

	log.Debug("preparing")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("%w: http.NewRequest", err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("X-CMC_PRO_API_KEY", c.apiKey)

	resp, err := c.client.Do(req)
	log.With(zap.Any("resp", resp), zap.Error(err)).Debug("client.Do")
	if err != nil {
		return fmt.Errorf("%w: client.Do", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &api.Error{Message: "resp.StatusCode != http.StatusOK", Details: err.Error()}
	}

	var result response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("%w: json.NewDecoder.Decode", err)
	}

	if result.Status.ErrorCode != 0 {
		return &api.Error{
			Code:    result.Status.ErrorCode,
			Message: "coinmarketcap api error",
			Details: util.Deref(result.Status.ErrorMessage),
		}
	}

	if rsp == nil {
		return nil
	}

	if err := json.Unmarshal(result.Data, &rsp); err != nil {
		return fmt.Errorf("%w: json.Unmarshal", err)
	}

	return nil
}
