package stoxapi

import (
	"context"
	"encoding/json"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"log"
	"os"
	stox "stox/gen/stox"
	"strings"
)

const (
	ALPACA_PAPER_API = "https://paper-api.alpaca.markets"
	ALPACA_LIVE_API  = "https://api.alpaca.markets"
)

// stox service example implementation.
// The example methods log the requests and return zero values.
type stoxsrvc struct {
	logger *log.Logger
}

// NewStox returns the stox service implementation.
func NewStox(logger *log.Logger) stox.Service {
	return &stoxsrvc{logger}
}

// Plan implements plan.
func (s *stoxsrvc) Plan(ctx context.Context, p *stox.PlanPayload) (int, error) {
	s.logger.Printf("stox.plan called with %+v", p)
	trade, err := GetLatestTrade(p.Symbol)

	if err != nil {
		log.Printf("ERROR: %s", err)
		return 1, err
	}

	formatted, _ := json.MarshalIndent(trade, "", "    ")
	log.Printf("%s", formatted)

	return 0, nil
}

func GetLatestTrade(symbol string) (*marketdata.Trade, error) {
	return GetLatestTradeAlpaca(symbol)
}

func GetLatestTradeAlpaca(symbol string) (*marketdata.Trade, error) {
	req := marketdata.GetLatestTradeRequest{
		Feed:     marketdata.IEX,
		Currency: "USD",
	}

	trade, err := marketdata.GetLatestTrade(symbol, req)
	if err != nil {
		return nil, err
	}

	return trade, nil
}

func GetLatestTradePolygon(symbol string) {

	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	params := models.GetLastTradeParams{
		Ticker: strings.ToUpper(symbol),
	}

	res, err := c.GetLastTrade(context.Background(), &params)
	if err != nil {
		log.Println(err)
	}

	prettyResult, _ := json.MarshalIndent(res, "", "    ")

	// do something with the result
	log.Printf("Result : %s", prettyResult)
}
