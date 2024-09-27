package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"

	"github.com/valanced/currency-converter/internal/api/coinmarketcap"
	"github.com/valanced/currency-converter/internal/app"
	"github.com/valanced/currency-converter/internal/converter"
)

var rootCmd = &cobra.Command{
	Use:   "convert [amount] [from_currency] [to_currency]",
	Short: "Converts currency from one to another",
	Long: `Converts the specified amount from the source currency to the target currency 
using CoinMarketCap API.`,
	Args: cobra.ExactArgs(3),
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("%w: zap.NewProduction", err))
	}
	defer logger.Sync()

	backgroundCtx, cancel := context.WithCancel(context.Background())
	backgroundCtx = ctxzap.ToContext(backgroundCtx, logger)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	rootCmd.SetContext(backgroundCtx)

	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: COINMARKETCAP_API_KEY is not set")
		os.Exit(1)
	}

	client := coinmarketcap.NewClient(apiKey)
	c := converter.New(client)
	a := app.New(c)

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		result, err := a.HandleConvert(cmd.Context(), args[0], args[1], args[2])
		if err != nil {
			return err
		}

		cmd.Println(result)

		return nil
	}

	eg, ctx := errgroup.WithContext(backgroundCtx)

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case sig := <-signalChan:
			ctxzap.Extract(ctx).Error("received signal", zap.Error(err), zap.Any("signal", sig))
			return err
		}
	})

	eg.Go(func() error {
		if err := rootCmd.Execute(); err != nil {
			return err
		}
		cancel()

		return nil
	})

	if err := eg.Wait(); err != nil {
		ctxzap.Extract(backgroundCtx).Error("eg.Wait", zap.Error(err))
	}
}
