package main

import (
	"fmt"
)

func main() {
	var x interface{} = 5

	switch v := x.(type) {
	case int:
		fmt.Println("int", v)
	default:
		fmt.Println("other", v)
	}
	// TODO
}
