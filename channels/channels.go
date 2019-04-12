package main

import (
	"fmt"
)

func main() {
	var ch chan byte
	fmt.Println(<-ch)
}
