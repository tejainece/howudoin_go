package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var err interface{}

	err = ioutil.WriteFile("hello.txt", []byte("Hello world!"), os.ModePerm)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile("hello.txt")
	fmt.Println(string(bytes))

	os.Remove("hello.txt")
}
