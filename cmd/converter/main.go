package main

import (
	"context"
	"fmt"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/valanced/currency-converter/internal/api/coinmarketcap"
	"github.com/valanced/currency-converter/internal/app"
	"github.com/valanced/currency-converter/internal/converter"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Currency converter using CoinMarketCap API",
	Long:  `A CLI tool to convert one currency to another using CoinMarketCap API as a data source.`,
}

var convertCmd = &cobra.Command{
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

	ctx := ctxzap.ToContext(context.Background(), logger)
	rootCmd.SetContext(ctx)

	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: COINMARKETCAP_API_KEY is not set")
		os.Exit(1)
	}

	client := coinmarketcap.NewClient(apiKey)
	c := converter.New(client)
	a := app.NewApp(c)

	convertCmd.RunE = func(cmd *cobra.Command, args []string) error {
		result, err := a.HandleConvert(cmd.Context(), args)
		if err != nil {
			return err
		}

		cmd.Println(result)

		return nil
	}

	rootCmd.AddCommand(convertCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
