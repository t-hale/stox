package stoxapi

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/t-hale/stox/gen/stox"
	"github.com/t-hale/stox/utils"
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
			},
			want: &stox.VestingPlanResponse{
				VestPlan: []*stox.VestEvent{
					{
						Date: utils.PtrTo(stox.Date("2023-09-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2023-10-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2023-11-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2023-12-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-01-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-02-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-03-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-04-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-05-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-06-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-07-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-08-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-09-13")),
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
			},
			want: &stox.VestingPlanResponse{
				VestPlan: []*stox.VestEvent{
					{
						Date: utils.PtrTo(stox.Date("2023-09-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2023-12-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-03-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-06-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-09-13")),
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
			},
			want: &stox.VestingPlanResponse{
				VestPlan: []*stox.VestEvent{
					{
						Date: utils.PtrTo(stox.Date("2023-09-13")),
					},
					{
						Date: utils.PtrTo(stox.Date("2024-09-13")),
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
			wantErr: fmt.Errorf("grant_date, vest_date and frequency do not align"),
		},
		{
			name: "yearly plan, bad end",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2024-03-13",
				VestFrequency: "yearly",
			},
			want:    nil,
			wantErr: fmt.Errorf("grant_date, vest_date and frequency do not align"),
		},
		{
			name: "vest date before grant date",
			req: &stox.VestingPlanRequest{
				GrantDate:     "2023-09-13",
				VestDate:      "2022-09-13",
				VestFrequency: "monthly",
			},
			want:    nil,
			wantErr: fmt.Errorf("grant_date, vest_date and frequency do not align"),
		},
	}

	for _, tc := range testCases {
		got, err := fakeStoxSrvc.calculateVestingPlan(tc.req)

		if err != nil {
			if tc.wantErr == nil {
				t.Errorf("TestCalculateVestDates: Received unexpected error (got : %s , want : %s)", err, tc.wantErr)
			}

		}

		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("TestCalculateVestDates() mismatch (-want +got):\n%s", diff)
		}

	}
}
