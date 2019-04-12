package main

import "fmt"

func main() {
	var i int = 5
	var x interface{} = &i

	fmt.Println(x)

	var iv int = *x.(*int)

	fmt.Println(iv)

	iv1, _ := x.(int)
	fmt.Println(iv1)
	// TODO
}
