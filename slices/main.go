package main

import (
	"fmt"
)

func main() {
	var ints = []int{5, 10, 15, 20, 25}
	var intsSubSlice = ints[1:3]

	fmt.Println(ints)
	fmt.Println(intsSubSlice)
	fmt.Println(len(ints), cap(ints))
	fmt.Println(len(intsSubSlice), cap(intsSubSlice))

	intsSubSlice = append(intsSubSlice, []int{-30, -35}...)

	fmt.Println(ints)
	fmt.Println(intsSubSlice)
	fmt.Println(len(intsSubSlice), cap(intsSubSlice))

	intsSubSlice = append(intsSubSlice, []int{-40}...)

	fmt.Println(ints)
	fmt.Println(intsSubSlice)
	fmt.Println(len(intsSubSlice), cap(intsSubSlice))

	intsSubSlice[0] = -10
	fmt.Println(ints)
	fmt.Println(intsSubSlice)
}
