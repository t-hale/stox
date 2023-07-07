package stoxapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	log "stox/gen/log"
	"stox/utils"
	"time"

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
func (s *stoxsrvc) Plan(ctx context.Context, p *stox.VestingPlanRequest) (*stox.VestingPlanResponse, error) {
	s.logger.Info().Msgf("stox.plan called with %+v", p)
	trade, err := s.GetLatestTrade(p.Symbol)

	if err != nil {
		s.logger.Error().Err(err)
		return &stox.VestingPlanResponse{}, err
	}
	schedule, err := s.calculateVestingSchedule(p)
	if err != nil {
		return &stox.VestingPlanResponse{}, err
	}

	s.logger.Debug().Interface("trade", trade).Msgf("")
	s.logger.Debug().Interface("schedule", schedule).Msgf("")
	return schedule, nil
}

func (s *stoxsrvc) calculateVestDates(startDate, endDate stox.Date, frequency stox.VestFrequency) (*stox.VestingPlanResponse, error) {
	var vestEvents []*stox.VestEvent

	curDate, _ := time.Parse(time.RFC3339, string(startDate))
	compareDate, _ := time.Parse(time.RFC3339, string(endDate))

	for {

		var vestEvent stox.VestEvent

		vestEvent.Date = utils.PtrTo(stox.Date(curDate.String()))

		vestEvents = append(vestEvents, &vestEvent)

		if frequency == "monthly" {
			curDate = curDate.AddDate(0, 1, 0)
		} else if frequency == "quarterly" {
			curDate = curDate.AddDate(0, 3, 0)
		} else if frequency == "yearly" {
			curDate = curDate.AddDate(1, 0, 0)
		} else {
			return &stox.VestingPlanResponse{}, fmt.Errorf("%q invalid vest frequency encountered while calculating vest dates", frequency)
		}

		if curDate.Equal(compareDate) || curDate.After(compareDate) {
			break
		}

	}
	return &stox.VestingPlanResponse{
		VestPlan: vestEvents,
	}, nil
}

func (s *stoxsrvc) calculateVestingSchedule(p *stox.VestingPlanRequest) (*stox.VestingPlanResponse, error) {
	return &stox.VestingPlanResponse{}, nil
}

func (s *stoxsrvc) GetLatestTrade(symbol string) (*marketdata.Trade, error) {
	return s.getLatestTradeAlpaca(symbol)
}

func (s *stoxsrvc) getLatestTradeAlpaca(symbol string) (*marketdata.Trade, error) {
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

func (s *stoxsrvc) getLatestTradePolygon(symbol string) {

	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	params := models.GetLastTradeParams{
		Ticker: strings.ToUpper(symbol),
	}

	res, err := c.GetLastTrade(context.Background(), &params)
	if err != nil {
		s.logger.Error().Err(err)
	}

	prettyResult, _ := json.MarshalIndent(res, "", "    ")

	// do something with the result
	s.logger.Info().Msgf("Result : %s", prettyResult)
}
