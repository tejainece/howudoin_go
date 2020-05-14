package main

import (
	"fmt"
	"runtime/debug"
)

func PanicerRecovered() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	panic("opps")
}

func PanicerDefered() {
	defer func() {
		fmt.Println("Hello")
		fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	}()

	panic("opps")
}

func main() {
	PanicerRecovered()
	PanicerDefered()
}
