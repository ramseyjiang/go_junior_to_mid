package iface

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {
	Exec()
	var fan Appliance = Fan(0)
	var coffeePot Appliance = CoffeePot("LuxBrew")
	heater := New()

	wantCoffeePotOutput := "CoffeePot has been turned off."
	wantFanOutput := "Fan has been turned off."
	wantHeaterOutput := "Heater has been turned off."
	t.Run("Test CoffeePot Turn Off", func(t *testing.T) {
		if got := coffeePot.TurnOff(); !reflect.DeepEqual(got, wantCoffeePotOutput) {
			t.Errorf("execute() = %v, want %v", got, wantCoffeePotOutput)
		}

		if got := fan.TurnOff(); !reflect.DeepEqual(got, wantFanOutput) {
			t.Errorf("execute() = %v, want %v", got, wantFanOutput)
		}

		// test for heater
		if got := heater.TurnOff(); !reflect.DeepEqual(got, wantHeaterOutput) {
			t.Errorf("execute() = %v, want %v", got, wantHeaterOutput)
		}
	})
}

func TestCoffeePotTurnOn(t *testing.T) {
	var coffeePot Appliance = CoffeePot("LuxBrew")
	wantCoffeePotOutput := "Hot Chocolate"
	t.Run("Test CoffeePot Turn On", func(t *testing.T) {
		if got := coffeePot.TurnOn(wantCoffeePotOutput); !reflect.DeepEqual(got, wantCoffeePotOutput) {
			t.Errorf("execute() = %v, want %v", got, wantCoffeePotOutput)
		}
	})
}

func TestCoffeePotTurnOff(t *testing.T) {
	var coffeePot Appliance = CoffeePot("LuxBrew")
	wantCoffeePotOutput := "CoffeePot has been turned off."
	t.Run("Test CoffeePot Turn Off", func(t *testing.T) {
		if got := coffeePot.TurnOff(); !reflect.DeepEqual(got, wantCoffeePotOutput) {
			t.Errorf("execute() = %v, want %v", got, wantCoffeePotOutput)
		}
	})
}

func TestFanTurnOn(t *testing.T) {
	var fan Appliance = Fan(1)
	wantFanOutput := "Today is too hot."
	t.Run("Test Fan Turn On", func(t *testing.T) {
		if got := fan.TurnOn(wantFanOutput); !reflect.DeepEqual(got, wantFanOutput) {
			t.Errorf("execute() = %v, want %v", got, wantFanOutput)
		}
	})
}

func TestFanTurnOff(t *testing.T) {
	var fan Appliance = Fan(0)
	wantFanOutput := "Fan has been turned off."
	t.Run("Test Fan Turn Off", func(t *testing.T) {
		if got := fan.TurnOff(); !reflect.DeepEqual(got, wantFanOutput) {
			t.Errorf("execute() = %v, want %v", got, wantFanOutput)
		}
	})
}

func TestHeaterTurnOn(t *testing.T) {
	heater := New()
	wantHeaterOutput := "Today is too cold."
	t.Run("Test Heater Turn On", func(t *testing.T) {
		if got := heater.TurnOn(wantHeaterOutput); !reflect.DeepEqual(got, wantHeaterOutput) {
			t.Errorf("execute() = %v, want %v", got, wantHeaterOutput)
		}
	})
}

func TestHeaterTurnOff(t *testing.T) {
	heater := New()
	wantHeaterOutput := "Heater has been turned off."
	t.Run("Test Heater Turn Off", func(t *testing.T) {
		if got := heater.TurnOff(); !reflect.DeepEqual(got, wantHeaterOutput) {
			t.Errorf("execute() = %v, want %v", got, wantHeaterOutput)
		}
	})
}
