package iface

func Exec() {
	var fan Appliance = Fan("Breeze")
	fan.TurnOn("Today is too hot.")
	fan.TurnOff()

	var coffeePot Appliance = CoffeePot("LuxBrew")
	coffeePot.TurnOn("Black Coffee.")
	coffeePot.TurnOff()
	// coffeePot.Brew() // It will has an error, because Brew() didn't define in the interface Appliance.
}
