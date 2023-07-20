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
	Field(1, "Symbol", String, "stock symbol to retrieve plan for")
	Field(2, "UnitsGranted", Float64, "number of stock units granted")
	Field(3, "GrantDate", Date, "initial grant date of equities")
	Field(4, "VestDate", Date, "date the equities vest completely")
	Field(5, "VestFrequency", VestFrequency, func() {
		Description("frequency of vesting schedule (monthly, quarterly, yearly)")
	})
	Required("Symbol", "UnitsGranted", "GrantDate", "VestDate", "VestFrequency")
})

var VestEvent = Type("VestEvent", func() {
	Attribute("UnitsGranted", Float64)
	Attribute("UnitsRemaining", Float64)
	Attribute("TotalUnitsGranted", Float64)
	Attribute("Date", Date)
	Attribute("AmountGranted", Float64)
	Attribute("TotalAmountGranted", Float64)

	Meta("type:generate:force")
})

var VestingPlanResponse = Type("VestingPlanResponse", func() {
	Field(1, "VestPlan", ArrayOf(VestEvent))
})
