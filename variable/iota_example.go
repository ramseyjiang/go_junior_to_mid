package main

import (
	"fmt"
	"reflect"
)

type OrderStatus int

func main() {
	const (
		OrderStatusUnknown OrderStatus = iota // 0
		OrderStatusNotActive
		OrderStatusUntriggered
		OrderStatusTriggered
		OrderStatusActive = iota - 1
		OrderStatusCreated
		OrderStatusRejected
		OrderStatusTested
	)

	fmt.Println(OrderStatusUnknown, reflect.TypeOf(OrderStatusNotActive), OrderStatusNotActive, OrderStatusUntriggered, OrderStatusTriggered,
		OrderStatusActive, OrderStatusCreated, OrderStatusRejected, OrderStatusTested)
}
