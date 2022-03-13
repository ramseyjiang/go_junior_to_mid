package main

import (
	"fmt"
	"reflect"
)

type OrderStatus int

func main() {
	const (
		OrderstatusUnknown OrderStatus = iota // 0
		OrderstatusNotactive
		OrderstatusUntriggered
		OrderstatusTriggered
		OrderstatusActive = iota - 1
		OrderstatusCreated
		OrderstatusRejected
		OrderstatusTested
	)

	fmt.Println(OrderstatusUnknown, reflect.TypeOf(OrderstatusNotactive), OrderstatusNotactive, OrderstatusUntriggered, OrderstatusTriggered,
		OrderstatusActive, OrderstatusCreated, OrderstatusRejected, OrderstatusTested)
}
