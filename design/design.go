package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zerologger" // Enables the plugin
)

var _ = API("stox", func() {
	Title("Stock Option Planner Service")
	Description("Service for planning stock options including vesting schedules")
	Server("server", func() {
		Host("0.0.0.0", func() {
			URI("http://0.0.0.0:8080")
		})
	})
})

var _ = Service("stox", func() {
	Description("The stox service provides advisors with a comprehensive view of a particular stock schedule.")

	Method("plan", func() {
		Meta("openapi:extension:x-google-backend", `{"address":"https://stox-xuqgkxzwta-uc.a.run.app"}`)
		Payload(VestingPlanRequest)
		Result(VestingPlanResponse)

		HTTP(func() {
			POST("/plan")
		})

		//GRPC(func() {
		//})
	})

	//Files("/openapi.json", "./gen/http/openapi.json")
})
