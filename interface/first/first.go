package main

import (
	"fmt"
)

type Vehicle interface {
	NumWheels() int
}

type Car struct {
}

func (v Car) NumWheels() int {
	return 4
}

func main() {
	var vehicle Vehicle = Car{}
	fmt.Println(vehicle.NumWheels())

	var c, ok = vehicle.(Car) // Type assertion
	if !ok {
		panic("Invalid type assertion!")
	}
	fmt.Println(c.NumWheels())
}
