package iface

func Exec() {
	// The Fan type is int, after it is defined in impl.go. Hence, when initial Fan, it should have any int in parentheses.
	var fan Appliance = Fan(0)
	fan.TurnOn("Today is too hot.")
	fan.TurnOff()

	// The CoffeePot type is string, after it is defined in impl.go. Hence, when initial Fan, it should have any string in parentheses.
	var coffeePot Appliance = CoffeePot("LuxBrew")
	coffeePot.TurnOn("Black Coffee.")
	coffeePot.TurnOff()
	// coffeePot.Brew() // It will has an error, because Brew() didn't define in the interface Appliance.

	// The Heater type is struct, after it is defined in impl.go. Hence, it should be initial with struct.
	var heater Appliance = Heater{ut: 26, nt: 20}
	heater.TurnOn("Electricity Heater")
	heater.TurnOff()
}
