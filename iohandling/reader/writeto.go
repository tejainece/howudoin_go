package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	buf := bytes.NewReader([]byte("Hello! Whats up? How are you?"))
	f, _ := os.Create("hello.txt")
	n, _ := buf.WriteTo(f)
	fmt.Println(n)
}
