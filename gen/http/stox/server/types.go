// Code generated by goa v3.11.3, DO NOT EDIT.
//
// stox HTTP server types
//
// Command:
// $ goa gen github.com/t-hale/stox/design

package server

import (
	"unicode/utf8"

	stox "github.com/t-hale/stox/gen/stox"
	goa "goa.design/goa/v3/pkg"
)

// PlanRequestBody is the type of the "stox" service "plan" endpoint HTTP
// request body.
type PlanRequestBody struct {
	// stock symbol to retrieve plan for
	Symbol *string `form:"Symbol,omitempty" json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	// number of stock units granted
	UnitsGranted *int64 `form:"UnitsGranted,omitempty" json:"UnitsGranted,omitempty" xml:"UnitsGranted,omitempty"`
	// initial grant date of equities
	GrantDate *string `form:"GrantDate,omitempty" json:"GrantDate,omitempty" xml:"GrantDate,omitempty"`
	// date the equities vest completely
	VestDate *string `form:"VestDate,omitempty" json:"VestDate,omitempty" xml:"VestDate,omitempty"`
	// frequency of vesting schedule (monthly, quarterly, yearly)
	VestFrequency *string `form:"VestFrequency,omitempty" json:"VestFrequency,omitempty" xml:"VestFrequency,omitempty"`
}

// PlanResponseBody is the type of the "stox" service "plan" endpoint HTTP
// response body.
type PlanResponseBody struct {
	Symbol   *string                  `form:"Symbol,omitempty" json:"Symbol,omitempty" xml:"Symbol,omitempty"`
	Price    *float64                 `form:"Price,omitempty" json:"Price,omitempty" xml:"Price,omitempty"`
	VestPlan []*VestEventResponseBody `form:"VestPlan,omitempty" json:"VestPlan,omitempty" xml:"VestPlan,omitempty"`
}

// VestEventResponseBody is used to define fields on response body types.
type VestEventResponseBody struct {
	UnitsGranted       *int64   `form:"UnitsGranted,omitempty" json:"UnitsGranted,omitempty" xml:"UnitsGranted,omitempty"`
	UnitsRemaining     *int64   `form:"UnitsRemaining,omitempty" json:"UnitsRemaining,omitempty" xml:"UnitsRemaining,omitempty"`
	TotalUnitsGranted  *int64   `form:"TotalUnitsGranted,omitempty" json:"TotalUnitsGranted,omitempty" xml:"TotalUnitsGranted,omitempty"`
	Date               *string  `form:"Date,omitempty" json:"Date,omitempty" xml:"Date,omitempty"`
	AmountGranted      *float64 `form:"AmountGranted,omitempty" json:"AmountGranted,omitempty" xml:"AmountGranted,omitempty"`
	TotalAmountGranted *float64 `form:"TotalAmountGranted,omitempty" json:"TotalAmountGranted,omitempty" xml:"TotalAmountGranted,omitempty"`
}

// NewPlanResponseBody builds the HTTP response body from the result of the
// "plan" endpoint of the "stox" service.
func NewPlanResponseBody(res *stox.VestingPlanResponse) *PlanResponseBody {
	body := &PlanResponseBody{
		Symbol: res.Symbol,
		Price:  res.Price,
	}
	if res.VestPlan != nil {
		body.VestPlan = make([]*VestEventResponseBody, len(res.VestPlan))
		for i, val := range res.VestPlan {
			body.VestPlan[i] = marshalStoxVestEventToVestEventResponseBody(val)
		}
	}
	return body
}

// NewPlanVestingPlanRequest builds a stox service plan endpoint payload.
func NewPlanVestingPlanRequest(body *PlanRequestBody) *stox.VestingPlanRequest {
	v := &stox.VestingPlanRequest{
		Symbol:        *body.Symbol,
		UnitsGranted:  *body.UnitsGranted,
		GrantDate:     stox.Date(*body.GrantDate),
		VestDate:      stox.Date(*body.VestDate),
		VestFrequency: stox.VestFrequency(*body.VestFrequency),
	}

	return v
}

// ValidatePlanRequestBody runs the validations defined on PlanRequestBody
func ValidatePlanRequestBody(body *PlanRequestBody) (err error) {
	if body.Symbol == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("Symbol", "body"))
	}
	if body.UnitsGranted == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("UnitsGranted", "body"))
	}
	if body.GrantDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("GrantDate", "body"))
	}
	if body.VestDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("VestDate", "body"))
	}
	if body.VestFrequency == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("VestFrequency", "body"))
	}
	if body.Symbol != nil {
		if utf8.RuneCountInString(*body.Symbol) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.Symbol", *body.Symbol, utf8.RuneCountInString(*body.Symbol), 1, true))
		}
	}
	if body.UnitsGranted != nil {
		if *body.UnitsGranted < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.UnitsGranted", *body.UnitsGranted, 1, true))
		}
	}
	if body.GrantDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.GrantDate", *body.GrantDate, goa.FormatDate))
	}
	if body.VestDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.VestDate", *body.VestDate, goa.FormatDate))
	}
	if body.VestFrequency != nil {
		if !(*body.VestFrequency == "monthly" || *body.VestFrequency == "quarterly" || *body.VestFrequency == "yearly") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.VestFrequency", *body.VestFrequency, []any{"monthly", "quarterly", "yearly"}))
		}
	}
	return
}