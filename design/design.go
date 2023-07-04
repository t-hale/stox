package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zerologger" // Enables the plugin
)

var _ = API("stox", func() {
	Title("Stock Option Planner Service")
	Description("Service for planning stock options including vesting schedules")
	Server("server", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			//URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("stox", func() {
	Description("The stox service provides advisors with a comprehensive view of a particular stock schedule.")

	Method("plan", func() {
		Payload(func() {
			Field(1, "symbol", String, "stock symbol to retrieve plan for")
			Field(2, "units", Int64, "number of stock units granted")
			Field(3, "grant_date", String, "initial grant date of equities")
			Field(4, "vest_date", String, "date the equities vest completely")
			Field(5, "vest_frequency", String, "frequency of vesting schedule (monthly, quarterly, yearly)")
			Required("symbol", "units", "grant_date", "vest_date", "vest_frequency")
		})

		Result(VestingSchedule)

		HTTP(func() {
			POST("/plan")
		})

		//GRPC(func() {
		//})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
