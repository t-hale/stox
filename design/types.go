package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zerologger" // Enables the plugin
)

var Date = Type("Date", String, func() {
	Format(FormatDate)

	//Meta("type:generate:force")
})

var VestFrequency = Type("VestFrequency", String, func() {
	Enum("monthly", "quarterly", "yearly")
})

var VestingPlanRequest = Type("VestingPlanRequest", func() {
	Field(1, "symbol", String, "stock symbol to retrieve plan for")
	Field(2, "units", Int64, "number of stock units granted")
	Field(3, "grant_date", Date, "initial grant date of equities")
	Field(4, "vest_date", Date, "date the equities vest completely")
	Field(5, "vest_frequency", VestFrequency, func() {
		Description("frequency of vesting schedule (monthly, quarterly, yearly)")
	})
	Required("symbol", "units", "grant_date", "vest_date", "vest_frequency")
})

var VestEvent = Type("VestEvent", func() {
	Attribute("UnitsGranted", Int64)
	Attribute("UnitsRemaining", Int64)
	Attribute("Date", Date)

	Meta("type:generate:force")
})

var VestingPlanResponse = Type("VestingPlanResponse", func() {
	Field(1, "VestPlan", ArrayOf(VestEvent))
})
