package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("How are you doing? I am doing fine!")
	byt := make([]byte, 10, 10)
	n, err := reader.Read(byt)
	if err != nil {
		panic(fmt.Sprint("Not enough space: ", err))
	}
	fmt.Println("Read", n, "bytes:", string(byt))
	fmt.Println("Buffer len:", reader.Len(), "Buffer Size:", reader.Size())
	
	writer := bytes.NewBuffer([]byte{})
	
	n1, err := reader.WriteTo(writer)
	if err != nil {
		panic(fmt.Sprint("Error writing: ", err))
	}
	fmt.Println("Written", n1, "bytes:", writer.String())
	
	w1 := bytes.NewBuffer([]byte{})
	w2 := bytes.NewBuffer([]byte{})
	
	w3 := io.MultiWriter(w1, w2)
	reader.Seek(0, io.SeekStart)
	n1, err = reader.WriteTo(w3)
	if err != nil {
		panic(fmt.Sprint("Error writing: ", err))
	}
	fmt.Println("Written", n1, "bytes:", w1.String())
	fmt.Println("Written", n1, "bytes:", w2.String())
}