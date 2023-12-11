package stoxapi

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"github.com/google/go-cmp/cmp"
	"github.com/t-hale/stox/gen/stox"
	"github.com/t-hale/stox/utils"
	"google.golang.org/protobuf/proto"
	"strings"
	"testing"
)

var (
	fakeStoxSrvc stoxsrvc
)

func setupTests() {
	fakeStoxSrvc = stoxsrvc{
		logger: nil,
	}
}

func init() {
	setupTests()
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestCalculateVestDates(t *testing.T) {

	type testCase struct {
		name    string
		req     *stox.VestingPlanRequest
		trade   *marketdata.Trade
		want    *stox.VestingPlanResponse
		wantErr error
	}

	testCases := []testCase{
		{
			name: "monthly plan",
			req: &stox.VestingPlanRequest{
				UnitsGranted:  120,
				GrantDate:     "2023-09-13",
				VestDate:      "2024-09-13",
				VestFrequency: "monthly",
				Symbol:        "FOO",
			},
			trade: &marketdata.Trade{
				Price: 100.13,
			},
			want: &stox.VestingPlanResponse{
				Symbol: utils.PtrTo("FOO"),
				Price:  proto.Float64(100.13),
				VestPlan: []*stox.VestEvent{
					{
						Date:               utils.PtrTo(stox.Date("2023-09-13")),
						UnitsGranted:       proto.Int64(0),
						TotalUnitsGranted:  proto.Int64(0),
						UnitsRemaining:     proto.Int64(120),
						AmountGranted:      proto.Float64(0.0),
						TotalAmountGranted: proto.Float64(0.0),
					},
					{
						Date:               utils.PtrTo(stox.Date("2023-10-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(10),
						UnitsRemaining:     proto.Int64(110),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(1001.3),
					},
					{
						Date:               utils.PtrTo(stox.Date("2023-11-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(20),
						UnitsRemaining:     proto.Int64(100),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(2002.6),
					},
					{
						Date:               utils.PtrTo(stox.Date("2023-12-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(30),
						UnitsRemaining:     proto.Int64(90),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(3003.8999999999996),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-01-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(40),
						UnitsRemaining:     proto.Int64(80),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(4005.2),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-02-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(50),
						UnitsRemaining:     proto.Int64(70),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(5006.5),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-03-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(60),
						UnitsRemaining:     proto.Int64(60),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(6007.799999999999),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-04-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(70),
						UnitsRemaining:     proto.Int64(50),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(7009.099999999999),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-05-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(80),
						UnitsRemaining:     proto.Int64(40),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(8010.4),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-06-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(90),
						UnitsRemaining:     proto.Int64(30),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(9011.699999999999),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-07-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(100),
						UnitsRemaining:     proto.Int64(20),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(10013.0),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-08-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(110),
						UnitsRemaining:     proto.Int64(10),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(11014.3),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-09-13")),
						UnitsGranted:       proto.Int64(10),
						TotalUnitsGranted:  proto.Int64(120),
						UnitsRemaining:     proto.Int64(0),
						AmountGranted:      proto.Float64(1001.3),
						TotalAmountGranted: proto.Float64(12015.599999999999),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "quarterly plan",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2024-09-13",
				VestFrequency: "quarterly",
				Symbol:        "FOO",
				UnitsGranted:  1000,
			},
			trade: &marketdata.Trade{
				Price: 72.45,
			},
			want: &stox.VestingPlanResponse{
				Symbol: utils.PtrTo("FOO"),
				Price:  proto.Float64(72.45),
				VestPlan: []*stox.VestEvent{
					{
						Date:               utils.PtrTo(stox.Date("2023-09-13")),
						UnitsGranted:       proto.Int64(0),
						TotalUnitsGranted:  proto.Int64(0),
						UnitsRemaining:     proto.Int64(1000),
						AmountGranted:      proto.Float64(0.0),
						TotalAmountGranted: proto.Float64(0.0),
					},
					{
						Date:               utils.PtrTo(stox.Date("2023-12-13")),
						UnitsGranted:       proto.Int64(250),
						TotalUnitsGranted:  proto.Int64(250),
						UnitsRemaining:     proto.Int64(750),
						AmountGranted:      proto.Float64(18112.5),
						TotalAmountGranted: proto.Float64(18112.5),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-03-13")),
						UnitsGranted:       proto.Int64(250),
						TotalUnitsGranted:  proto.Int64(500),
						UnitsRemaining:     proto.Int64(500),
						AmountGranted:      proto.Float64(18112.5),
						TotalAmountGranted: proto.Float64(36225.0),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-06-13")),
						UnitsGranted:       proto.Int64(250),
						TotalUnitsGranted:  proto.Int64(750),
						UnitsRemaining:     proto.Int64(250),
						AmountGranted:      proto.Float64(18112.5),
						TotalAmountGranted: proto.Float64(54337.5),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-09-13")),
						UnitsGranted:       proto.Int64(250),
						TotalUnitsGranted:  proto.Int64(1000),
						UnitsRemaining:     proto.Int64(0),
						AmountGranted:      proto.Float64(18112.5),
						TotalAmountGranted: proto.Float64(72450.0),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "yearly plan",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2024-09-13",
				VestFrequency: "yearly",
				Symbol:        "BAR",
				UnitsGranted:  500,
			},
			trade: &marketdata.Trade{
				Price: 5.71,
			},
			want: &stox.VestingPlanResponse{
				Symbol: utils.PtrTo("BAR"),
				Price:  proto.Float64(5.71),
				VestPlan: []*stox.VestEvent{
					{
						Date:               utils.PtrTo(stox.Date("2023-09-13")),
						AmountGranted:      proto.Float64(0.0),
						UnitsGranted:       proto.Int64(0),
						TotalUnitsGranted:  proto.Int64(0),
						UnitsRemaining:     proto.Int64(500),
						TotalAmountGranted: proto.Float64(0.0),
					},
					{
						Date:               utils.PtrTo(stox.Date("2024-09-13")),
						AmountGranted:      proto.Float64(500 * 5.71),
						UnitsGranted:       proto.Int64(500),
						UnitsRemaining:     proto.Int64(0),
						TotalUnitsGranted:  proto.Int64(500),
						TotalAmountGranted: proto.Float64(2855.0),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "monthly plan, bad end",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2024-09-07",
				VestFrequency: "monthly",
			},
			want:    nil,
			wantErr: fmt.Errorf("GrantDate, VestDate and VestFrequency do not align"),
		},
		{
			name: "yearly plan, bad end",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2024-03-13",
				VestFrequency: "yearly",
			},
			want:    nil,
			wantErr: fmt.Errorf("GrantDate, VestDate and VestFrequency do not align"),
		},
		{
			name: "vest date before grant date",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2022-09-13",
				VestFrequency: "monthly",
			},
			want:    nil,
			wantErr: fmt.Errorf("GrantDate must occur before VestDate"),
		},
	}

	for _, tc := range testCases {
		got, err := fakeStoxSrvc.calculateVestingPlan(tc.req, tc.trade)

		if err != nil {
			if tc.wantErr == nil {
				t.Errorf("TestCalculateVestDates: Received unexpected error (got : %s , want : %s)", err, tc.wantErr)
			}

			if tc.wantErr != nil && !strings.Contains(err.Error(), tc.wantErr.Error()) {
				t.Errorf("TestCalculateVestDates: Received unexpected error (got : %s , want : %s)", err, tc.wantErr)
			}

		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("TestCalculateVestDates() mismatch (-want +got):\n%s", diff)
		}

	}
}
