// Code generated by goa v3.11.3, DO NOT EDIT.
//
// stox service
//
// Command:
// $ goa gen github.com/t-hale/stox/design

package stox

import (
	"context"
)

// The stox service provides advisors with a comprehensive view of a particular
// stock schedule.
type Service interface {
	// Plan implements plan.
	Plan(context.Context, *VestingPlanRequest) (res *VestingPlanResponse, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "stox"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"plan"}

type Date string

type VestEvent struct {
	UnitsGranted       *int64
	UnitsRemaining     *int64
	TotalUnitsGranted  *int64
	Date               *Date
	AmountGranted      *float64
	TotalAmountGranted *float64
}

type VestFrequency string

// VestingPlanRequest is the payload type of the stox service plan method.
type VestingPlanRequest struct {
	// stock symbol to retrieve plan for
	Symbol string
	// number of stock units granted
	UnitsGranted int64
	// initial grant date of equities
	GrantDate Date
	// date the equities vest completely
	VestDate Date
	// frequency of vesting schedule (monthly, quarterly, yearly)
	VestFrequency VestFrequency
}

// VestingPlanResponse is the result type of the stox service plan method.
type VestingPlanResponse struct {
	Symbol   *string
	Price    *float64
	VestPlan []*VestEvent
}
