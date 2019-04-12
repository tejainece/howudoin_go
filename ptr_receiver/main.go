package main

import "fmt"

type Struct struct {
	X int
}

func (s *Struct) GetX() int {
	s = &Struct{X: 20}
	return s.X
}

func main() {
	s := Struct{X: 10}
	fmt.Println(s.GetX())
}
