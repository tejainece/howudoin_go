package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := bytes.NewBuffer([]byte{})

	buf.WriteString("Hello world!")

	fmt.Println(buf.Bytes())

	fmt.Println(len(buf.Bytes()))

	b := make([]byte, 5)

	n, err := buf.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(string(b))

	n, err = buf.Read(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(string(b))
}
