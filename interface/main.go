package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Printable interface {
	Print()
	GetPrintString() string
}

type PrintableInt struct {
	X int
}

func (p PrintableInt) Print() {
	fmt.Println(p.X)
}

func (p PrintableInt) GetPrintString() string {
	return strconv.Itoa(p.X)
}

func main() {
	pi := PrintableInt{X: 10}
	pi.Print()

	var i Printable = pi
	i.Print()

	pi1, ok := i.(PrintableInt)
	pi1.Print()
	fmt.Println(ok)

	fmt.Println(reflect.TypeOf(pi1))

	var ii interface{} = pi
	fmt.Println(reflect.TypeOf(ii))

	pi2, ok := ii.(PrintableInt)
	pi2.Print()
	fmt.Println(ok)

	fmt.Println(reflect.TypeOf(pi2))
}
