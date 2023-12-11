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
	Attribute("Symbol", String, "stock symbol to retrieve plan for", func() {
		MinLength(1)
	})
	Attribute("UnitsGranted", Int64, "number of stock units granted", func() {
		Minimum(1)
	})
	Attribute("GrantDate", Date, "initial grant date of equities")
	Attribute("VestDate", Date, "date the equities vest completely")
	Attribute("VestFrequency", VestFrequency, func() {
		Description("frequency of vesting schedule (monthly, quarterly, yearly)")
	})
	Required("Symbol", "UnitsGranted", "GrantDate", "VestDate", "VestFrequency")
})

var VestEvent = Type("VestEvent", func() {
	Attribute("UnitsGranted", Int64)
	Attribute("UnitsRemaining", Int64)
	Attribute("TotalUnitsGranted", Int64)
	Attribute("Date", Date)
	Attribute("AmountGranted", Float64)
	Attribute("TotalAmountGranted", Float64)

	Meta("type:generate:force")
})

var VestingPlanResponse = Type("VestingPlanResponse", func() {
	Attribute("Symbol", String)
	Attribute("Price", Float64)
	Attribute("VestPlan", ArrayOf(VestEvent))
})
