package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("Point{X: %.2f, Y: %.2f}", p.X, p.Y)
}

func main() {
	fmt.Println(Point{X: 10, Y: 10})
	// TODO
}
