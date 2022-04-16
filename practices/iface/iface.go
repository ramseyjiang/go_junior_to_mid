package iface

type Appliance interface {
	TurnOn(name string) string
	TurnOff() string
}
