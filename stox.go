package stoxapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/t-hale/stox/errors"
	log "github.com/t-hale/stox/gen/log"
	"github.com/t-hale/stox/utils"
	"time"

	stox "github.com/t-hale/stox/gen/stox"
	"os"
	"strings"
)

const (
	ALPACA_PAPER_API = "https://paper-api.alpaca.markets"
	ALPACA_LIVE_API  = "https://api.alpaca.markets"
)

const (
	MONTHLY   = "monthly"
	QUARTERLY = "quarterly"
	YEARLY    = "yearly"
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
	schedule, err := s.calculateVestingPlan(p, trade)
	if err != nil {
		s.logger.Error().Err(err)
		return &stox.VestingPlanResponse{}, err
	}

	s.logger.Debug().Interface("trade", trade).Msgf("")
	s.logger.Debug().Interface("schedule", schedule).Msgf("")
	return schedule, nil
}

func (s *stoxsrvc) calculateVestingPlan(p *stox.VestingPlanRequest, trade *marketdata.Trade) (*stox.VestingPlanResponse, error) {
	var vestEvents []*stox.VestEvent

	curDate, _ := time.Parse(time.DateOnly, string(p.GrantDate))
	compareDate, _ := time.Parse(time.DateOnly, string(p.VestDate))

	if curDate.After(compareDate) {
		return nil, fmt.Errorf("%w : GrantDate must occur before VestDate", errors.ErrInvalidInput)
	}

	for curDate.Before(compareDate) {

		vestEvent := &stox.VestEvent{
			Date:           utils.PtrTo(stox.Date(curDate.Format(time.DateOnly))),
			UnitsGranted:   utils.PtrTo(int64(0)),
			UnitsRemaining: &p.UnitsGranted,
		}

		vestEvents = append(vestEvents, vestEvent)

		switch p.VestFrequency {
		case MONTHLY:
			curDate = curDate.AddDate(0, 1, 0)
		case QUARTERLY:
			curDate = curDate.AddDate(0, 3, 0)
		case YEARLY:
			curDate = curDate.AddDate(1, 0, 0)
		default:
			return nil, fmt.Errorf("%w : unknown VestFrequency %q encountered while calculating vesting dates", errors.ErrInvalidInput, p.VestFrequency)
		}

	}

	if curDate.After(compareDate) {
		return nil, fmt.Errorf("%w : GrantDate, VestDate and VestFrequency do not align", errors.ErrInvalidInput)
	}

	if curDate.Equal(compareDate) {
		vestEvents = append(vestEvents, &stox.VestEvent{Date: utils.PtrTo(stox.Date(curDate.Format(time.DateOnly)))})
	}

	numEvents := len(vestEvents)
	unitsGrantedPerEvent := p.UnitsGranted / (int64(numEvents) - 1)
	totalUnitsGranted := int64(0)
	unitsRemaining := p.UnitsGranted

	for i := 0; i < numEvents; i++ {

		vestEvents[i].UnitsRemaining = utils.PtrTo(unitsRemaining)
		vestEvents[i].TotalUnitsGranted = utils.PtrTo(totalUnitsGranted)
		vestEvents[i].AmountGranted = utils.PtrTo(0.0)
		vestEvents[i].TotalAmountGranted = utils.PtrTo(0.0)

		if i == 0 {
			continue
		}

		if i == numEvents-1 {
			unitsGrantedPerEvent = unitsRemaining
		}

		totalUnitsGranted += unitsGrantedPerEvent
		unitsRemaining -= unitsGrantedPerEvent

		vestEvents[i].UnitsGranted = utils.PtrTo(unitsGrantedPerEvent)
		vestEvents[i].UnitsRemaining = utils.PtrTo(unitsRemaining)
		vestEvents[i].AmountGranted = utils.PtrTo(float64(unitsGrantedPerEvent) * trade.Price)
		vestEvents[i].TotalAmountGranted = utils.PtrTo(float64(totalUnitsGranted) * trade.Price)
		vestEvents[i].TotalUnitsGranted = utils.PtrTo(totalUnitsGranted)
	}

	return &stox.VestingPlanResponse{
		Symbol:   &p.Symbol,
		Price:    &trade.Price,
		VestPlan: vestEvents,
	}, nil
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
