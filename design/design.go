package design

import (
	. "goa.design/goa/v3/dsl"
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
	Description("The sop service provides advisors with a comprehensive view of a particular stock schedule.")

	Method("plan", func() {
		Payload(func() {
			Field(1, "symbol", String, "stock symbol to retrieve plan for")
			Field(2, "units", Int64, "number of stock units granted")
			Required("symbol", "units")
		})

		Result(Int)

		HTTP(func() {
			POST("/plan")
		})

		//GRPC(func() {
		//})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
