package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/zerologger" // Enables the plugin
)

var Date = Type("Date", String, func() {
	Format(FormatDate)

	//Meta("type:generate:force")
})

var VestEvent = Type("VestDate", func() {
	Attribute("UnitsGranted", Int64)
	Attribute("UnitsRemaining", Int64)
	Attribute("Date", Date)
})

var VestingSchedule = Type("VestingSchedule", func() {
	Attribute("VestEvents", ArrayOf(VestEvent))
})
